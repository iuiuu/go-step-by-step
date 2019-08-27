package main

import "fmt"

func main() {
	var j = 20
	for i := 0; i < 9; i++ {
		var j = 100
		fmt.Println(j + i)
	}
	fmt.Printf("external: j = %d\n", j)
}

// 返回 x 除以 y 的商和余数
func divide(m, n int) (quotient, remainder int) {
	quotient = m / n
	remainder = m % n
	return
}


