package main

import (
	"io"
	"os"

	"github.com/gobuffalo/plush"
)

func main() {
	render(os.Stdout)
}

func render(output io.Writer) {
	ctx := plush.NewContext()

	ctx.Set("concat", func(a string, b string) string {
		return a + b
	})

	res, err := plush.Render("<% let s = [\"foo\", \"bar\", \"asd\"] %>\n<%= concat(s[0], s[1]) %>", ctx)
	if err != nil {
		panic(err)
	}

	output.Write([]byte(res))
}
