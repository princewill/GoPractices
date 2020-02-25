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

	

}
