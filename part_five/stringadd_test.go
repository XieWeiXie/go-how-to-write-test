package part_five

import "testing"

func BenchmarkOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		One()
	}
}
