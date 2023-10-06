package main

import (
	"io"
	"os"
	"reflect"

	"github.com/CloudyKit/jet/v6"
)

func main() {
	render(os.Stdout)
}

func render(out io.Writer) {
	set := jet.NewSet(
		jet.NewInMemLoader(),
	)

	set.AddGlobalFunc("concat", func(args jet.Arguments) reflect.Value {
		args.RequireNumOfArguments("concat", 2, 2)
		return reflect.ValueOf(
			args.Get(0).String() + args.Get(1).String(),
		)
	})

	template, err := set.Parse("test", "{{ s := slice(\"foo\", \"bar\", \"asd\") }}\n{{ concat(s[0], s[1]) }}")
	if err != nil {
		panic(err)
	}

	if err := template.Execute(out, nil, nil); err != nil {
		panic(err)
	}
}
