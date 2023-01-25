package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
// *Account는 Account를 보고있다는 뜻
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// 이렇게 하면 main에서 account를 이용하여 Deposit을 사용 할 수 있음
// a Account 부분을 receiver라고 함
// a는 원하는 이름 Account는 사용할 타입
// receiver( == method ) 쓸때 주의사항은 사용하고자 하는 struct의 알파벳 소문자를 따오는 것
// 근데 a Acoount로 하면 값을 복사해오기때문에 main에서 값을 넣어도 인식 못함 -> a *Account로 바꿈
// a *Account로 바꾸면 recevier나 account를 복사하지말고 실제 receiver를 달라고 하는것 ( pointer receiver )
func (a *Account) Deposit(amount int) {
	fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

// Balacne of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
// go는 try catch가 없기 때문에 직접 error 넣어줘야함
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw you are poor")
	}
	a.balance -= amount
	return nil
}
