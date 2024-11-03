package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}

	// 만약 위에서 문자열을 입력할 경우. 버퍼를 지워주는 과정이 없기 때문에 아래에서도 에러가 터져버림.
	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
