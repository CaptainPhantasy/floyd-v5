# Extensibility Selection UI - Optimal UX Design

## Design Decision: Unified Pattern

After analysis, the optimal pattern is **Category Tabs + Collapsible Groups**:

```
┌─────────────────────────────────────────────────────────────────┐
│ Agent Library                                           90     │
│ [●All] [Arch] [Infra] [Orch] [Code] [Sec] [Qual] [Test] [...]  │
├─────────────────────────────────────────────────────────────────┤
│ Filter: [_______________________________]                       │
├─────────────────────────────────────────────────────────────────┤
│ ▾ ARCHITECTURE (36)                                            │
│   Data Flow Cartographer                                       │
│     Maps data, queries, and control flow...                    │
│                                                                 │
│   Monorepo Boundary Cartographer                               │
│     Maps module boundaries...                                   │
│                                                                 │
│ ▸ INFRASTRUCTURE (16)                                          │
│ ▸ ORCHESTRATION (7)                                            │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│ tab cycle | 1-9 cat | / search | e expand | enter select | esc │
└─────────────────────────────────────────────────────────────────┘
```

## Key Bindings (Consistent Across All Libraries)

| Key         | Action                          |
|-------------|--------------------------------|
| `↑/↓` or `j/k` | Navigate items              |
| `Tab`       | Next category tab              |
| `Shift+Tab` | Previous category tab          |
| `0`         | Show all categories            |
| `1-9`       | Jump to category by number     |
| `/`         | Focus filter input             |
| `Enter`     | Select item                    |
| `e`         | Expand/collapse description    |
| `?`         | Show keyboard help overlay     |
| `Esc`       | Close dialog                   |

## Category Mappings

### Agents (90 total)
```
0 = All        (90)
1 = Architecture (36)
2 = Infrastructure (16)
3 = Orchestration (7)
4 = Coding (6)
5 = Security (6)
6 = Quality (5)
7 = Testing (4)
8 = Monitoring (3)
9 = DX (3)
```

### Skills (20 planned)
```
0 = All
1 = Git
2 = Testing
3 = Linting
4 = Refactoring
5 = Documentation
6 = Deployment
7 = Security
8 = Data
9 = Debugging
```

### Commands (Unified)
```
Tab cycles: System → Agents → Skills → Plugins → System
```

## Implementation

1. Create shared `CategoryList` component
2. Update `AgentLibrary` to use it
3. Update `SkillsLibrary` to use it
4. Update `Commands` dialog with same pattern
5. Add `PluginsLibrary` dialog
