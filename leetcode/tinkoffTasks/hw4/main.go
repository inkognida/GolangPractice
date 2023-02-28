package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	school string
}

func (h Human) SayHi() {
	fmt.Println("Human", h.name, h.age)
}

func (s Student) SayHi(arg string) {
	fmt.Println("Student", s.name, s.age, arg)
}

type Sayer interface {
	SayHi()
}

func Say(s Sayer) {
	s.SayHi()
}

type hideStudent struct {
	Sayer
}

func main() {

	s := Student{
		Human:  Human{name: "Bob", age: 19},
		school: "21",
	}

	h := Human{
		name: "Aboba",
		age:  21,
	}

	Say(h)

	s.SayHi("embedding")
	s.Human.SayHi()

	w := struct {
		Human
		id int
	}{
		Human: Human{name: "German", age: 18},
		id:    1,
	}
	w.SayHi()

	/* hide name and age for hs, it has Human struct which
	implement Sayer interface and has own methods from Human
	*/

	hs := hideStudent{Human{
		name: "a",
		age:  1,
	}}

	hs.SayHi()
	Say(hs)
}

// TODO get more info about Embedding, logger, go mod vendor
