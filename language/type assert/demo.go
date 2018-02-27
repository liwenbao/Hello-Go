//Demo
//use variable.(type) to switch variable's type, variable must be interface
//use variable.(typename) to assert variable's type, variable must be interface
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Talker interface {
	Talk()
}

type Eat interface {
	Eat()
}

type Animal struct {
}

func (a *Animal) Talk() {
	fmt.Printf("$@^&`^#@*\n")
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Talk() {
	fmt.Printf("I'm %s, i'm %d years old.\n", p.Name, p.Age)
}

func (p *Person) PTalk() {
	fmt.Printf("I'm a person, my name is %s, i'm %d years old.\n", p.Name, p.Age)
}

func main() {
	var i int = 100
	//log.Println(i.(type)) err: use of .(type) outside type switch
	TypeSwitch(i)
	TypeAssert(i)

	p := Person{"zhangsan", 20}
	TypeSwitch(p)
	TypeSwitch(&p)

	TypeAssert(p)
	TypeAssert(&p)

	a := Animal{}
	TypeSwitch(a)
	TypeSwitch(&a)

	TypeAssert(a)
	TypeAssert(&a)

	//t := i.(Person) err: invalid type assertion: i.(Person) (non-interface type int on left)
	//t := (&i).(Person) err: invalid type assertion: (&i).(Person) (non-interface type *int on left)
	//t := p.(Person) err: invalid type assertion: p.(Person) (non-interface type Person on left)

}

func TypeSwitch(a interface{}) {
	fmt.Print("TypeSwitch ")
	switch a.(type) {
	case int:
		s := a.(int)
		fmt.Println(strconv.Itoa(s))
	case Person:
		p := a.(Person)
		p.PTalk()
	case Talker:
		t := a.(Talker)
		t.Talk()
	default:
		fmt.Println("unknow", reflect.ValueOf(a).Type())
	}
}

func TypeAssert(a interface{}) {
	fmt.Print("TypeAssert ")
	if v, ok := a.(int); ok {
		fmt.Println(strconv.Itoa(v))
		return
	}
	if v, ok := a.(Person); ok {
		v.PTalk()
		return
	}
	if v, ok := a.(Talker); ok {
		v.Talk()
		return
	}
	fmt.Println("unknow", reflect.ValueOf(a).Type())
}
