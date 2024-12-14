package lab6

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Height float64
}

func NewPerson(name string, age int, height float64) Person {
	return Person{
		Name:   name,
		Age:    age,
		Height: height,
	}
}

func (p Person) GetInfo() string {
	return fmt.Sprintf("Имя: %s, Возраст: %d, Рост: %.2f м", p.Name, p.Age, p.Height)
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p *Person) UpdateName(newName string) {
	p.Name = newName
}

func Runlab6() {

	person := NewPerson("Анна", 22, 1.68)

	fmt.Println(person.GetInfo())

	fmt.Printf("Совершеннолетний: %v\n", person.IsAdult())

	person.UpdateName("Екатерина")
	fmt.Println(person.GetInfo())
}
