# Go Profiling

This repository demonstrates how to use Go's built-in profiling tools to analyze program performance.

## Prerequisites

- Go 1.16 or later
- graphviz (for viewing profile graphs)

## Running Profiles

### CPU Profile

```bash
go test -cpuprofile cpu.prof -bench .
```

### Memory Profile

```bash
go test -memprofile mem.prof -bench .
```

### Goroutine Profile

```bash
go test -blockprofile block.prof -bench .
```

## Analyzing Profiles

Use `go tool pprof` to analyze the generated profiles:

```bash
# CPU profile analysis
go tool pprof cpu.prof

# Memory profile analysis
go tool pprof mem.prof

# Block profile analysis
go tool pprof block.prof
```

### Common pprof commands

Once in the pprof interactive mode:
- `top`: Show top consumers
- `web`: Generate a graph visualization (requires graphviz)
- `list <function-name>`: Show annotated source for a function
- `quit`: Exit pprof

## Visualization

To generate a visual graph of the profile:

```bash
go tool pprof -http=:6060 cpu.prof
```

## Performance Results
### Comparison Summary
The following table compares Push Time and Pop Time for ArrayStack and LinkedListStack after profiling, including the impact of memory escape and performance ratios.

| Operation   | ArrayStack (After)      | LinkedListStack (After)   | Escape Difference            | Ratio (LL/Array)         |
|-------------|-------------------------|---------------------------|------------------------------|--------------------------|
| **Push Time** | 59 µs (5.9 ns/push)     | 336.417 µs (33.6417 ns/push) | No escape (stack) vs. escape (heap) | Ratio (LL/Array) ~5.7x slower |
| **Pop Time**  | 9.541 µs (1.9082 ns/pop) | 13.208 µs (2.6416 ns/pop)  | No new escape vs. prior heap | Ratio (LL/Array) ~1.4x slower |


### Full Stack Comparison Table
The complete profiling data, including memory allocation and GC metrics, is shown below:

| Metric                | ArrayStack (Before) | ArrayStack (After) | LinkedListStack (Before) | LinkedListStack (After) |
|-----------------------|---------------------|--------------------|--------------------------|-------------------------|
| Alloc                 | 227 KB              | 227 KB             | 227 KB                   | 386 KB                  |
| HeapSys               | 3808 KB             | 3808 KB            | 3808 KB                  | 3776 KB                 |
| GC CPU Fraction       | 0.00 %              | 0.00 %             | 0.00 %                   | 0.00 %                  |
| Push Time (10,000)    | -                   | 59 µs              | -                        | 336.417 µs              |
| Pop Time (5,000)      | -                   | 9.541 µs           | -                        | 13.208 µs               |
| After GC Alloc        | -                   | -                  | -                        | 163 KB                  |
| After GC HeapSys      | -                   | -                  | -                        | 3712 KB                 |
| After GC CPU Fraction | -                   | -                  | -                        | 2.94 %                  |

### Calculations

- **Push Time**:
    - **ArrayStack**: 59 µs / 10,000 = 5.9 ns per push.
    - **LinkedListStack**: 336.417 µs / 10,000 = 33.6417 ns per push.
    - **Ratio (LL/Array)**: 33.6417 / 5.9 ≈ 5.7x slower.

- **Pop Time**:
    - **ArrayStack**: 9.541 µs / 5,000 = 1.9082 ns per pop.
    - **LinkedListStack**: 13.208 µs / 5,000 = 2.6416 ns per pop.
    - **Ratio (LL/Array)**: 2.6416 / 1.9082 ≈ 1.4x slower.