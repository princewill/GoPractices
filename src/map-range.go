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

	n := map[string]int {"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	nums := []int{1,2,3}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	kvs := map[string]string{"a":"apple","b":"banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
