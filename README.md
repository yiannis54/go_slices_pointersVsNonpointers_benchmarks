# go_slices_pointersVsNonpointers_benchmarks
Benchmark tests to compare slice of pointers to slice of non pointers.

## Run

```
go test -bench . -count 10
```

## Results

```
cpu: Intel(R) Core(TM) i7-4720HQ CPU @ 2.60GHz
BenchmarkSlicePointers-8        1000000000               0.0003097 ns/op               0 B/op          0 allocs/op
BenchmarkSlicePointers-8        1000000000               0.0003177 ns/op               0 B/op          0 allocs/op
BenchmarkSlicePointers-8        1000000000               0.0002143 ns/op               0 B/op          0 allocs/op
BenchmarkSlicePointers-8        1000000000               0.0002816 ns/op               0 B/op          0 allocs/op
BenchmarkSlicePointers-8        1000000000               0.0003322 ns/op               0 B/op          0 allocs/op
BenchmarkSlicePointers-8        1000000000               0.0002001 ns/op               0 B/op          0 allocs/op
BenchmarkSlicePointers-8        1000000000               0.0002103 ns/op               0 B/op          0 allocs/op
BenchmarkSlicePointers-8        1000000000               0.0002009 ns/op               0 B/op          0 allocs/op
BenchmarkSlicePointers-8        1000000000               0.0003416 ns/op               0 B/op          0 allocs/op
BenchmarkSlicePointers-8        1000000000               0.0002024 ns/op               0 B/op          0 allocs/op
BenchmarkSliceNoPointers-8      1000000000               0.0002621 ns/op               0 B/op          0 allocs/op
BenchmarkSliceNoPointers-8      1000000000               0.0001408 ns/op               0 B/op          0 allocs/op
BenchmarkSliceNoPointers-8      1000000000               0.0001491 ns/op               0 B/op          0 allocs/op
BenchmarkSliceNoPointers-8      1000000000               0.0001426 ns/op               0 B/op          0 allocs/op
BenchmarkSliceNoPointers-8      1000000000               0.0001803 ns/op               0 B/op          0 allocs/op
BenchmarkSliceNoPointers-8      1000000000               0.0001428 ns/op               0 B/op          0 allocs/op
BenchmarkSliceNoPointers-8      1000000000               0.0001767 ns/op               0 B/op          0 allocs/op
BenchmarkSliceNoPointers-8      1000000000               0.0001380 ns/op               0 B/op          0 allocs/op
BenchmarkSliceNoPointers-8      1000000000               0.0001526 ns/op               0 B/op          0 allocs/op
BenchmarkSliceNoPointers-8      1000000000               0.0001459 ns/op               0 B/op          0 allocs/op
PASS
```
