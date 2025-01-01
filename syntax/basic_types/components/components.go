package components

import "fmt"

// Out 结构体组合结构体
type Out struct {
	Inner
}

type NameI interface {
	name() string
}

func (i Inner) name() string {
	return "innner"
}

// Out2 结构体组合指针
type Out2 struct {
	Inner *Inner
}

type Inner struct {
}

func (i Inner) hello() {
	fmt.Println("hello this is ", i.name())
}

func Components() {
	var o Out
	o.hello()
	//var o2 Out2
	o2 := Out2{
		Inner: &Inner{},
	}
	o2.Inner.hello()

}
