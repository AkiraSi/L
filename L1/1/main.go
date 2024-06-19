package main

import "fmt"

type Human struct { // структура человека
	Name  string
	Age   int
	Hobby string
}

func (h *Human) vivod() { // вывод информации о человеке (некий геттер всех полей разом)
	fmt.Println(h.Name, h.Age, h.Hobby)
}

func (h *Human) GetAge() int { // геттер возраста
	return h.Age
}

type Action struct { // дочерний класс со своим методом
	Human
	Skill string
}

func (a *Action) DoAction() { // действие дочернего класса
	fmt.Println(a.Skill)
}

func main() {
	action := Action{
		Human: Human{
			Name:  "Ilfat",
			Age:   19,
			Hobby: "lomat-krusit",
		},
		Skill: "coding",
	}
	action.vivod()
	fmt.Println("Возраст:", action.GetAge())
	fmt.Println("Способность:", action.Skill)

	action.DoAction()
}
