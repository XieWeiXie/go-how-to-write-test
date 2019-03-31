package part_six

import "testing"

func BenchmarkOne(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		One()
	}
}

func BenchmarkTwo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Two()
	}
}
