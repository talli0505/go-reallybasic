// 조건문
package main

import "fmt"

func canIDrink(age int) bool {
	// variable 사용 ( koreanAge := age +2 )
	// if문 밖에서 쓰면 func 내에서의 지역변수가 될 수 있지만, if 내에서 variable 쓸 경우 if에서만 사용한다는걸 확인 가능
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
