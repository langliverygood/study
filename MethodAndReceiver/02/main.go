package main

import "fmt"

type myInt int

func (i *myInt) hello() {
	fmt.Printf("我是int:%d\n", i)
}

func main() {
	m := myInt(15)
	m.hello()
}