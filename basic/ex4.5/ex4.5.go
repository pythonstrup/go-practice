package main

import "fmt"

func main() {
	var a int16 = 3456   // int16 => 2byte 정수
	var b int8 = int8(a) // int8 => 1byte 정수
	// a는 1byte 범위를 벗어나기 때문에 stackoverflow 발생
	// int 16 a => 0000110110000000
	// int 8 a => (앞의 8자리가 날라감) [00001101]10000000 => -127
	// 컴퓨터에서 이진수 맨 앞자리는 부호 비트임. 0인 경우 +를, 1인 경우 -를 뜻함
	var a2 int16 = 3460
	var b2 int8 = int8(a2)

	fmt.Println(a, b)
	fmt.Println(a2, b2)

	// 2진수 실수 표현
	// 32비트 실수 => 1비트(부호비트) + 8비트(지수부) + 23비트(소수부)
}
