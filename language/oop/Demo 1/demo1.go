//Demo
//type receiver function use copy of struct
//pointer receiver function use copy of point
package main

import (
	"log"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Talk() {
	log.Printf("I'm %s. I'm %d years old.\n", p.Name, p.Age)
}

func (p Person) AddAge() {
	p.Age += 1
	log.Printf("%s add age by type receiver function\n", p.Name)
}

func (p *Person) Grow() {
	p.Age += 1
	log.Printf("%s add age by pointer receiver function\n", p.Name)
}

func main() {
	p1 := Person{"张三", 10}  //p1 is a struct
	p2 := &Person{"李四", 10} //p2 is a point of struct

	p1.Talk()
	p2.Talk() //either struct or pointer can call type receiver function

	p1.AddAge()
	p1.Talk() //in type receiver function, struct is a copy. so can't change field in type receiver function.

	p2.AddAge()
	p2.Talk() //call type receiver function by pointer, field is not changed too.

	p2.Grow()
	p2.Talk() //in pointer receiver function, address is a copy. tow different address point to same struct. so can change field in pointer receiver function.

	p1.Grow()
	p1.Talk() //call pointer receiver function by struct, field is changed too.
}
