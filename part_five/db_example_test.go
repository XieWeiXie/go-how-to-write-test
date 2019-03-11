package part_five

import "testing"

func BenchmarkInsertByFirstOrInit(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertByFirstOrInit(uint(i))
	}
}

func BenchmarkInsertByFirstOrCreate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertByFirstOrCreate(uint(i))
	}
}

//func TestInsertByFirstOrInit(t *testing.T) {
//	InsertByFirstOrInit(2)
//}
