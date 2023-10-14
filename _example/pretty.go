package main

import (
	"fmt"

	"github.com/dselans/go-prettyjson-tview"
)

func main() {
	v := map[string]interface{}{
		"str":   "foo",
		"num":   100,
		"bool":  false,
		"null":  nil,
		"array": []string{"foo", "<bar>", "baz"},
		"map": map[string]interface{}{
			"foo": "bar",
		},
	}
	s, _ := prettyjson.Marshal(v)
	fmt.Println(string(s))
}
