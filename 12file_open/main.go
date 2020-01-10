package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("Open file failed, er:%v", err)
		return
	}

	defer fileObj.Close()

	//var tmp = make([]byte, 128)
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("Read finish")
			return
		}
		if err != nil {
			fmt.Printf("Read file failed, err:%v\n", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:]))
		if n < 128 {
			return
		}
	}
}

func readFromFileByBufio() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("Read file failed, err:%v\n", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Read finish")
			return
		}
		if err != nil {
			fmt.Printf("Read file failed, err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("Read file failed, err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	//readFromFile1()
	readFromFileByBufio()
	//readFromFileByIoutil()

}
