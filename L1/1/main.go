package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
	hobby string
}

func (h *Human) SetName(name string) {
	h.name = name
}
func (h *Human) SetAge(age int) {
	h.age = age
}
func (a *Action) SetHobby(hobby string) {
	a.hobby = hobby
}
func (h *Human) GetName() string {
	return h.name
}
func (h *Human) GetAge() int {
	return h.age
}
func (a *Action) GetHobby() string {
	return a.hobby
}

func main() {
	a := Action{}

	a.SetName("Mikhail")
	a.SetAge(18)
	a.SetHobby("Programming")

	fmt.Println(a.GetName())
	fmt.Println(a.GetAge())
	fmt.Println(a.GetHobby())
}
