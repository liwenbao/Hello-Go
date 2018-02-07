//Demo
//struct inherit
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

type Teacher struct {
	Person
	Course string
}

func (t Teacher) Talk() {
	log.Printf("I'm a teacher. I teach %s.\n", t.Course)
}

type Student struct {
	*Person
	Grade int
}

func (s Student) Talk() {
	log.Printf("I'm a student. I study at %d grade.\n", s.Grade)
}

func main() {
	t1 := Teacher{Person{"张三", 30}, "语文"}
	t2 := &Teacher{Person{"李四", 30}, "数学"}

	t1.Person.Talk()
	t1.Talk()
	t1.Grow()
	t1.Person.Talk() //age changed
	log.Printf("meathed set of %s is\n", t1.Name)
	PrintMeathedSet(reflect.TypeOf(t1)) //not contains Grow

	log.Println()

	t2.Person.Talk()
	t2.Talk()
	t2.Grow()
	t2.Person.Talk() //age changed
	log.Printf("meathed set of %s is\n", t2.Name)
	PrintMeathedSet(reflect.TypeOf(t2))

	log.Println()

	s1 := Student{&Person{"王五", 10}, 4}
	s2 := &Student{&Person{"马六", 9}, 3}

	s1.Person.Talk()
	s1.Talk()
	s1.Grow()
	s1.Person.Talk() //age changed
	log.Printf("meathed set of %s is\n", s1.Name)
	PrintMeathedSet(reflect.TypeOf(s1))

	log.Println()

	s2.Person.Talk()
	s2.Talk()
	s2.Grow()
	s2.Person.Talk() //age changed
	log.Printf("meathed set of %s is\n", s2.Name)
	PrintMeathedSet(reflect.TypeOf(s2))
}

func PrintMeathedSet(v reflect.Type) {
	for i := 0; i < v.NumMethod(); i++ {
		m := v.Method(i)
		log.Printf("\tmeathed %d: %s\n", i+1, m.Name)
	}
}
