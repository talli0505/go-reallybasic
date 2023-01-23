package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

// Go는 constructor method가 없다
func main() {
	favFood := []string{"kimchi", "ramen"}
	James := person{name: "James", age: 18, favFood: favFood}
	fmt.Println(James.name)
}
