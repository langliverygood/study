package main

import (
	"fmt"
	"../calc"
)

//空接口没有必要起名字，通常定义成下面的格式
//interface{}
//所有的类型都实现了空接口，也就是一体类型的变量都能保存到空接口中


// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("Type:%T Value:%v\n", a, a)
}

//类型断言
//空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println(str)
	}
	
}

func assign2(a interface{}) {
	switch t := a.(type) {
	case string:
		fmt.Println("是字符串:", a.(string))
	case int:
		fmt.Println("是整型:", a.(int))
	case bool:
		fmt.Println("是布尔:", a.(bool))
		fmt.Println(t)
	}

	fmt.Println(calc.Add(3, 2))
	
}
func main() {

	// 当保存值不确定时，可以使用空接口
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "郎力"
	m1["age"] = 23
	m1["hobby"] = [...]string{"象棋", "骑车", }
	fmt.Println(m1)

	show(false)
	show(nil)
	show(24)

	assign(100)
	assign2("aaa")
	assign2(true)

}