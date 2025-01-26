package main

import (
	"modtest/mylib"
	"strconv"
	"fmt"
	"os"
)

func getPassedArgs(minArgs int) []int {
	if len(os.Args) < minArgs {
		fmt.Printf("Input should longer than %v", minArgs)
		os.Exit(1)
	}

	var args []int
	for i:=1; i<len(os.Args); i++ {
		val, _ := strconv.Atoi(os.Args[i])
		args=append(args, val)
	}

	return args
}

func appendLocal(args []int) []int {
	maxVal := 0

	for i:=0; i<len(args); i++ {
		maxVal = max(maxVal, args[i])
	}

	local := []int {maxVal+1, maxVal+2, maxVal+3, maxVal+4}
	args=append(args, local...)

	return args
}

func main() {
	fmt.Println("---------------------------------")
	fmt.Println("          Basic Feature          ")
	fmt.Println("---------------------------------")
	var args []int
	args = getPassedArgs(4)

	local:=appendLocal(args)
	fmt.Println("Local Args: ", local)
	fmt.Println("Os Args: ", args)

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println("   compare cpoy() and append()   ")
	fmt.Println("---------------------------------")

	// slice as an argument of func - pass by reference
	mylib.Linked(local)
	fmt.Println("Local Args: ", local, fmt.Sprintf("addr:%p, len:%d, cap:%d", local, len(local), cap(local)), "\n")

	// if slice doesn't has enough capacity, it will point to the new addr and expand capacity
	mylib.Append(local)
	fmt.Println("Local Args: ", local, fmt.Sprintf("addr:%p, len:%d, cap:%d", local, len(local), cap(local)), "\n")

	mylib.Copy(local)
	fmt.Println("Local Args: ", local, fmt.Sprintf("addr:%p, len:%d, cap:%d", local, len(local), cap(local)))
}