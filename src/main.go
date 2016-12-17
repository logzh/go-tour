package main

import (
	"fmt"
	"./basics"
)

type person struct {
	name string
	age  int
}

func older(p1, p2 person)(person, int)  {
	if p1.age > p2.age {
		return p1, p1.age -p2.age
	}

	return p2, p2.age - p1.age
}

func add(x int, y int)(z int)  {
	return x + y
}

func max(x, y int)int  {
	if x > y{
		return x
	}
	return y
}

func init()  {
	fmt.Println("main package init")
}

func main() {
	var P1, P2 person
	P1.name = "spence"
	P1.age = 29
	fmt.Println(P1.name)
	P2.name = "jin"
	P2.age = 28
	fmt.Println(P2.name)

	fmt.Println(older(P1, P2))

	var i int = 5
	fmt.Println(i)

	var a, b, c int = 1, 2, 3
	fmt.Println(a + b + c)

	const PI = 3.141592653
	fmt.Println(PI)

	var enable bool = false
	fmt.Println("enable:", enable)

	var s string = "test me"
	s = "123e"
	fmt.Println(s)

	var ll string = `
	new line
	is hh
	`
	fmt.Println(ll + " append")

	fmt.Println(len(ll))
	fmt.Println(ll[7:])

	arr := [10]int{1, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}
	fmt.Println(arr)

	var numbers map[string]int
	numbers = make(map[string]int)
	fmt.Println(numbers)
	numbers["aa"] = 12
	numbers["bb"] = 23

	fmt.Println(numbers)
	for k, v := range numbers {
		fmt.Println("numbers's key:", k)
		fmt.Println("numbers's val:", v)
	}

	if i > 4 {
		fmt.Println("i>4的")
	} else if i == 4 {
		fmt.Println("i=4的")
	} else {
		fmt.Println("i<4的")
	}

	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	switch i {
	case 1:
		fmt.Println("case 1")
		break
	case 2:
		fmt.Println("case 2")
		break
	default:
		fmt.Println("case default")
		break

	}

	fmt.Println(add(i, P1.age))

	basics.Imports(3.14)

}
