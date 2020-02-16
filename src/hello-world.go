package main

import "fmt"

func main() {
	fmt.Println("hello world"); fmt.Println("words" + "good"); fmt.Println(23); fmt.Println(true)

	var a = "variable"
	fmt.Println(a)

	var ab, b int = 1, 2
	fmt.Println(ab,b)

	z := "friend"
	fmt.Println(z)

	i := 0

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for i:=0; i<9; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("loop")
		break
	}

}
