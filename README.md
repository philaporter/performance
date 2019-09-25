### `performance` 
this repo provides terrible code samples for demoing go performance profiling.

#### Useful commands

Running benchmarks:
```
go test -run=XXX -bench=. -benchtime=20x -cpu=1,2,4 -cpuprofile=c.p .
go test -run=XXX -bench=. -cpu=1,2,4 -memprofile=m.p .
go test -run=XXX -bench=. -cpu=1,2,4 -blockprofile=b.p .
```

Capturing a profile on a running application:
```
curl "http://127.0.0.1:8085/debug/pprof/profile?seconds=30" > cpu.pprof
```

Analyzing the profile via command line:
```
go tool pprof cpu.profile
```

Analyzing the profile via web tool:
```
go tool pprof -http=:8089 cpu.profile
```