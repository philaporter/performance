package main

import (
	i "performance/server/initialize"
	"sort"
	"sync"
	"testing"
)

/*
	go test -run=XXX -bench=. -benchtime=20x -cpu=1,2,4 -cpuprofile=c.p .
	go test -run=XXX -bench=. -cpu=1,2,4 -memprofile=m.p .
	go test -run=XXX -bench=. -cpu=1,2,4 -blockprofile=b.p .
 */

var Result string

func setup() {
	Result = ""
	Response = ""
	i.Match = []string{}
	i.Lookup = sync.Map{}
	i.Init()
	sort.Strings(i.Match)
}

func BenchmarkBuildResponse(b *testing.B) {
	b.StopTimer()
	setup()
	var s string
	b.StartTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		s = buildResponse()
	}
	Result = s
}

func BenchmarkBuildResponseV2(b *testing.B) {
	b.StopTimer()
	setup()
	var s string
	b.StartTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		s = buildResponseV2()
	}
	Result = s
}

func BenchmarkBuildResponseV3(b *testing.B) {
	b.StopTimer()
	setup()
	var s string
	b.StartTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		s = buildResponseV3()
	}
	Result = s
}
