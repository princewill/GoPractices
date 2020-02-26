package main

import "fmt"

func main() {

	a := make(map[string]int)

	a["k1"] = 1
	a["k2"] = 2

	fmt.Println(a)

	v1 := a["k1"]
	fmt.Println(v1)

	fmt.Println(len(a))

	delete(a, "k1")
	fmt.Println(a)

	_, prs := a["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
