package main

import "fmt"

// loop은 오직 for로만 가능 (foreach, for of, map 등 없음)
func superAdd(numbers ...int) int {
	// range는 array에 loop를 적용시켜줌

	// 이렇게 하면 이제 index를 가져와서 0부터 가져옴
	// 이거 밑에꺼 index붙여있는걸로 확인 가능
	// for number := range numbers {
	// 	fmt.Println(number)
	// }

	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
