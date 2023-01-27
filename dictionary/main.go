package main

import (
	"fmt"

	"github.com/talli0505/learngo/dictionary/dic"
)

func main() {
	dictionary := dic.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
