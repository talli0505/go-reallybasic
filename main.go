// 내가 작설할 패키지의 이름을 작성
// main은 오로지 컴파일을 위해 필요한 것
package main

import (
	"fmt"

	"github.com/talli0505/learngo/banking"
)

// owner , balance 첫 글자 대문자 아니면 private으로 인식
func main() {
	account := banking.Account{Owner: "James", Balance: 1000}
	fmt.Println(account)
}
