package main

import "fmt"

func main() {
	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)

	var b int64 = 394809235823953
	fmt.Println(b)

	// var c uint64 = -4248190418 // error => unsigned int에는 마이너스를 담을 수 없다.
	// var c uint8 = 256 // error => 8비트로는 0~255까지의 숫자만 표현 가능하기에 overflow 에러가 발생한다.
	var c uint16 = 256
	fmt.Println(c)

	var d float32 = 3.2523663
	fmt.Println(d)

	var e byte = 255
	fmt.Println(e)

	// rune - UTF_8(3byte)
	// int32의 별칭(이진수라 2의 배수가 효율적 => rune은 4byte를 쓴다.)
	var f1 rune = 'a'
	fmt.Println(f1)
	var f2 rune = 127
	fmt.Println(f2)

	var boolTrue bool = true
	var boolFalse bool = false
	fmt.Println(boolTrue)
	fmt.Println(boolFalse)

	// 별칭 타입을 선언해 만들 수도 있다.
	// rune => int32
	// byte ==> uint8
	type myInt int
	var number myInt = 125122
	fmt.Println(number)

	// 초기값 생략 가능. 대신 이런 경우 타입은 필수로 들어가야 함.
	var num1 int // 0이라는 초기값이 들어감.
	fmt.Println(num1)
	// 타입 추론 가능
	var num2 = 10
	fmt.Println(num2)
	// 콜론(:)은 var를 내포
	num3 := 20
	fmt.Println(num3)
}
