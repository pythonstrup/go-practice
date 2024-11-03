package main

import "fmt"

// 패키지 전역 변수
// 해당 패키지 안에서는 다 쓸수 있다.
var g int = 10

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	// 컴파일 에러
	// 변수 s는 block scope 안에 들어가 있고, 중괄호 닫힘 '}'이 되는 순간 변수가 사라져 버린다.
	// m = s * 20
}
