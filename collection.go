package main

import "strconv"

func main() {
	// Slice
	// Slice 선언
	numbers := []int{90, 80, 110}

	for i := range numbers {
		print(strconv.Itoa(i) + " ")
	}

	// make 를 활용한 선언

	numbers = make([]int, 5)     //길이가 5인 슬라이스
	numbers = make([]int, 5, 10) //길이가 5 용량인 10 인 슬라이스
}
