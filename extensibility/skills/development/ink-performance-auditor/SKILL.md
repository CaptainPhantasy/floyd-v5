---
name: ink-performance-auditor
description: Performance critic and optimization specialist for Ink (React CLI) apps with streaming events. Use this skill when streaming output exists (tokens/logs/progress), IPC emits more than 5 events/sec, UI has frequently updating lists, layout changes dynamically, app feels jittery/high CPU, or long sessions occur. Audits CLI features and outputs: throughput failure modes, measurable performance gates, concrete refactors (batching, throttling, memoization), and stress-test harness plans. Acts as a perf SRE for terminal dashboards.
---

# Ink Performance & Event Throughput Auditor (2026)

## Skill Mission

Prevent the "cool CLI" from becoming a laggy, memory-leaking, flickery mess under real throughput.

When invoked, this skill audits a CLI feature or subsystem and outputs:

- Top throughput failure modes (re-render storms, unbounded buffers, layout churn)
- Measurable performance gates (FPS-ish, max renders/sec, buffer sizes, CPU limits)
- Concrete refactors (batching, throttling, memoization, reducer design, ring buffers)
- A stress-test harness plan that any agent can run locally

**Acts as:** "Performance SRE" for a terminal dashboard.

---

## When to Invoke

Use this skill whenever **any** of these are true:

- Streaming output exists (tokens/logs/progress events)
- IPC emits more than ~5 events/sec
- UI has lists (agents/tools/logs) that update frequently
- Layout changes dynamically (panes collapse/expand, overlays, resizing)
- App feels "jittery," high CPU usage, or flickers
- Long sessions happen (minutes to hours)

---

## Required Inputs

Caller provides (best effort; skill will proceed with defaults):

| Input | Description |
|-------|-------------|
| **Subsystem/Feature name** | e.g., "Monitor Pane", "Tool Timeline", "Streaming Viewport" |
| **Event sources** | IPC / engine / tool logs / LLM tokens |
| **Expected throughput** | events/sec, tokens/sec, tool log volume |
| **UI surfaces affected** | Which panes/components update |
| **Session length target** | 5 min? 2 hours? |
| **Constraints** | No extra deps? Cross-platform? Node version? Ink renderer? |

---

## Mandatory Output Format

This skill always returns the following sections:

### A) Threat Model (Perf Failure Modes)

Top 5 ways this will break under load, **ranked**.

### B) Observability Plan (What to Measure)

Counters/timers/log points to add.

### C) Performance Gates (Numbers, Not Vibes)

Hard thresholds for render rate, buffer sizes, CPU, memory.

### D) Architecture Fixes (Concrete)

Specific patterns with implementation notes.

### E) Component Rules (Do/Don't List)

Ink/React anti-patterns and safe patterns.

### F) Stress Test Harness

How to simulate load deterministically and verify gates.

### G) Definition of Done

Checklist + acceptance tests.

---

## A) Threat Model (Perf Failure Modes) — Ranked

| Rank | Failure Mode | Description |
|------|--------------|-------------|
| **1** | **Re-render storm** | Every event triggers state update → entire tree re-renders |
| **2** | **Unbounded buffers** | Logs/tokens/events grow forever → memory climbs → crash or GC thrash |
| **3** | **Layout churn / flicker** | Components change height/width frequently → terminal redraw cost spikes |
| **4** | **High-frequency list diffs** | Re-rendering large lists without memoization or stable keys |
| **5** | **Async race + double-run** | Multiple streams update same state slice → inconsistent UI + wasted renders |

---

## B) Observability Plan (Add These Counters Everywhere)

Add a lightweight metrics module (no heavy deps) with counters you can print in a debug overlay.

**Required counters:**

| Counter | Description |
|---------|-------------|
| `events_in_total` | By type (token, log, ipc, tool) |
| `events_in_per_sec` | Rolling window |
| `renders_total` | Total render count |
| `renders_per_sec` | Rolling window |
| `state_updates_total` | Number of state changes |
| `ring_buffer_dropped_lines_total` | Lines dropped due to buffer cap |
| `avg_reduce_time_ms` | Event → view-model reduce latency |
| `avg_render_commit_ms` | Time between state set and next tick |
| `max_list_size_logs/tools/agents` | Peak list sizes |

**Debug UI overlay:**

- Toggle with `Ctrl+D`
- Shows: events/sec, renders/sec, buffer sizes, dropped lines, CPU-ish (if available), memory (`process.memoryUsage()`)

---

## C) Performance Gates (Hard Numbers)

Use these default gates unless you have a reason to change them.

### Render / Update Gates

| Metric | Target | Hard Ceiling |
|--------|--------|--------------|
| UI update rate cap | 10–20 updates/sec | 30 max |
| Renders/sec | ≤ 20 | **40 (perf bug above this)** |
| Event reduction latency (avg) | < 2ms | p95 < 8ms |

### Buffer Gates

| Buffer Type | Limit |
|-------------|-------|
| Log ring buffer | 2,000 lines per stream surface |
| Token stream buffer | 20,000 chars per active stream |
| Tool output store | Summaries + last N lines; details on-demand |

### Session Gates (1–2 hours usage)

| Metric | Requirement |
|--------|-------------|
| RSS memory growth | **Must plateau** (no linear growth with time) |
| GC churn | No visible stutter during sustained streaming |

### Layout Gates

- No full-screen reflow on per-token updates
- Expanding/collapsing sections must be **user-triggered**, not automatic

---

## D) Architecture Fixes (Concrete Patterns)

### D1) Batch Events → One UI Update Per Frame

**Rule:** Events can arrive fast; UI must update at a controlled cadence.

**Pattern:**

1. Push incoming events into an in-memory queue
2. Every 50–100ms, drain queue → reduce to view model → single state set

**Result:** 200 events/sec → ~10–20 UI updates/sec.

```typescript
const eventQueue = useRef<Event[]>([]);

// Push events as they arrive
function pushEvent(event: Event) {
  eventQueue.current.push(event);
}

// Drain on interval
setInterval(() => {
  if (eventQueue.current.length === 0) return;

  const events = eventQueue.current.splice(0);
  const newViewModel = reducer(currentViewModel, events);
  setViewModel(newViewModel);
}, 50); // 20 updates/sec max
```

### D2) Reduce into a View Model (Not Raw Event History)

Keep **minimal** state needed to render:

- Current agent statuses
- Current tool runs + last update
- Bounded logs
- Progress checklist states

**Avoid:** Storing every raw event forever.

### D3) Ring Buffers for Anything Stream-Like

Implement a ring buffer utility with:

- `maxItems`
- `push(item)` returning `droppedCount`
- Optional `mergeConsecutiveSimilar()` for repeated log spam

```typescript
class RingBuffer<T> {
  private buffer: T[] = [];
  private pointer = 0;
  private dropped = 0;

  constructor(private size: number) {}

  push(item: T): number {
    if (this.buffer.length < this.size) {
      this.buffer.push(item);
      return 0;
    }
    this.buffer[this.pointer] = item;
    this.pointer = (this.pointer + 1) % this.size;
    return ++this.dropped;
  }

  toArray(): T[] {
    return [
      ...this.buffer.slice(this.pointer),
      ...this.buffer.slice(0, this.pointer),
    ];
  }

  getDroppedCount(): number {
    return this.dropped;
  }
}
```

### D4) Separate "Hot" and "Cold" State

| State Type | Definition | Examples |
|------------|------------|----------|
| **Hot** | Changes frequently | Progress counters, last log lines |
| **Cold** | Changes rarely | Theme, layout mode, config |

Keep hot state in a small store slice so updates don't invalidate the world.

### D5) Stable Keys + Memoization for Lists

**Rules:**

- Every list item needs a stable key (`toolRunId`, `agentId`) — **never array index**
- Use `React.memo` for list rows
- Provide `areEqual(prev, next)` when row updates are sparse

```typescript
// Bad - unstable keys
{items.map((item, index) => (
  <Row key={index} item={item} />
))}

// Good - stable keys
{items.map((item) => (
  <Row key={item.id} item={item} />
))}

// Better - memoized
const Row = memo(function Row({ item }: { item: Item }) {
  return <Box>{item.label}</Box>;
}, areEqual);
```

### D6) Don't Rerender Layout Chrome for Stream Updates

Header/footer should not depend on rapidly changing props.

Feed them only **coarse-grained state** (mode, major status).

---

## E) Component Rules (Do / Don't)

### DO

- Use a reducer or store that batches updates
- Use bounded height containers for logs (fixed height)
- Render "last N lines" only
- Prefer "summary + expand" over "everything at once"
- Use `useMemo` for derived filtered lists (commands, files)
- Use `useRef` to store raw high-frequency stuff until next batch tick

### DON'T

- Don't call `setState` for every token/log line
- Don't map giant arrays each tick without memoization
- Don't auto-expand "thinking" blocks as they stream (layout churn)
- Don't store raw events indefinitely
- Don't use unstable keys (index keys) for dynamic lists

---

## F) Stress Test Harness (Deterministic Load)

The harness should simulate:

| Scenario | Simulation |
|----------|------------|
| High event rate IPC | 100 events/sec |
| LLM token stream | 50–200 tokens/sec |
| Tool logs burst | 500 lines in 2 seconds |
| Long session | 10 minutes continuous |

**Harness Requirements:**

1. Must run without network
2. Must be reproducible (seeded RNG)
3. Must print metrics summary at end:
   - max renders/sec observed
   - dropped lines count
   - peak memory
   - average reduce time
   - **Pass/Fail**

**Pass/Fail Criteria:**

| Condition | Result |
|-----------|--------|
| Renders/sec spikes over 40 | **FAIL** |
| Buffers exceed limits | **FAIL** |
| Memory increases linearly | **FAIL** |
| UI becomes non-responsive | **FAIL** |
| All gates met | **PASS** |

---

## G) Definition of Done (Perf Auditor)

A subsystem passes when:

- [ ] Update cadence is capped (batch tick present)
- [ ] All streams are bounded (ring buffers, truncation)
- [ ] Lists are stable-keyed and memoized
- [ ] Layout remains stable under streaming (no flicker)
- [ ] Stress harness passes gates with printed metrics
- [ ] "Debug overlay" can show current throughput stats

---

## Skill Invocation Template

```
SKILL: Ink Performance & Event Throughput Auditor (2026)

Subsystem:
Event Sources:
Expected Throughput:
UI Surfaces Affected:
Session Length Target:
Constraints:

Deliver:
A) Threat Model (Perf Failure Modes)
B) Observability Plan (What to measure)
C) Performance Gates (Numbers, not vibes)
D) Architecture Fixes (Concrete)
E) Component Rules (Do/Don't list)
F) Stress Test Harness
G) Definition of Done
```

---

## Related Skills

- **cli-x-2026** — Use this skill first when designing the UI components that this skill will audit
- **mcp-builder** — Use when MCP tool throughput is causing performance issues
- **chrome-extension-bridge** — Use when auditing browser bridge performance
