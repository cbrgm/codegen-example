// Code generated by "stringer -type=Genre"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[horror-0]
	_ = x[fantasy-1]
	_ = x[crime-2]
	_ = x[drama-3]
	_ = x[comedy-4]
}

const _Genre_name = "horrorfantasycrimedramacomedy"

var _Genre_index = [...]uint8{0, 6, 13, 18, 23, 29}

func (i Genre) String() string {
	if i < 0 || i >= Genre(len(_Genre_index)-1) {
		return "Genre(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Genre_name[_Genre_index[i]:_Genre_index[i+1]]
}
