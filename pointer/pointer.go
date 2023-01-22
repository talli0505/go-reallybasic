package main

import "fmt"

func main() {
	// & 가 붙으면 메모리 주소가 확인 가능
	// * 메모리 주소를 볼수가 있음
	// &* 이 붙으면 그 값을 볼 수 있음
	a := 2
	b := &a
	a = 10
	fmt.Println(a, *b)
}
