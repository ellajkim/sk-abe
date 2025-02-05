# Inner-Product Secret-Key ABE

This repository contains an implementation of our secret-key attribute based encryption (SK-ABE) scheme using inner products. We use two tools to instantiate our scheme: constrained pseudorandom functions (CPRFs) and symmetric-key encryption (SKE).

## Code Organization

| Directory | Description |
| :--- | :--- |
| [abe/](abe/) | IP-SK-ABE construction |
| [ro-cprf/](ro-cprf/) | Random oracle based CPRF construction |
| [ske/](ske/) | AES-GCM SKE construction |

## Prerequisites

- Go (version 1.20 or later)

## Running Benchmarks

To run benchmarks for each implementation:

1. Select the implementation:
   ```
   cd abe
   ```

2. Run the benchmarks:
   ```
   go test -bench=.
   ```

## Interpreting the Results

The benchmark results are presented in the following format:
```
BenchmarkEval/length=10-10                555568              2150 ns/op
BenchmarkEval/length=50-10                118142              9941 ns/op
BenchmarkEval/length=100-10                59997             19887 ns/op
BenchmarkEval/length=500-10                12088             97314 ns/op
BenchmarkEval/length=1000-10                5764            196743 ns/op
```

- `length=X`: Indicates the vector length used in the CPRF evaluation.
- `-Y`: Indicates number of benchmark iterations.
- `Z ns/op`: Represents the average time in nanoseconds it takes for a single CPRF evaluation operation.

For example, `BenchmarkEval/length=100-10 60027 19896 ns/op` means:
- Vector length: 100
- Benchmark iterations: 10
- Average time per evaluation: 19,896 nanoseconds (about 0.02 milliseconds)

## Acknowledgements
We use Sacha Servan-Schreiber's [random oracle based CPRF](https://github.com/sachaservan/cprf) in our implementation.


## ⚠️ Important Warning

**This implementation is intended for _research purposes only_. The code has NOT been vetted by security experts. As such, no portion of the code should be used in any real-world or production setting!**
