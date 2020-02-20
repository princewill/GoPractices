package main

import (
	"fmt"
	"time"
)

func main() {

	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if a:=7; a % 2 == 0 {
		fmt.Println("7 is even")
	}

	b:="Tuesday"; switch b {
	case "Monday":
		fmt.Println("mon")
	case "Tuesday":
		fmt.Println("tues")
	default:
		fmt.Println("week-day")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("week-end")
	default:
		fmt.Println("week-day")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
