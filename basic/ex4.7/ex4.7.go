package main

import "fmt"

func main() {
	var a float32 = 1234.623
	var b float32 = 3456.123
	var c float32 = a * b // 약 0.3의 오차 발생
	var d float32 = c * 3 // * 3을 했기 때문에 약 1의 오차가 생김

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
