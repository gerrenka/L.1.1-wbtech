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
		fmt.Println(  "я здесь")
	}

func main() {
	Action := Action{Human{"Ivan", 25}, 1}

	Action.do()
}

