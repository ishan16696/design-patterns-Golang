package main

import "design-patterns-Golang/factoryMethod/factory"

func main() {
	cfg := factory.PersonConfig{
		First:    "ishan",
		Last:     "tyagi",
		Availale: false,
	}

	fact := factory.NewFactory(cfg)

	// with factory method, we can construct the instances of both person/agent.
	person := fact.NewPerson()
	agent := fact.NewAgent()

	person.Speak()
	agent.Speak()

	config := factory.PersonConfig{
		First:    "ayush",
		Last:     "tyagi",
		Availale: true,
	}

	//now we have new config, we need to have a new factory to create a object.
	clientFactory := factory.NewFactory(config)
	agent = clientFactory.NewAgent()
	agent.Speak()

}
