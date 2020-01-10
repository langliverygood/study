
package main

import "fmt"

/*type myInt int
type yourInt = int

type person struct {
	name string
	age int
	gender string
	hobby []string
}

func main() {
	var langli person
	langli.name = "郎力"
	langli.age = 23
	langli.gender = "男"
	langli.hobby = []string{ "篮球", "羽毛球", "象棋"}
	fmt.Printf("%T\n", langli)
	fmt.Println(langli.hobby)

}*/

//指针结构体
type person struct {
	name string
	age int
}

func f(x *person) {
	x.age = 20
}

func main() {

	//匿名结构体
	/*var s struct {
		name string
		age int
	}
	fmt.Printf("%T\n", s)*/

	/* 
	var p person
	p.name = "郎力"
	p.age = 18
	f(&p)
	fmt.Println(p.age)*/

	/* 声明结构体指针
	var p2 = new(person)
	//p2.name = "张三"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2) */
	
	/* 声明结构体并且初始化 */
	var p3 = person {
		name: "郎力",
		age: 18,
	}
	fmt.Printf("%#v\n", p3)

	p4 := person {
		"张三",
		16,
	}
	fmt.Printf("%#v\n", p4)
}

// 结构体占用一块连续的内存空间

