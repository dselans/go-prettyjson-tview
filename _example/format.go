package main

import (
	"fmt"

	"github.com/dselans/go-prettyjson-tview"
)

func main() {
	s, _ := prettyjson.Format([]byte(`{"foo":"bar"}`))
	fmt.Println(string(s))
}
