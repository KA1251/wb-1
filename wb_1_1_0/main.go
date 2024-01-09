package main

import "fmt"

// Структура Human
type Human struct {
	Name string
}

// Структура Action, встроенная в Human
type Action struct {
	Human      // Встроенная структура Human
	hobby_name string
}

// Метод встроенной структуры Action
func (a *Human) DoHobby(action_name string) {
	fmt.Print(a.Name, " is doing:", action_name)
}

func main() {
	// Создание объекта Human с встроенной структурой Action
	person := Action{
		Human{
			Name: "Kirill",
		},
		"programming",
	}

	// Использование метода DoSomething
	person.DoHobby(person.hobby_name)

}
