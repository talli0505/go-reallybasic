package main

import (
	"fmt"

	"github.com/talli0505/learngo/dictionary/dic"
)

func main() {
	// dictionary := dic.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }
	dictionary := dic.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition: ", hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
