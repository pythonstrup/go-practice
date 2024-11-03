package main

import "fmt"

func main() {
	var a int = 3
	var b int
	var c = 4
	d := 5

	fmt.Println(a, b, c, d)

	var e = "Hello"
	f := 3.14
	fmt.Println(e, f)

	// 타입별 기본값
	// 모든 정수 타입 => 0
	// 모든 실수 타입 => 0.0
	// 불리언 => false
	// 문자열 => "" (빈 문자열)
	// 그외 => nil (정의되지 않은 메모리 주소를 나타내는 Go 키워드)
}
