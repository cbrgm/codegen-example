package main

import "fmt"

func (v Genre) String() string {
	switch v {
		case comedy:
		return "comedy"
	case crime:
		return "crime"
	case drama:
		return "drama"
	case fantasy:
		return "fantasy"
	case horror:
		return "horror"
	default:
		panic(fmt.Errorf("unknown Genre value %d", v))
	}
}