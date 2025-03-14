package factory

import (
	"fmt"
)

type name string

type PersonConfig struct {
	First    name
	Last     name
	Availale bool
}

type person struct {
	first name
	last  name
}

type secretAgent struct {
	man         person
	isAvailable bool
}

type Factory interface {
	NewPerson() *person
	NewAgent() *secretAgent
}

type factoryImpl PersonConfig

func NewFactory(cfg PersonConfig) Factory {
	var f = factoryImpl(cfg)
	return &f
}

func (f *factoryImpl) NewPerson() *person {
	return &person{
		first: f.First,
		last:  f.Last,
	}
}

func (f *factoryImpl) NewAgent() *secretAgent {
	return &secretAgent{
		man: person{
			first: f.First,
			last:  f.Last,
		},
		isAvailable: f.Availale,
	}
}

func (s *secretAgent) Speak() {
	fmt.Println("I am agent", s.man.first, s.man.last, "- on a mission...")
	if s.isAvailable {
		fmt.Println("I'm availale...")
	} else {
		fmt.Println("I'm not available...")
	}
}

func (p *person) Speak() {
	fmt.Println("I am", p.first, p.last, "- the Person speak")
}
