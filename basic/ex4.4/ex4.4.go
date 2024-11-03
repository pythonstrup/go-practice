package main

import "fmt"

// 타입 변환에 관하여
func main() {
	a := 3 // int (64 bit computer) = 사실상 int64
	var b float64 = 3.5

	// var c int = b // 타입이 달라 에러
	var c int = int(b)

	// d := a * b // 마찬가지로 a와 b의 타입이 달라 에러
	d := float64(a) * b

	var e int64 = 7
	// go는 강타입 언어라 사실상 같은 타입으로 보이더라도(int, int64) 타입이 다르면 컴파일 에러가 발생
	// f := a * e
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)
}
