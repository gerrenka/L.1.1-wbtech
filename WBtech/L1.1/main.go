package main

import (
	"fmt"
	)

type Human struct{
	name string
	age int
}

type Action struct{
	Human
	id int
}

func (h Human) do(){
	fmt.Println("I'm", h.name, "and I'm", h.age, "years old.")
}

func main() {
	Action := Action{Human{"Ivan", 25}, 1}
	Action.do()
}

