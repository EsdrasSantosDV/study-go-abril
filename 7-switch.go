package main

import "fmt"

func main() {
	switch 100 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 100:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}

}
