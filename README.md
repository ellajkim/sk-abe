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
BenchmarkABEScheme/Vector_Size_5-8         	  611896	      1655 ns/op
BenchmarkABEScheme/Vector_Size_10-8        	  442987	      2832 ns/op
BenchmarkABEScheme/Vector_Size_20-8        	  243865	      5559 ns/op
BenchmarkABEScheme/Vector_Size_50-8        	   87523	     13784 ns/op
BenchmarkABEScheme/Vector_Size_100-8       	   47011	     26989 ns/op
BenchmarkABEScheme/Vector_Size_500-8       	    7552	    141130 ns/op
BenchmarkABEScheme/Vector_Size_1000-8      	    4897	    249314 ns/op
BenchmarkABEScheme/Vector_Size_2500-8      	    1753	    663087 ns/op
BenchmarkABEScheme/Vector_Size_10000-8     	     372	   3064785 ns/op
BenchmarkABEScheme/Vector_Size_100000-8    	      40	  27116415 ns/op
```

- `Vector_Size_X`: Indicates the vector length used in the IP-SK-ABE evaluation.
- `-Y`: Indicates number of benchmark iterations.
- `Z ns/op`: Represents the average time in nanoseconds it takes for a single IP-SK-ABE evaluation operation.

For example, `BenchmarkABEScheme/Vector_Size_1000-8 4897	249314 ns/op` means:
- Vector length: 1000
- Benchmark iterations: 8
- Average time per evaluation: 249,314 nanoseconds (about 0.25 milliseconds)

## Acknowledgements
We use Sacha Servan-Schreiber's [random oracle based CPRF](https://github.com/sachaservan/cprf) in our implementation.


## ⚠️ Important Warning

**This implementation is intended for _research purposes only_. The code has NOT been vetted by security experts. As such, no portion of the code should be used in any real-world or production setting!**
