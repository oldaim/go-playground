package main

import . "fmt"

func main() {

	//배열 선언
	var numbers = [5]int{1, 2, 3, 4, 5}
	//배열 선언 (자동으로 크기 조정)
	numbers2 := [...]int{1, 2, 3, 4, 5}

	Println(numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])

	Println(numbers2[0], numbers2[1], numbers2[2], numbers2[3], numbers2[4])

	var scores = []int{10, 20, 30, 40, 50}
	scores2 := make([]int, 5) // new int[5] 와 유사

	scores2[0] = 10
	scores2[6] = 20               // 인덱스 범위를 벗어나면 에러 발생
	scores2 = append(scores2, 20) // append() 함수를 사용하여 요소 추가 add 와 유사

	Println(scores[0], scores[1], scores[2], scores[3], scores[4])
	Println(scores2[0], scores2[1], scores2[2], scores2[3], scores2[4], scores2[5])

	scores3 := scores2[0:3] // 0 ~ 2 까지의 요소를 가지는 배열을 생성

	Println(scores3[0], scores3[1], scores3[2])
}
