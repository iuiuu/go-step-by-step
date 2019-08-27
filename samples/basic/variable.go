/**
代码注释
*/
package main

import "fmt"
import "math/rand"

func main() {
	var i int = 20
	var j int
	var k = 85
	fmt.Println(i) // print 20
	fmt.Println(j) // print 0
	fmt.Println(k) // print 85

	j = 10
	fmt.Println(j) // print 10

	// 声明和初始化多个变量
	var m, n int = 100, 200
	fmt.Printf("m = %d", m) // 输出 m = 100
	fmt.Println()           // 换行
	fmt.Printf("n = %d", n) // 输出 n = 200
	fmt.Println()           // 换行

	// 声明多个变量（零值初始化）
	var p, q int
	fmt.Printf("p = %d", p) // 输出 p = 0
	fmt.Println()           // 换行
	fmt.Printf("q = %d", q) // 输出 q = 0
	fmt.Println()           // 换行

	// var a int, var b bool
	// var x int
	// var y bool
	// var x int, y bool	// syntax error: unexpected comma at end of statement
	// fmt.Println(x) // print 0
	// fmt.Println(y) // print false
	var x, y, z = true, "abc", 20
	fmt.Printf("x = %t", x) // 输出 x = true
	fmt.Println()           // 换行
	fmt.Printf("y = %s", y) // 输出 y = "abc"
	fmt.Println()           // 换行
	fmt.Printf("z = %d", z) // 输出 z = 20
	fmt.Println()           // 换行

	t := 999
	r := rand.Float64() * 3.0

	fmt.Printf("t = %d", t) // 输出 t = 999
	fmt.Println()           // 换行
	fmt.Printf("r = %f", r) // 输出 r = 1.813981
	fmt.Println()

	c, d := divide(m, n)
	fmt.Printf("c = %d, d = %d", c, d) // 输出 c = 0, d = 100
	fmt.Println()

	global = "in main()"
	fmt.Println(global) // 输出 in main()
	global := "new variable"
	fmt.Println(global) // 输出 new variable
	// f := "f"
	// // global := "local"
	// fmt.Println(f)      // "f"; local var f shadows package-level func f
	// fmt.Println(global) // "g"; package-level var
	// fmt.Println(h) // compile error: undefined: h
}

func f() {}

var global = "global"

// 返回 x 除以 y 的商和余数
func divide(m, n int) (quotient, remainder int) {
	quotient = m / n
	remainder = m % n
	return
}
