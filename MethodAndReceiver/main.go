package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string, age int) person {
	return person {
		name : name,
		age : age,
	}
}

func (p *person) talk() {
	fmt.Printf("%s:你好!\n", p.name)
}

func (p *person) grow() {
	p.age++
}

func main() {
	p1 := newPerson("郎力", 18)
	p1.talk()
	p1.grow()
	fmt.Printf("%d\n", p1.age)

}