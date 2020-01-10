package main

import "fmt"

//接口是一种类型,它规定了变量有哪些方法
type speaker interface {
	speak() //只要实现了speak方法，都是speaker类型
}

type cat struct {

}

type dog struct {

}

type person struct {

}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func (p person) speak() {
	fmt.Println("啊啊啊~")
}

func da(x speaker) {
	// 接收一个参数，传进来什么，打印什么
	x.speak() 
}


// 在编程中会遇到以下场景：
// 1.不关心变量的类型，只关心能调用的它的方法。
// 
func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)
}