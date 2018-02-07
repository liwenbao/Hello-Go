//Demo
package main

import (
	"log"
	"math/rand"
)

func main() {
	Create()
	Append()
	Foreach()
	Sub()
	Remove()
	Insert()
	Copy()
	SliceAndArray()
	SliceAndAddress()
}

func Create() {
	s1 := []int{}
	log.Printf("slice has %d elements, they are:%d", len(s1), s1)

	s2 := make([]int, 0)
	log.Printf("slice has %d elements, they are:%d", len(s2), s2)

	s3 := make([]int, 10)
	log.Printf("slice has %d elements, they are:%d", len(s3), s3)

	s4 := new([]int)
	//log.Printf("slice has %d elements, they are:%d", len(s4), s4) //err invalid argument s4 (type *[]int) for len
	log.Printf("slice has %d elements, they are:%d", len(*s4), s4) //type of s4 is &[]int

	s5 := []int{1, 2, 3}
	log.Printf("slice has %d elements, they are:%d", len(s5), s5)

	s6 := [...]int{}
	log.Printf("slice has %d elements, they are:%d", len(s6), s6)

	s7 := make([]int, 0, 10)
	log.Printf("slice has %d elements, they are:%d", len(s7), s7)
}

func Append() {
	s1 := []int{1, 2, 3}
	s1 = append(s1, 4)

	s1 = append(s1, 5, 6)

	s2 := []int{7, 8}
	s1 = append(s1, s2...)

	log.Printf("slice has %d elements, they are:%d", len(s1), s1)
}

func Foreach() {
	s1 := []int{}
	for i := 0; i < 10; i++ {
		s1 = append(s1, rand.Intn(100))
	}
	log.Printf("slice has %d elements, they are:%d", len(s1), s1)

	for i, v := range s1 {
		if v > 50 {
			log.Printf("%d, %d", i, v)
		}
	}
}

func Sub() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := s1[3:7]
	log.Printf("slice has %d elements, they are:%d", len(s2), s2)

	s2 = s1[:7]
	log.Printf("slice has %d elements, they are:%d", len(s2), s2)

	s2 = s1[3:]
	log.Printf("slice has %d elements, they are:%d", len(s2), s2)

	s2 = s1[:]
	log.Printf("slice has %d elements, they are:%d", len(s2), s2)

	s2 = s1[:3:7]
	log.Printf("slice has %d elements, they are:%d", len(s2), s2)
}

func Remove() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := 5 //index to remove
	s2 := append(s1[:index], s1[index+1:]...)
	log.Printf("remove index %d: %d", index, s2)
}

func Insert() {
	s1 := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
	index := 5
	inserted := 5
	rear := s1[index:]
	s1 = append(s1[:index], inserted)
	s1 = append(s1, rear...)
	log.Printf("insert into index %d: %d", index, s1)
}

func Copy() {
	s1 := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
	s2 := []int{}
	copy(s2, s1)
	log.Printf("copy to []int{}: %d", s2)

	s2 = []int{10, 11, 12}
	copy(s2, s1)
	log.Printf("copy to []int{10, 11, 12}: %d", s2)

	s2 = make([]int, 5, 10)
	copy(s2, s1)
	log.Printf("copy to make([]int, 5, 10): %d", s2)

	s2 = make([]int, len(s1))
	copy(s2, s1)
	log.Printf("copy to make([]int, len(s1)): %d", s2)

	s2 = make([]int, len(s1)+1)
	copy(s2, s1)
	log.Printf("copy to make([]int, len(s1)+1): %d", s2)
}

func SliceAndArray() {
	//var array [10]int
	array := make([]int, 6)
	log.Printf("array: %d", array)

	s1 := array[:5]
	log.Printf("s1: %d", s1)

	array[1] = 1
	log.Printf("array: %d", array)
	log.Printf("s1: %d", s1)

	s1[2] = 2
	log.Printf("array: %d", array)
	log.Printf("s1: %d", s1)

	//以上示例中，s1始终未超出实际切片的容量，因此操作的是同一个切片

	s1 = append(s1, 10)
	log.Printf("array: %d", array)
	log.Printf("s1: %d", s1)

	//对s1进行append，未超出实际切片容量，s1追加了元素，array的5位置上被修改。

	array[3] = 3
	log.Printf("array: %d", array)
	log.Printf("s1: %d", s1)

	s1[4] = 4
	log.Printf("array: %d", array)
	log.Printf("s1: %d", s1)

	array[5] = 5
	log.Printf("array: %d", array)
	log.Printf("s1: %d", s1)

	//仍未超出实际切片的容量，因此操作的是同一个切片

	s1 = append(s1, 0)
	log.Printf("array: %d", array)
	log.Printf("s1: %d", s1)

	s1[5] = 0
	log.Printf("array: %d", array)
	log.Printf("s1: %d", s1)

	//s1再次append，超出了实际切片容量，此时重新分配了切片对象。s1和array内部不再是同一个切片对象。修改s1的元素不再影响array。
}

func SliceAndAddress() {
	//append时，如果现有切片容量不足，则会重新分配一个切片。新切片与现有切片数据是各自独立的。
	var s1 []int
	log.Printf("s1 len: %d, cap: %d, content: %v", len(s1), cap(s1), s1)

	s2 := append(s1, 0)
	log.Printf("s1 len: %d, cap: %d, content: %v", len(s1), cap(s1), s1)
	log.Printf("s2 len: %d, cap: %d, content: %v", len(s2), cap(s2), s2)

	//因此需要通过切片共享数据时，为了避免切片重新分配导致引用变化，应该使用切片的指针。
	var s3 []int
	var s *[]int = &s3
	log.Printf("s len: %d, cap: %d, content: %v", len(*s), cap(*s), *s)
	*s = append(*s, 0)
	log.Printf("s len: %d, cap: %d, content: %v", len(*s), cap(*s), *s)
}
