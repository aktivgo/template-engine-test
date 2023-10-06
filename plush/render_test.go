package main

import (
	"io"
	"testing"
)

func BenchmarkPlushRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		render(io.Discard)
	}
}
