# stringer example

Basic example of Go's `stringer` tool

> Stringer is a tool to automate the creation of methods that satisfy the fmt.Stringer interface. Given the name of a (signed or unsigned) integer type T that has constants defined, stringer will create a new self-contained Go source file implementing: func (t T) String() string

See: https://pkg.go.dev/cmd/go/internal/generate

```
go install golang.org/x/tools/cmd/stringer 
```
