package main

import "fmt"

var level = "global"
var val = 1

func funcA() {
	fmt.Println("funcA start, level: ", level)
	fmt.Println("funcA start, val: ", val)
}

func main() {
	fmt.Println("main func start, level: ", level)
	fmt.Println("main func start, val: ", val)

	level := "local"
	val = 0

	if true {
		funcA()
		fmt.Println("main func start, level: ", level)
		fmt.Println("main func start, val: ", val)
	}

	fmt.Println("main func start, level: ", level)
	fmt.Println("main func start, val: ", val)
}
