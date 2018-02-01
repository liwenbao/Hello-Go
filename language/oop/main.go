package main

import (
	"log"
	"reflect"
)

type Talker interface {
	Talk() string
}

type Person struct {
}

func (p Person) Talk() string {
	return "i'm a person"
}

type Student struct {
	Person
}

type Animal struct {
}

func (a *Animal) Talk() string {
	return "#!$%@^~"
}

type Dog struct {
	Animal
	Name string
	Age  int
}

func (d Dog) ChangeName(newName string) {
	d.Name = newName
}

func (d *Dog) AddAge() {
	d.Age++
}

func (d *Dog) Talk() string {
	return "Wang Wang Wang"
}

type TalkerBase struct {
}

func (t *TalkerBase) Talk() string {
	return Talk(t)
}

func Talk(t Talker) string {
	log.Println("t is", reflect.ValueOf(t).Type().String())
	return "i'm a talker"
}

type Talker1 struct {
	TalkerBase
}

type Talker2 struct {
	TalkerBase
}

func (t *Talker2) Talk() string {
	return Talk(t)
}

func main() {
	var person Person = Person{}
	log.Println(reflect.TypeOf(person).String(), "talk", person.Talk())

	//用Person类型直接初始化Talker接口类型的变量
	var personTalk Talker = Person{}
	log.Println(reflect.TypeOf(personTalk).String(), "talk", personTalk.Talk())

	//将一个Person类型的对象赋值给Talker接口类型的变量
	personTalk = person

	//对值为Person类型的interface变量，进行类型断言
	var personInterface interface{} = person
	p, ok := personInterface.(Talker)
	if ok {
		log.Println(reflect.TypeOf(p).String(), "talk", p.Talk())
	}

	//Person类型上定义了Talk方法，*Person类型也自动具有该方法
	personPoint := &person
	log.Println(reflect.TypeOf(personPoint).String(), "talk", personPoint.Talk())

	var student Student = Student{}
	log.Println(reflect.TypeOf(student).String(), "talk", student.Talk())

	var t Talker = student
	p, ok = t.(Talker)
	if ok {
		log.Println("student is Talker")
	}

	var animal Animal = Animal{}
	log.Println(reflect.TypeOf(animal).String(), "talk", animal.Talk())
	//t = animal err: cannot use animal (type Animal) as type Talker in assignment:Animal does not implement Talker (Talk method has pointer receiver)
	t = &animal
	p, ok = t.(Talker)
	if ok {
		log.Println("student is Talker")
	}

	/*
		NOTE:
		ChangeName是Type Reciever方法，不论申明的变量是Dog还是*Dog，实际调用都会以Dog.ChangeName方式进行。在方法内，Dog被复制，无法修改原始变量的字段。
		AddAge是Pointer Reciever方法，不论申明的变量是Dog还是*Dog，实际调用都会以*Dog.ChangeName方式进行。在方法内，Dog的指针被复制，可以修改原始变量的字段。
		虽然Dog类型变量可以直接调用AddAge方法，但是实际上调用被转换为先对Dog变量取地址，而后再调用指针的AddAge方法。Dog类型的方法集上并没有AddAge方法。
		但是，*Dog类型的方法集上有ChangeName方法。
		Dog继承Animal后，继承了Pointer Reciever的Talk方法。Talk在*Dog的方法集中，不在Dog的方法集中。因此Dog不是Talker接口的实现。
	*/
	var dog1 Dog = Dog{Name: "Dog1"}
	log.Println(dog1) //{Dog1 0}
	dog1.ChangeName("Dog001")
	dog1.AddAge()
	log.Println(dog1) //{Dog1 1}
	v := reflect.TypeOf(dog1)
	log.Printf("type %s has %d meatheds\n", v.String(), v.NumMethod()) //print: type main.Dog has 1 meatheds
	for i := 0; i < v.NumMethod(); i++ {
		log.Printf("meathed %d: %s\n", i, v.Method(i).Name)
	}
	//t = dog1 err:cannot use dog1 (type Dog) as type Talker in assignment:Dog does not implement Talker (Talk method has pointer receiver)

	var dog2 *Dog = &Dog{Name: "Dog2"}
	log.Println(dog2) //&{Dog2 0}
	dog2.ChangeName("Dog002")
	dog2.AddAge()
	log.Println(dog2) //&{Dog2 1}
	v = reflect.TypeOf(dog2)
	log.Printf("type %s has %d meatheds\n", v.String(), v.NumMethod()) //print: type *main.Dog has 3 meatheds
	for i := 0; i < v.NumMethod(); i++ {
		log.Printf("meathed %d: %s\n", i, v.Method(i).Name)
	}
	t = dog2
	p, ok = t.(Talker)
	if ok {
		log.Println("*Dog is Talker", p.Talk())
	}

	/*
		Talk方法接收一个Talker接口的实例
		t1调用Talk方法时，因为Talker1并没有显示实现Talk，所以实际是在调用TalkerBase.Talk。在TalkerBase.Talk方法内，t的类型是*TalkerBase
		t2调用Talk方法时，实际是在调用Talker2.Talk。在Talker2.Talk方法内，t的类型是*Talker2
		因此Talk方法中做出了不同的类型判断。
		通t1.Talk调用，在Talk方法中得到的Talker实际是一个*TalkerBase
	*/
	var t1 = Talker1{}
	var t2 = Talker2{}
	log.Println("t1.Talk", t1.Talk()) //t is *main.TalkerBase
	log.Println("t2.Talk", t2.Talk()) //t is *main.Talker2
}
