package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	x := 9
	var y int = 7
	var sum int = x + y
	fmt.Println(sum)

	if y > 6 {
		fmt.Println("More than 6")
	}

	var a [5]int
	a[2] = 7

	b := [5]int{2, 5, 3, 9, 6}

	c := []int{5, 6}

	c = append(c, 18)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Objects (maps)
	vertices := make(map[string]int)
	vertices["square"] = 2
	vertices["triangle"] = 3

	fmt.Println(vertices)
	//delete(vertices, "square")
	fmt.Println(vertices)

	//For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//For loop over an array
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	//For loop over a map
	for key, value := range vertices {
		fmt.Println("key:", key, "value:", value)
	}

	// while loop
	z := 0
	for z < 5 {
		fmt.Println(z)
		z++
	}

	//Build A DA
	type person struct {
		name string
		age  int
	}

	people := [5]person{
		{
			name: "Jake",
			age:  23,
		},
	}

	fmt.Println((people[0].name))

}

func sum(x int, y int) int {
	return x + y
}
