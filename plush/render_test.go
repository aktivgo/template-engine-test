package main

import (
	"io"
	"os"
	"testing"
)

func BenchmarkPlushRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in, err := os.Open("./source/test.plush")
		if err != nil {
			panic(err)
		}

		render(in, io.Discard)
	}
}
