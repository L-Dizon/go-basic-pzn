package main

import "fmt"

func main() {
	// basic arithmetic, will explore more later
	var a = 10
	var b = 20
	var d = 5
	var c = a + b/d
	fmt.Println(c)
	// + - * / and % for modulus

	// modulus
	var e = 10
	var f = 3
	var g = e % f
	fmt.Println(g)

	// augmented assignment or operation with itself (variable +=, -=, *=, /=, %=).
	var h = 10
	h += 10 // h = h + 10
	fmt.Println(h)
	h -= 15 // h = h - 15
	fmt.Println(h)
	h *= 5 // h = h * 5
	fmt.Println(h)
	h /= 5 // h = h / 5
	fmt.Println(h)
	h %= 5 // h = h % 5
	fmt.Println(h)

	// unary operators, or operations that act on a single operand
	var i = 10
	var j = -i
	i++ // i = i + 1
	fmt.Println(i)
	i-- // i = i - 1
	fmt.Println(i)
	j++
	fmt.Println(j)

	// ++i and --i are not supported
}
