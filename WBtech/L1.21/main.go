package main

import "fmt"

type EuropeanSocket interface {
	GetVoltage() int 
}

type EuropeanPowerSocket struct{}

func (e *EuropeanPowerSocket) GetVoltage() int {
	return 220 
}

type AmericanPlug struct{}

func (a *AmericanPlug) ConnectTo110V() {
	fmt.Println("Подключено к 110 Вольт")
}

type Adapter struct {
	socket EuropeanSocket 
	plug   *AmericanPlug  
}

func NewAdapter(socket EuropeanSocket, plug *AmericanPlug) *Adapter {
	return &Adapter{
		socket: socket,
		plug:   plug,
	}
}


func (a *Adapter) Use() {
	voltage := a.socket.GetVoltage() 
	if voltage == 220 {
		fmt.Println("Преобразование 220 Вольт в 110 Вольт...")
		a.plug.ConnectTo110V() 
	} else {
		fmt.Println("Неверное напряжение!")
	}
}

func main() {

	europeanSocket := &EuropeanPowerSocket{}

	americanPlug := &AmericanPlug{}


	adapter := NewAdapter(europeanSocket, americanPlug)


	adapter.Use()
}
