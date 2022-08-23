package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/tools/go/packages"
)

func main() {
	typeName := os.Args[1]

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo |
			packages.NeedSyntax | packages.NeedName,
	}
	pkgs, err := packages.Load(cfg)
	if err != nil {
		panic(err)
	}
	if len(pkgs) != 1 {
		panic(fmt.Errorf("got unexpected number of packages %v", len(pkgs)))
	}
	pkg := pkgs[0]
	const outputTmpl = `package {{.PackageName}}
import "fmt"

// THIS FILE IS GENERATED, DO NOT EDIT MANUALLY

func (v {{.Type}}) String() string {
	switch v {
		{{range .DeclaredConstants -}}
	case {{.}}:
		return "{{.}}"
	{{end -}}
	default:
		panic(fmt.Errorf("unknown {{.Type}} value %d", v))
	}
}`

	outputFile := fmt.Sprintf("gen_%v_string.go", strings.ToLower(typeName))
	ExecuteTemplate(outputFile, outputTmpl, struct {
		PackageName       string
		Type              string
		DeclaredConstants []string
	}{
		PackageName:       pkg.Name,
		Type:              typeName,
		DeclaredConstants: os.Args[1:],
	})
}

// ExecuteTemplate renders the named template and writes to io.Writer wr.
func ExecuteTemplate(file string, tmpl string, data interface{}) error {
	wr := os.Stdout
	if output := file; output != "" {
		wri, err := os.Create(output)
		if err != nil {
			return err
		}
		wr = wri
		defer wr.Close()
	}

	buf := new(bytes.Buffer)

	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(buf, "", data)
	if err != nil {
		return err
	}

	src, err := format(buf)
	if err != nil {
		return err
	}
	_, err = io.Copy(wr, src)
	return err
}

// format formats a template using gofmt.
func format(in io.Reader) (io.Reader, error) {
	var out bytes.Buffer

	gofmt := exec.Command("gofmt", "-s")
	gofmt.Stdin = in
	gofmt.Stdout = &out
	gofmt.Stderr = os.Stderr
	err := gofmt.Run()
	return &out, err
}
