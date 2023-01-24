package banking

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
