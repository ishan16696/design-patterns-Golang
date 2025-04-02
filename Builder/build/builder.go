package build

import (
	"fmt"
)

type Employee struct {
	empID int
	name  string
	tech  []string
}

type EmployeeBuilder interface {
	SetName(string) EmployeeBuilder
	SetTechStack([]string) EmployeeBuilder
	SetEmpID(int) EmployeeBuilder
	BuildDev() *Employee
}

type Developer struct {
	dev Employee
}

func (d *Developer) SetName(name string) EmployeeBuilder {
	d.dev.name = name
	return d
}

func (d *Developer) SetTechStack(tech []string) EmployeeBuilder {
	d.dev.tech = tech
	return d
}

func (d *Developer) SetEmpID(empID int) EmployeeBuilder {
	d.dev.empID = empID
	return d
}

func (d *Developer) BuildDev() *Employee {
	return &Employee{
		empID: d.dev.empID,
		tech:  d.dev.tech,
		name:  d.dev.name,
	}
}

func (e *Employee) AssignProject() {
	fmt.Printf("Welcome %s\n", e.name)
	fmt.Printf("Assigning the project acc to tech: %v\n", e.tech)
}
