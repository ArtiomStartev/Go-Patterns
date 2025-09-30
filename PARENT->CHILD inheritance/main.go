package main

import "fmt"

func main() {
	parentInstance := NewParentInstance()

	fmt.Println(parentInstance.ParentMethodGetProp1())
	fmt.Println(parentInstance.ParentMethodGetProp2())
	fmt.Println(parentInstance.ParentMethodGetGreeting())
	fmt.Println(parentInstance.ChildMethodGetProp1())
	fmt.Println(parentInstance.ChildMethodGetProp2())
	fmt.Println(parentInstance.ChildMethodGetGreeting())
}

// CHILD

type ChildInterface interface {
	ChildMethodGetProp1() string
	ChildMethodGetProp2() string
	ChildMethodGetGreeting() string
}

type ChildStruct struct {
	prop1 string
	prop2 string
}

func (cs *ChildStruct) ChildMethodGetProp1() string {
	return cs.prop1
}

func (cs *ChildStruct) ChildMethodGetProp2() string {
	return cs.prop2
}

func (cs *ChildStruct) ChildMethodGetGreeting() string {
	return "Hello from Child Method!"
}

func NewChildInstance() ChildInterface {
	return &ChildStruct{
		prop1: "Child Property 1",
		prop2: "Child Property 2",
	}
}

// PARENT

type ParentInterface interface {
	ParentMethodGetProp1() string
	ParentMethodGetProp2() string
	ParentMethodGetGreeting() string
	ChildInterface
}

type ParentStruct struct {
	prop1 string
	prop2 string
	ChildInterface
}

func (ps *ParentStruct) ParentMethodGetProp1() string {
	return ps.prop1
}

func (ps *ParentStruct) ParentMethodGetProp2() string {
	return ps.prop2
}

func (ps *ParentStruct) ParentMethodGetGreeting() string {
	return "Hello from Parent Method!"
}

func NewParentInstance() ParentInterface {
	return &ParentStruct{
		prop1:          "Parent Property 1",
		prop2:          "Parent Property 2",
		ChildInterface: NewChildInstance(),
	}
}

