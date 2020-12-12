package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

func main() {
	mylist := list.New()
	mylist.PushFront("frist")
	mylist.PushBack(101)
	mylist.PushBack(102)
	mylist.PushBack(103)
	mylist.PushBack("end")
	//重点!!!!()一定要加括号!!!大概表示结构体
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	mynewlist := list.New()
	for i := 0; i < 10; i++ {
		mynewlist.PushBack(i + 1)
	}

	fmt.Println("-------我是分割线-------")

	for i := mynewlist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("-------我是分割线-------")

	var i int
	var i2 int16
	var i1 int8
	var i3 int32
	var i4 int64
	fmt.Println(unsafe.Sizeof(i), unsafe.Sizeof(i1), unsafe.Sizeof(i2), unsafe.Sizeof(i3), unsafe.Sizeof(i4))
}
