package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Open file failed, err:%v", err)
		return
	}
	file.Write([]byte("你好\n"))
	file.WriteString("郎力\n")
	file.Close()

}
