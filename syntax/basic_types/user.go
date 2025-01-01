package main

import (
	"basic-project/syntax/basic_types/components"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func (u User) ChangeName(name string) {
	fmt.Printf("u address %p \n", &u)
	u.Name = name
}

// ChangeAge *User 表示指针类型
func (u *User) ChangeAge(age int) {
	u.Age = age
}

func NewUser() {
	//// u1 是指向一个 User 对象的指针
	//u1 := &User{}
	//println(u1)
	//
	//// u2 中的字段都是零值
	//u2 := User{}
	//println(u2)
	//// 修改 u2 的字段
	//u2.Name = "Jerry"
	//
	//// u3 中的字段也都是零值
	//var u3 User
	//println(u3)
	//
	//// 初始化的同时，还赋值了 Name
	//var u4 User = User{Name: "Tom"}
	//println(u4)
	//
	//// 没有指定字段名，按照字段顺序赋值
	//// 必须全部赋值
	//var u5 User = User{"Tom", 18}
	//println(u5)

	u1 := &User{Name: "kevin", Age: 22}
	println(u1)
	fmt.Printf("%+v", *u1)
	u1.ChangeAge(56)
	fmt.Printf("%+v", *u1)
}

func main() {
	//NewUser()
	//结构体包含结构体
	//o1 := &Out{}

	//o1.inner.hello()
	//Components()
	components.Components()
}
