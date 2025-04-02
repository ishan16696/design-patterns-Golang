package main

import "design-patterns-Golang/Builder/build"

func main() {
	dev := build.Developer{}

	emp1 := dev.SetName("Ishan").SetTechStack([]string{"C++", "Docker", "Go"}).SetEmpID(2).BuildDev()
	emp1.AssignProject()
}
