package main

import "fmt"

func main() {

	// go 에서는 반복문 for 형태만 있음
	// 반복문 기본형
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	//while 문 형태
	sum := 1
	for sum < 100000000000000000 {
		sum += sum
	}

	fmt.Printf("%d\n", sum)

	// 무한 루프 형태
	count := 1

	for {
		fmt.Printf("%d Still Running\n", count)
		count++
		if count > 10 {
			break
		}
	}

	// 슬라이스 / 배열 순회
	intArray := [5]int{5, 4, 3, 2, 1}

	for index, value := range intArray {
		fmt.Printf("%d 번 index의 값 %d\n", index, value)
	}

	// 컬랙션 순회
	ages := map[string]int{
		"철수": 20,
		"영희": 22,
	}

	for key, value := range ages {
		fmt.Printf("%s의 나이: %d\n", key, value)
	}
}
