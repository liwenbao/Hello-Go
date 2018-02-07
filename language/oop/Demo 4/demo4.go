//Demo
//interface & polymorphism
package main

import (
	"log"
)

type Talker interface {
	Talk()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Talk() {
	log.Printf("I'm %s. I'm %d years old.\n", p.Name, p.Age)
}

type Animal struct {
}

func (a *Animal) Talk() {
	log.Printf("@#$~&^*…….\n")
}

func main() {
	p1 := Person{"张三", 30}
	Talk(p1)
	p2 := &Person{"李四", 25}
	Talk(p2)
	a1 := Animal{}
	//Talk(11) //err: Animal does not implement Talker (Talk method has pointer receiver)
	a1.Talk()
	a2 := &Animal{}
	Talk(a2)
}

func Talk(t Talker) {
	t.Talk()
}
