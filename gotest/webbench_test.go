package gotest

import "testing"

func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(4, 5)
	}
}

func Benchmark_Accumulate(b *testing.B) {
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		accumulate(1000)
	}
}
