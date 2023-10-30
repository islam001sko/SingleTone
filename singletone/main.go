package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
