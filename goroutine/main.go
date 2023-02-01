package main

import (
	"fmt"
	"time"
)

// goroutine : 기본적으로 다른 함수와 동시에 실행되는 함수
func main() {
	go sexyCount("James")
	go sexyCount("flynn")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
