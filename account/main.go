// 내가 작성할 패키지의 이름을 작성
// main은 오로지 컴파일을 위해 필요한 것
package main

import (
	"fmt"

	accounts "github.com/talli0505/learngo/account/banking"
)

// owner , balance 첫 글자 대문자 아니면 private으로 인식
func main() {
	// account := banking.Account{Owner: "James", Balance: 1000}
	// fmt.Println(account)
	account := accounts.NewAccount("James")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		// 프로그램 종료를 시킴
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}
