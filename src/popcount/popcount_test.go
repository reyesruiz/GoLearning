package popcount_test

import (
	//"popcount"
	"github.com/reyesruiz/GoLearning/src/popcount"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}
