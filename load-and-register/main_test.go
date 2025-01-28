package main

import "testing"

func Benchmark_LoadRegisterLoop(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		LoadRegisterLoop()
	}
}
