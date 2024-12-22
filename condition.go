package main

import "fmt"

func main() {

	// 기본 반복문
	// 조건 괄호는 필수 x, 중괄호 (조건이 참일때) 필수

	conditions := [5]int{1, 2, 3, 4, 5}

	for _, el := range conditions {

		if el > 3 {
			fmt.Println("조건은 참 입니다.")
		} else {
			fmt.Println("조건은 거짓 입니다.")
		}
	}

	// 조건문에서 초기화?

	if f := False(); f == false {
		fmt.Println("조건문에서 초기화 후 분기 실행")
	}

	// switch 문
	days := [8]string{"mon", "tue", "wed", "thu", "fri", "sat", "sun", "wrong"}

	for _, el := range days {
		switch el {
		case "mon", "tue", "wed", "thu", "fri":
			fmt.Println("평일 입니다.")
		case "sat", "sun":
			fmt.Println("주말 입니다.")
		default:
			fmt.Println("잘못 입력 되었습니다.")

		}
	}

	// switch 문 fallthrough?
	daysFall := [8]string{"mon", "tue", "wed", "thu", "fri", "sat", "sun", "wrong"}

	for _, el := range daysFall {
		switch el {
		case "mon", "tue", "wed", "thu", "fri":
			fmt.Println("평일 입니다.")
			fallthrough //조건 검사 없이 다음 구문을 실행 시키는 제어문 (java 에서 switch 에 break 없이 사용하는것과 동일)
		case "sat", "sun":
			fmt.Println("주말 입니다.")
			fallthrough
		default:
			fmt.Println("잘못 입력 되었습니다.")
		}
	}
}

func False() bool {
	return false
}
