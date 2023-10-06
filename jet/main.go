package main

import (
	"github.com/CloudyKit/jet/v6"
	"io"
	"os"
	"reflect"
)

func main() {
	render(getTemplate("./jet/source", "test.jet"), os.Stdout)
}

func getTemplate(dir string, templatePath string) *jet.Template {
	source := jet.NewSet(
		jet.NewOSFileSystemLoader(dir),
	)

	source.AddGlobalFunc("concat", func(args jet.Arguments) reflect.Value {
		args.RequireNumOfArguments("concat", 2, 2)
		return reflect.ValueOf(
			args.Get(0).String() + args.Get(1).String(),
		)
	})

	template, err := source.GetTemplate(templatePath)
	if err != nil {
		panic(err)
	}

	return template
}

func render(template *jet.Template, out io.Writer) {
	if err := template.Execute(out, nil, nil); err != nil {
		panic(err)
	}
}
