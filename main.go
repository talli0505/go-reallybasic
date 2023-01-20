// 내가 작설할 패키지의 이름을 작성
// main은 오로지 컴파일을 위해 필요한 것
package main

import (
	"fmt"
	"strings"
)

// func multiply(a, b int) int {
// 	return a * b
// }

// // go는 기본적으로 func이 있어야 작동
// func main() {
// 	//fmt 부분을 지우고 저장하면 import "fmt"도 자동적으로 사라짐 (vscode에서만 작동)
// 	// fmt : formatting을 위한 package
// 	// go 같은경우 function 대문자로 시작
// 	// fmt.Println("Hello World!")

// 	// const면 똑같이 불변, var면 변경 가능
// 	const name string = "James"
// 	fmt.Println(name)

// 	// 아래 두개는 같은 뜻
// 	// var name2 string = "Lynn"
// 	name2 := "Sam"
// 	fmt.Println(name2)

// 	fmt.Println(multiply(2, 2))
// }

// lenAndUpper는 return에 초기화한거, lenAndUpper2는 위에 초기화하고 대입한것
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 이런형식을 naked return이라고 한다. return값에 variable을 명시 x
func lenAndUpper2(name string) (length int, uppercase string) {
	// defer는 return 이후에 작동하는 형태
	// 여기서는 main에서 lendAndUpper2를 사용하고 나서 끝난 후 defer가 나오고 그다음에 결과 출력
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	name := "nico"
	name = "lynn"
	fmt.Println(name)

	// // 두개 다 선언해서 말하고싶을때
	// totalLenght, upperName := lenAndUpper("Jame")
	// fmt.Println(totalLenght, upperName)

	// 둘 중 하나만 얘기하고싶을때는 _ 로 ignore를 시킴
	totalLenght, _ := lenAndUpper("Jame")
	fmt.Println(totalLenght)

	totalLength, uppercase := lenAndUpper2("Smith")
	fmt.Println(totalLength, uppercase)

	repeatMe("A", "B", "C", "D", "E")
}
