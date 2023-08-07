# Benchmarks

goos: linux
goarch: amd64
pkg: github.com/tauraamui/kvs-bench
cpu: Intel(R) Core(TM) i5-8200Y CPU @ 1.30GHz
BenchmarkBHLoad-4                                	   13290	     94342 ns/op
BenchmarkBHLoad500Records-4                      	      72	  14768587 ns/op
BenchmarkBHLoad100RecordsQueryColour-4           	     400	   3086616 ns/op
BenchmarkBHLoad500RecordsQuerySizeNoMatches-4    	      88	  14084911 ns/op
BenchmarkBHLoad500RecordsQueryColour-4           	      73	  17030642 ns/op
BenchmarkKVSLoad-4                               	   57556	     20270 ns/op
BenchmarkKVSLoad500Records-4                     	     494	   2786213 ns/op
BenchmarkKVSLoad100RecordsQueryColour-4          	    1654	    679295 ns/op
BenchmarkKVSLoad500RecordsQuerySizeNoMatches-4   	     429	   3191895 ns/op
BenchmarkKVSLoad500RecordsQueryColour-4          	     496	   2703187 ns/op
BenchmarkSLLoad-4                                	   12842	     91069 ns/op
PASS
ok  	github.com/tauraamui/kvs-bench	53.426s
