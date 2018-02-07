//Demo
//meathed set of struct contains type receiver function only.
//meathed set of pointer contains both type receiver function and prointer receiver funtion.
package main

import (
	"log"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Talk() {
	log.Printf("I'm %s. I'm %d years old.\n", p.Name, p.Age)
}

func (p *Person) Grow() {
	p.Age += 1
	log.Printf("%s add age.\n", p.Name)
}

func main() {
	p1 := Person{"张三", 10}  //p1 is a struct
	p2 := &Person{"李四", 10} //p2 is a point of struct

	log.Printf("meathed set of %s is\n", p1.Name)
	PrintMeathedSet(reflect.TypeOf(p1)) //meathed set of struct contains type receiver function only.

	log.Printf("meathed set of %s is\n", p2.Name)
	PrintMeathedSet(reflect.TypeOf(p2)) //meathed set of pointer contains both type receiver function and prointer receiver funtion.
}

func PrintMeathedSet(v reflect.Type) {
	for i := 0; i < v.NumMethod(); i++ {
		m := v.Method(i)
		log.Printf("\tmeathed %d: %s\n", i+1, m.Name)
	}
}
