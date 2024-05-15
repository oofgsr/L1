package main

import (
	"fmt"
)

type Human struct {
	name, surname, gender, eyeColor string
	weight, height                  float32
	age                             uint8
}

func (human *Human) GetName() string {
	return human.name
}

func (human *Human) GetSurname() string {
	return human.surname
}

func (human *Human) GetGender() string {
	return human.gender
}

func (human *Human) GetEyeColor() string {
	return human.eyeColor
}

func (human *Human) GetWeight() float32 {
	return human.weight
}

func (human *Human) GetHeight() float32 {
	return human.height
}

func (human *Human) GetAge() uint8 {
	return human.age
}

type Action struct {
	Human
}

func (action *Action) ShowHuman() {
	fmt.Println("Name: " + action.GetName() + "\tsurname: " + action.GetSurname() + "\tgender: " + action.GetGender() + "\teye color: " + action.GetEyeColor() + "\tweight: " + fmt.Sprint(action.GetWeight()) + "\theight: " + fmt.Sprint(action.GetHeight()) + "\tage: " + fmt.Sprint(action.GetAge()))

}

func main() {

	var tom = Human{name: "Grigory",
		surname: "Filyaev", gender: "man",
		eyeColor: "black",
		weight:   85.5, height: 175.6, age: 30}

	var act = Action{tom}

	act.ShowHuman()

}
