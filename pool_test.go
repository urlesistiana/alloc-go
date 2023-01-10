package alloc

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantCap int
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"", 2, 2},
		{"", 3, 4},
		{"", 128, 128},
		{"", 129, 256},
		{"", 255, 256},
		{"", SmallBufSize - 1, SmallBufSize},
		{"", SmallBufSize, SmallBufSize},
		{"", SmallBufSize + 1, SmallBufSize + SmallBufSize/4},
		{"", SmallBufSize + SmallBufSize/4, SmallBufSize + SmallBufSize/4},
		{"", SmallBufSize + SmallBufSize/4 + 1, SmallBufSize + SmallBufSize/2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Get(tt.size)
			if len(b) != tt.size {
				t.Fatalf("want size %d, got %v", tt.size, len(b))
			}
			if cap(b) != tt.wantCap {
				t.Fatalf("want cap %d, got %v", tt.wantCap, cap(b))
			}
			Release(b)
		})
	}
}

func Benchmark_Pool(b *testing.B) {
	for i := 0; i < 24; i += 4 {
		size := 1 << i
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				b := Get(size)
				Release(b)
			}
		})
	}
}
