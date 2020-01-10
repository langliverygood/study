package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("132.txt")
	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}
	fmt.Printf("%T\n", file)
	defer file.Close()

}