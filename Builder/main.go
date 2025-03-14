package main

import "design-patterns-Golang/Builder/build"

func main() {
	fact := build.NewEmployeeBuild()

	emp1 := fact.SetName("Ishan").SetTechStack([]string{"C++", "Docker", "Go"}).BuildDev()
	emp1.AssignProject()

	emp2 := fact.SetName("Ayush").SetFields([]string{"Finance", "BusinessConsultant"}).BuildCon()
	emp2.AssignProject()
}
