package main

import . "fmt"

func main() {

	//배열 선언
	var numbers = [5]int{1, 2, 3, 4, 5}
	//배열 선언 (자동으로 크기 조정)
	numbers2 := [...]int{1, 2, 3, 4, 5}

	Println(numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])

	Println(numbers2[0], numbers2[1], numbers2[2], numbers2[3], numbers2[4])

}
