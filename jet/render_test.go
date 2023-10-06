package main

import (
	"io"
	"testing"
)

func BenchmarkJetRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		render(io.Discard)
	}
}
