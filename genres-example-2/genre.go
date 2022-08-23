package main

import (
	"fmt"
)

type Genre int

const (
	horror Genre = iota
	fantasy
	crime
	drama
	comedy
)

//go:generate ./generator Genre
func main() {
	fmt.Printf("Let's watch a %v movie today!", horror)
}
