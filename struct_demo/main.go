package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson_value(name string, age int) person {
	return person {
		name: name,
		age : age,
	}
}

func newPerson_point(name string, age int) *person {
	return &person {
		name: name,
		age : age,
	}
}

func main() {
	p1 := newPerson_value("元帅", 18)
	p2 := newPerson_value("周林", 9200)
	fmt.Printf("%p: %v\n", &p1, p1)
	fmt.Printf("%p: %v\n", &p2, p2)

	p3 := newPerson_point("元帅", 18)
	p4 := newPerson_point("周林", 9200)

	fmt.Printf("%p: %#v\n", p3, *p3)
	fmt.Printf("%p: %#v\n", &p3, *p3)
	fmt.Printf("%p: %#v\n", p4, *p4)
	fmt.Printf("%p: %#v\n", &p4, *p3)
}