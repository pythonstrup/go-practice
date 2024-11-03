package main

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%5d, %5d\n", a, b)   // 5칸에 맞춰서 출력하라.
	fmt.Printf("%05d, %05d\n", a, b) // 5칸에 맞춰서 출력하되, 빈 공간을 0으로 채운다.
	fmt.Printf("%-5d, %-05d\n", a, b)
	// %-5d: 5칸을 확보하고 빈 칸이 있다면 오른쪽에 공간을 둠.
	// %-05d: 위와 동일

	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-5d, %-05d\n", c, c)
}
