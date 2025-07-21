# go_sorting - Sorting Algorithm Benchmarking Suite

A comprehensive Go-based benchmarking suite for comparing the performance of various sorting algorithms across different data distributions and input sizes.

## Overview

go_sorting is a performance evaluation tool that benchmarks multiple sorting algorithms using configurable data generators and input sizes. It provides detailed timing measurements, memory usage statistics, and generates formatted reports for analysis.

## Features

### Sorting Algorithms

The project implements and benchmarks the following sorting algorithms:

- **Insertion Sort** - Simple comparison-based algorithm, efficient for small datasets
- **Radix Sort** - Non-comparative integer sorting algorithm
- **American Flag Sort** - In-place variant of radix sort
- **American Flag Sort (Parallel)** - Multi-threaded version of American Flag Sort
- **Standard Sort** - Go's built-in `sort.Slice()` implementation

### Data Generators

Multiple data distribution patterns are available for testing:

- **Random** - Randomly distributed integers
- **Sorted** - Already sorted ascending data
- **Reversed** - Reverse-sorted (descending) data
- **SortedDoubled** - Two concatenated sorted sequences
- **NearSorted** - Mostly sorted data with some random swaps

## Installation

### Prerequisites

- Go 1.21.0 or higher

### Setup

1. Clone the repository:
```bash
git clone https://github.com/maxdolliger/go_sorting.git
cd go_sorting
```

2. Build the project:
```bash
go build
```

## Usage

### Basic Execution

Run the benchmarking suite with default settings:

```bash
./go_sorting
```

### Configuration

The main configuration is in `main.go`:

```go
sortingSliceSizes := []int{10, 100, 1000, 10_000, 100_000, 1_000_000}
runsPerSize := 1
```

- `sortingSliceSizes`: Array sizes to test
- `runsPerSize`: Number of iterations per size for averaging

### Output

The program generates:

1. **Console output**: Formatted table showing performance results
2. **File output**: Timestamped results file (e.g., `Random_2024-01-15_14-30-45.txt`)

## Benchmarking

### Test Sizes

Default test sizes range from small (10 elements) to large (1,000,000 elements):
- 10
- 100  
- 1,000
- 10,000
- 100,000
- 1,000,000

### Metrics Collected

For each algorithm and input size:
- **Execution Time**: Precise timing using `time.Since()`
- **Memory Usage**: Memory snapshots during execution
- **Correctness Verification**: Ensures data is properly sorted
- **Garbage Collection**: Forced GC between runs for consistent measurements

## Project Structure

```
go_sorting/
├── main.go                    # Entry point and orchestration
├── executer.go               # Benchmark execution engine
├── go.mod                    # Go module definition
├── LICENSE.md                # Project license
├── benchmarks/               # Benchmark test files
│   ├── bench_key_test.go
│   ├── bench_less_test.go
│   └── bench_sort_test.go
├── data/                     # Data generation and evaluation
│   ├── eval.go              # Performance evaluation logic
│   ├── format.go            # Output formatting
│   ├── generator.go         # Test data generators
│   └── table.go             # Table formatting utilities
└── sorting/                  # Sorting algorithm implementations
    ├── types.go             # Common interfaces and types
    ├── insertionSort.go     # Insertion sort implementation
    ├── radixSort.go         # Radix sort implementation
    ├── americanFlagSort.go  # American flag sort implementation
    ├── americanFlagSortP.go # Parallel American flag sort
    └── stdSort.go           # Standard library wrapper
```

## Extending the Project

### Adding New Sorting Algorithms

1. Implement the algorithm in the `sorting/` directory
2. Follow the `SortingFn[T sorting.Sortable]` function signature
3. Add the algorithm to the benchmark suite in `main.go`

### Adding New Data Generators

1. Create generator functions in `data/generator.go`
2. Implement the `DataGenerator[T sorting.Sortable]` interface
3. Return slices of `SortableNumber` or custom `Sortable` types

### Custom Data Types

Implement the `Sortable` interface for custom data types:

```go
type Sortable interface {
    SortValue() int64
}
```

## Performance Considerations

- **Memory Management**: Garbage collection is forced between runs
- **Verification**: Each run includes correctness checking
- **Isolation**: Each algorithm runs independently
- **Reproducibility**: Consistent random seeding for repeatable results

## Contributing

1. Fork the repository
2. Create a feature branch
3. Implement your changes
4. Add appropriate tests
5. Submit a pull request

## License

This project is licensed under the GLWTS (Good Luck With That Shit) Public License. See `LICENSE.md` for details.

## Author

Created by [@maxdolliger](https://github.com/maxdolliger)

---

*Note: This benchmarking suite is designed for educational and research purposes. Results may vary based on hardware, Go version, and system load.* 