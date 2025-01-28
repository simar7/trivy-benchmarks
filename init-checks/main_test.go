package main

import "testing"

func Benchmark_InitLoop(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		InitLoop()
	}
}
