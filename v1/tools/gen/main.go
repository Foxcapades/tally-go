// build ignore

package main

import (
	"fmt"
	"os"
	"text/template"
)

const (
	ver = "v1"
	tools = ver + "/tools/gen"
	tally = ver + "/tally"

	tpl  = tools + "/*.gotpl"
	outImpl = tally + "/%s.go"
	outTest = tally + "/%s_test.go"
)

var types = [...]Data{
	{"ITally", "int", "ints", []string{"", "8", "16", "32", "64"}},
	{"UTally", "uint", "uints", []string{"", "8", "16", "32", "64"}},
}

func main() {
	tpl := template.Must(template.ParseGlob(tpl))

	for i := range types {
		render(tpl, &types[i])
	}
}

func render(t *template.Template, d *Data) {
	file, err := os.Create(fmt.Sprintf(outImpl, d.File))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err = t.ExecuteTemplate(file, "tally", d); err != nil {
		panic(err)
	}

	test, err := os.Create(fmt.Sprintf(outTest, d.File))
	if err != nil {
		panic(err)
	}
	defer test.Close()

	if err = t.ExecuteTemplate(test, "test", d); err != nil {
		panic(err)
	}
}

type Data struct {
	Name string
	Type string
	File string
	Suffixes []string
}
