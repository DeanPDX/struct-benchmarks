# Struct Benchmarks
This is a contrived benchmark to illustrate differences between iterating over an array of type `[]T` and `[]*T` and how much it matters if we `range` over them versus index into them. See [this thread](https://forum.golangbridge.org/t/can-someone-kindly-eli5-the-difference-between-using-slice-of-pointers-and-slice-of-structs/34905) for more details.

## Running the benchmarks
After cloning the repository, `go test -bench .` should yield something like this:

```bash
goos: windows
goarch: amd64
pkg: github.com/DeanPDX/struct-benchmarks
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkRangePointers-8          654621              1839 ns/op
BenchmarkRangeStructs-8           499588              2332 ns/op
BenchmarkIndexPointers-8          704833              1732 ns/op
BenchmarkIndexStructs-8           666658              1844 ns/op
PASS
ok      github.com/DeanPDX/struct-benchmarks    5.002s
```