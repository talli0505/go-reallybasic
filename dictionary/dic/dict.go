package dic

import "errors"

// Dictionary type
// Dictionary는 struct가 아니고 alias(별명, 가명)임
// ex ) type Money int 라고 만들면 Money(1) -> 1 과 같은 뜻
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	// 에러처리를 먼저 한다
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}
