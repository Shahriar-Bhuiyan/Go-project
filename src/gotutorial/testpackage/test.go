package testpackage

import "fmt"

func MyFunction(step string) {
	if step == "1" {
		fmt.Println("Step 1")
	} else if step == "2" {
		fmt.Println("Step 2")
	} else if step == "3" {
		fmt.Println("Step 3")
	} else if step == "4" {
		fmt.Println("Step 4")
	} else if step == "This is a string  " { 
		fmt.Println(step)
	} else {
		fmt.Println("Step not supported")
	}
}



