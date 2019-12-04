package rgc

import (
	"testing"
)

func BenchmarkGetHubenyDistance(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetHubenyDistance(35.648971, 139.743023, 35.648065, 139.741654)
	}
}

func BenchmarkPythagoreanDistance(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetPythagoreanDistance(35.648971, 139.743023, 35.648065, 139.741654)
	}
}
