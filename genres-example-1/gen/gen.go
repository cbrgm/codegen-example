package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"html/template"
	"os"
	"sort"
	"strings"
)

func main() {
	// the name of the type we're looking for
	typeName := os.Args[1]
	// constants definitions
	constantNames := []string{}

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

	targetType := pkg.Types.Scope().Lookup(typeName)
	if targetType == nil {
		panic(fmt.Errorf("failed to find type declaration for %v", typeName))
	}

	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			gd, ok := decl.(*ast.GenDecl) // keywords type, const or var
			if !ok {
				continue
			}
			if gd.Tok != token.CONST {
				continue
			}
			for _, spec := range gd.Specs {
				spec := spec.(*ast.ValueSpec)
				for _, name := range spec.Names {
					if pkg.TypesInfo.Defs[name].Type() == targetType.Type() {
						constantNames = append(constantNames, name.Name)
					}
				}
			}
		}
		sort.Strings(constantNames)
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

		outputFileName := fmt.Sprintf("gen_%v_string.go", strings.ToLower(typeName))
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			panic(err)
		}
		tmpl := template.Must(template.New("out").Parse(outputTmpl))
		if err := tmpl.Execute(outputFile, struct {
			PackageName       string
			Type              string
			DeclaredConstants []string
		}{
			PackageName:       pkg.Name,
			Type:              typeName,
			DeclaredConstants: constantNames,
		}); err != nil {
			panic(err)
		}

		if err := outputFile.Close(); err != nil {
			panic(err)
		}
	}
}
