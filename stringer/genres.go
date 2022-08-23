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

//go:generate stringer -type=Genre
func main() {
	fmt.Printf("Let's watch a %v movie today!", horror)
}

// String returns the iota's string representation
func (g Genre) String() string {
	switch g {
	case horror:
		return "horror"
	case comedy:
		return "comedy"
	default:
		return ""
	}
}
