# State Management Patterns for Ink CLI

This reference covers state management patterns for agentic Ink CLI applications following CLI-X 2026 principles.

## Engine State vs UI State

```typescript
// Engine State - facts from orchestration/runtime
interface EngineState {
  agentStatus: 'idle' | 'thinking' | 'executing' | 'done' | 'error';
  activeTool: string | null;
  steps: Step[];
  logs: LogEntry[];
  progress: number;
}

// UI State - presentation state
interface UIState {
  overlayOpen: boolean;
  focusedPane: 'main' | 'side' | 'logs';
  selectedStepIndex: number;
  searchQuery: string;
  followMode: boolean;
}
```

## Zustand Store Pattern

```typescript
import { create } from 'zustand';
import { persist } from 'zustand/middleware';

interface AppState {
  engine: EngineState;
  ui: UIState;

  // Actions - UI never mutates engine directly
  engineActions: {
    setAgentStatus: (status: AgentStatus) => void;
    addLog: (log: LogEntry) => void;
    updateStep: (id: string, update: Partial<Step>) => void;
  };
  uiActions: {
    openOverlay: () => void;
    closeOverlay: () => void;
    setFocusedPane: (pane: UIPane) => void;
  };
}

const useStore = create<AppState>()(
  persist(
    (set, get) => ({
      engine: {
        agentStatus: 'idle',
        activeTool: null,
        steps: [],
        logs: [],
        progress: 0,
      },
      ui: {
        overlayOpen: false,
        focusedPane: 'main',
        selectedStepIndex: 0,
        searchQuery: '',
        followMode: true,
      },

      engineActions: {
        setAgentStatus: (status) =>
          set((state) => ({
            engine: { ...state.engine, agentStatus: status },
          })),

        addLog: (log) =>
          set((state) => ({
            engine: {
              ...state.engine,
              logs: [...state.engine.logs.slice(-99), log], // Bounded to 100
            },
          })),

        updateStep: (id, update) =>
          set((state) => ({
            engine: {
              ...state.engine,
              steps: state.engine.steps.map((s) =>
                s.id === id ? { ...s, ...update } : s
              ),
            },
          })),
      },

      uiActions: {
        openOverlay: () =>
          set((state) => ({ ui: { ...state.ui, overlayOpen: true } })),
        closeOverlay: () =>
          set((state) => ({ ui: { ...state.ui, overlayOpen: false } })),
        setFocusedPane: (pane) =>
          set((state) => ({ ui: { ...state.ui, focusedPane: pane } })),
      },
    }),
    { name: 'floyd-cli-state' }
  )
);
```

## Event Ring Buffer Pattern

```typescript
class RingBuffer<T> {
  private buffer: T[] = [];
  private pointer = 0;

  constructor(private size: number) {}

  push(item: T): void {
    if (this.buffer.length < this.size) {
      this.buffer.push(item);
    } else {
      this.buffer[this.pointer] = item;
      this.pointer = (this.pointer + 1) % this.size;
    }
  }

  toArray(): T[] {
    return [
      ...this.buffer.slice(this.pointer),
      ...this.buffer.slice(0, this.pointer),
    ];
  }

  clear(): void {
    this.buffer = [];
    this.pointer = 0;
  }
}

// Usage for logs
const logBuffer = new RingBuffer<LogEntry>(100);
```

## Reducer Pattern for Event Streams

```typescript
type Event =
  | { type: 'AGENT_STATUS_CHANGED'; status: AgentStatus }
  | { type: 'TOOL_STARTED'; tool: string; id: string }
  | { type: 'TOOL_COMPLETED'; id: string; result: unknown }
  | { type: 'LOG_ADDED'; log: LogEntry }
  | { type: 'STEP_UPDATED'; id: string; status: StepStatus };

interface ViewModel {
  agentStatus: AgentStatus;
  activeTools: Map<string, ToolState>;
  logs: LogEntry[];
  steps: Step[];
}

function reducer(state: ViewModel, event: Event): ViewModel {
  switch (event.type) {
    case 'AGENT_STATUS_CHANGED':
      return { ...state, agentStatus: event.status };

    case 'TOOL_STARTED':
      return {
        ...state,
        activeTools: new Map(state.activeTools).set(event.id, {
          name: event.tool,
          status: 'running',
          startTime: Date.now(),
        }),
      };

    case 'TOOL_COMPLETED':
      const tools = new Map(state.activeTools);
      const tool = tools.get(event.id);
      if (tool) {
        tool.status = event.result ? 'done' : 'failed';
        tool.endTime = Date.now();
      }
      return { ...state, activeTools: tools };

    case 'LOG_ADDED':
      return {
        ...state,
        logs: [...state.logs.slice(-99), event.log],
      };

    case 'STEP_UPDATED':
      return {
        ...state,
        steps: state.steps.map((s) =>
          s.id === event.id ? { ...s, status: event.status } : s
        ),
      };

    default:
      return state;
  }
}
```

## Command Dispatcher Pattern

```typescript
// UI never mutates engine state directly
type Command =
  | { type: 'START_TASK' }
  | { type: 'CANCEL_TASK' }
  | { type: 'RETRY_STEP'; stepId: string }
  | { type: 'OPEN_LOGS' }
  | { type: 'COPY_SUMMARY' };

interface CommandDispatcher {
  dispatch: (command: Command) => void;
}

// In component
const { dispatch } = useCommandDispatcher();

// Usage
<Flex>
  <Button onPress={() => dispatch({ type: 'START_TASK' })}>Start</Button>
  <Button onPress={() => dispatch({ type: 'CANCEL_TASK' })}>Cancel</Button>
</Flex>
```

## Throttled UI Updates

```typescript
import { useMemo, useEffect, useRef } from 'react';

// Throttle high-frequency updates to 30-60 FPS
function useThrottledValue<T>(value: T, fps = 30): T {
  const throttled = useRef(value);
  const lastUpdate = useRef(Date.now());

  useEffect(() => {
    const interval = 1000 / fps;
    const now = Date.now();

    if (now - lastUpdate.current >= interval) {
      throttled.current = value;
      lastUpdate.current = now;
    }
  }, [value, fps]);

  return throttled.current;
}

// Usage with streaming tokens
const displayTokens = useThrottledValue(tokens);
```

## Persistence Rules

```typescript
// Persist only what survives restarts
const persistedState = {
  // Always persist
  ui: {
    overlayOpen: false, // Don't persist - reset on restart
    focusedPane: 'main',
    searchQuery: '', // Don't persist - reset on restart
  },

  // Persist engine recovery state
  engine: {
    steps: [], // Persist for resume capability
    sessionId: 'uuid', // Persist
  },

  // Never persist
  transient: {
    logs: [], // Too large, not needed on restart
    tokens: [], // Ephemeral stream data
  },
};
```
