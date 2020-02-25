package main

import "fmt"

func main() {

	var a[5] int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)

	b := [4]int{1, 2, 3, 4}
	fmt.Println(b)

	var c [3][3]int
	fmt.Println(c)

	for i:=0; i < 3; i++ {
		for j:=0; j < 3; j++ {
			c[i][j] = i + j
		}
	}

	foo := make([]string, 3)
	fmt.Println(foo)

	foo[0] = "a"
	foo[1] = "b"
	foo[2] = "c"

	foo = append(foo, "d")
	fmt.Println(foo)

	bar := make([]string, len(foo))
	copy(foo, bar)
	fmt.Println(bar)

	l := bar[2:5]
	fmt.Println("sl1:", l)

	l = bar[:5]
	fmt.Println("sl2:", l)

	l = bar[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

}
