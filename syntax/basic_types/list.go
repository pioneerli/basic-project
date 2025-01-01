package main

type List interface {
	Add(index int, val any)
	Append(val any)
	Delete(index int)
}
type LinkedList struct {
	head string
}

func (l LinkedList) Add(index int, val any) {
	//TODO implement me
	panic("implement me")
}

func (l LinkedList) Append(val any) {
	//TODO implement me
	panic("implement me")
}

func (l LinkedList) Delete(index int) {
	//TODO implement me
	panic("implement me")
}
