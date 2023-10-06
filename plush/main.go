package main

import (
	"github.com/gobuffalo/plush"
	"io"
	"os"
)

func main() {
	in, err := os.Open("./plush/source/test.plush")
	if err != nil {
		panic(err)
	}

	render(in, os.Stdout)
}

func render(input *os.File, output io.Writer) {
	ctx := plush.NewContext()

	ctx.Set("concat", func(a string, b string) string {
		return a + b
	})

	res, err := plush.RenderR(input, ctx)
	if err != nil {
		panic(err)
	}

	output.Write([]byte(res))
}
