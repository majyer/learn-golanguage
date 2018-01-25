package main

import "fmt"

func main() {
	func_b_0()
	func_b_1()
	func_b_2()
}

func func_b_0() {
	fmt.Println("func_b_0...top")
	a := 5
	defer fmt.Println("a=", a)
	a++
}
func func_b_1() {
	fmt.Println("func_b_1...top")
	a := 5
	defer func() {
		fmt.Println("a=", a)
	}()
	a++
}
func func_b_2() {
	fmt.Println("func_b_2...top")
	a := 5
	defer func(a int) {
		fmt.Println("a=", a)
	}(a)
	a++
}