package main

import "fmt"

func main() {
	// 만약 크기를 정한다면 [] 안에 크기를 넣는다
	// 크기가 있는 경우 변수[몇번째 자리] = 넣을값 해서 추가
	// [] 크기가 없으면 slice인데 이때는 추가할때 append를 이용
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn")
	fmt.Println(names)
}
