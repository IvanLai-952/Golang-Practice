package main

import "fmt"

func main() {
	var name string
	var age int

	name, age = "Oscar", 30

	fmt.Println(name, "is", age, "years old")
	fmt.Printf("%s is %d years old\n", name, age)

	fmt.Printf("-------------------------------------------\n")

	t1 := "text"
	t2 := []string{"apple", "banana"}
	t3 := map[string]float64{"straberry": 32.2, "blueberry": 30.9}
	t4 := 2
	t5 := 3.141596
	t6 := true

	fmt.Printf("Data Type for each variable:\n")
	fmt.Printf("t1: %T\n", t1)
	fmt.Printf("t2: %T\n", t2)
	fmt.Printf("t2: %T\n", t3)
	fmt.Printf("t2: %T\n", t4)
	fmt.Printf("t2: %T\n", t5)
	fmt.Printf("t2: %T\n", t6)

	fmt.Printf("-------------------------------------------\n")

	var a, b int = 10, 20
	var ua uint = 113
	fmt.Println(a, b)
	fmt.Printf("Hex of variable a: 0x%08x \n", ua)

	var c float32 = 3.1415926
	fmt.Println(c)

	var d string = "Hello World"
	fmt.Println(d)

	fmt.Printf("-------------------------------------------\n")

	// anonymous placeholder
	var name2 string
	name2, _ = "Iron man", "Dr.strange"
	fmt.Println(name2)
}
