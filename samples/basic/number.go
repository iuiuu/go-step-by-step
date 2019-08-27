package main

import "math"
import "fmt"

func main() {
	fmt.Println(math.MinInt8, "~", math.MaxInt8)
	fmt.Println(math.MinInt16, "~", math.MaxInt16)
	fmt.Println(math.MinInt32, "~", math.MaxInt32)
	fmt.Println(math.MinInt64, "~", math.MaxInt64)
	fmt.Println(0, "~", math.MaxUint8)
	fmt.Println(0, "~", math.MaxUint16)
	fmt.Println(0, "~", math.MaxUint32)

	//===============================

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	fmt.Println(math.NaN)
	fmt.Println(math.Inf(-3))

	var a float64
	// var b float64

	fmt.Printf("1/%g is: %g\n", a, 1/a)
	fmt.Printf("-1/%g is: %g\n", a, -1/a)
	fmt.Printf("0/%g is: %g\n", a, 0/a)

	fmt.Printf("0/%g is: %g\n", a, a == math.Inf(1))
	fmt.Println(a == math.Inf(1))
	fmt.Println(1/a == math.Inf(1))
	fmt.Println(1/a == math.NaN())
	fmt.Println(0/a == math.NaN())
	fmt.Println(math.IsNaN(0 / a))

	// var z float64
	// fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	// value := uint64(math.MaxUint64)
	// fmt.Println("Hello, world") // print "hello, world"
}
