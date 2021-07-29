package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai : ")
	fmt.Scan(&nilai)

	if nilai%3 == 0 {
		fmt.Println("Hello")
	} else if nilai%5 == 0 {
		fmt.Println("World")
	} else if nilai%5 == 0 && nilai%3 == 0 {
		fmt.Println("Hello World")
	}
}
