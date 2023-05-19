package main

import "fmt"

func main() {
	var Star int
	for i := 0; i < 11; i++ {
		Star = 1
		for j := 5; j >= 0; j++ {
			fmt.Print(" ")
		}
		count := Star
		for count == 0 {
			fmt.Print("*")
		}
		fmt.Println()
		Star += 2
	}
}
