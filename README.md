# Benchmarks

Here you can see benchmarks showing speed differences of equivalent operations implemented using KVS, Badgerhold and SQLite respectfully.
```
go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/tauraamui/kvs-bench
BenchmarkBHLoad-8                                	   44186	     26428 ns/op
BenchmarkBHLoad500Records-8                      	     264	   4526775 ns/op
BenchmarkBHLoad100RecordsQueryColour-8           	    1276	    938912 ns/op
BenchmarkBHLoad500RecordsQuerySizeNoMatches-8    	     259	   4615119 ns/op
BenchmarkBHLoad500RecordsQueryColour-8           	     259	   4618628 ns/op
BenchmarkKVSLoad-8                               	  120360	      9933 ns/op
BenchmarkKVSLoad500Records-8                     	    2215	    540268 ns/op
BenchmarkKVSLoad100RecordsQueryColour-8          	    9252	    139602 ns/op
BenchmarkKVSLoad500RecordsQuerySizeNoMatches-8   	    2235	    539261 ns/op
BenchmarkKVSLoad500RecordsQueryColour-8          	    2480	    484517 ns/op
BenchmarkSLLoad-8                                	   60076	     19706 ns/op
BenchmarkSLLoad500Records-8                      	     802	   1495841 ns/op
BenchmarkSLLoad100RecordsQueryColour-8           	    7113	    166161 ns/op
BenchmarkSLLoad500RecordsQuerySizeNoMatches-8    	   24186	     49457 ns/op
BenchmarkSLLoad500RecordsQueryColour-8           	    7190	    166854 ns/op
PASS
ok  	github.com/tauraamui/kvs-bench	24.693s
```
