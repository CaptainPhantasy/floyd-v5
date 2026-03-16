# Ink Component Patterns Reference

This reference provides common patterns for Ink (React for CLI) components.

## Box Layout Patterns

```typescript
import { Box, Text } from 'ink';

// Fixed size
<Box width={20} height={10}>
  <Text>Fixed dimensions</Text>
</Box>

// Flex grow
<Box flexGrow={1}>
  <Text>Takes remaining space</Text>
</Box>

// Padding and margin
<Box padding={1} paddingLeft={2} marginBottom={1}>
  <Text>Padded content</Text>
</Box>

// Border
<Box borderStyle="round" borderColor="green">
  <Text>Bordered content</Text>
</Box>
```

## Flex Layout Patterns

```typescript
import { Box, Flex } from 'ink';

// Row layout
<Flex alignItems="center" gap={1}>
  <Text>Label:</Text>
  <TextInput value={value} onChange={setValue} />
</Flex>

// Column layout
<Flex flexDirection="column" gap={1}>
  <Text>Line 1</Text>
  <Text>Line 2</Text>
  <Text>Line 3</Text>
</Flex>

// Space between
<Flex justifyContent="space-between">
  <Text>Left</Text>
  <Text>Right</Text>
</Flex>
```

## Scrollable Content Patterns

```typescript
import { ScrollBox } from 'ink-scrollbox';

// Bounded scroll area
<ScrollBox height={10} width={50}>
  {longContent}
</ScrollBox>
```

## Input Patterns

```typescript
import { TextInput } from 'ink-text-input';
import { Select } from 'ink-select-input';
import { MultiSelect } from 'ink-multi-select';

// Text input with focus management
const [value, setValue] = useState('');
const inputRef = useRef<HTMLInputElement>(null>;

<TextInput
  ref={inputRef}
  value={value}
  onChange={setValue}
  placeholder="Type something..."
  onSubmit={handleSubmit}
/>

// Select input
<Select
  items={[
    { label: 'Option 1', value: '1' },
    { label: 'Option 2', value: '2' },
  ]}
  onSelect={handleSelect}
/>

// Multi-select
<MultiSelect
  options={[
    { label: 'Choice A', value: 'a' },
    { label: 'Choice B', value: 'b' },
  ]}
  onChange={setSelected}
/>
```

## Gradient Text Pattern

```typescript
import { GradientText } from 'ink-gradient';

// For headers only (per CLI-X 2026 rules)
<GradientText colors={['#6B50FF', '#8B75FF']}>
  <Text bold>FLOYD CLI</Text>
</GradientText>
```

## Spinner Pattern

```typescript
import { Spinner } from 'ink-spinner';

<Spinner type="dots" />
<Text> Processing...</Text>
```

## Syntax Highlighting Pattern

```typescript
import { SyntaxHighlight } from 'ink-syntax-highlight';

<SyntaxHighlight code={codeBlock} language="typescript" />
```

## Progress Indicator Pattern

```typescript
import { ProgressBar } from 'ink-progress-bar';

<ProgressBar percent={progress} />
```
