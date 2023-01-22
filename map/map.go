package main

import "fmt"

func main() {
	//map[key의 타입]value의 타입
	Jame := map[string]string{"name": "Jame", "age": "111"}
	for key, value := range Jame {
		fmt.Println(key, value)
	}
}
