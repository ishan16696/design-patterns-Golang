package build

import (
	"fmt"
)

type Developer struct {
	name string
	tech []string
}

type Consultant struct {
	name   string
	fields []string
}

type EmployeeBuild struct {
	name   string
	tech   []string
	fields []string
}

func (e *EmployeeBuild) SetName(name string) *EmployeeBuild {
	e.name = name
	return e
}

func (e *EmployeeBuild) SetTechStack(tech []string) *EmployeeBuild {
	e.tech = tech
	return e
}

func (c *EmployeeBuild) SetFields(f []string) *EmployeeBuild {
	c.fields = f
	return c
}

func (e *Developer) AssignProject() {
	fmt.Printf("Welcome %s\n", e.name)
	fmt.Printf("Assigning the project acc to tech: %v\n", e.tech)
}

func (c *Consultant) AssignProject() {
	fmt.Printf("Welcome %s\n", c.name)
	fmt.Printf("Assigning the project acc to fileds: %v\n", c.fields)
}

func NewEmployeeBuild() *EmployeeBuild {
	return &EmployeeBuild{}
}

func (e *EmployeeBuild) BuildDev() *Developer {
	return &Developer{
		name: e.name,
		tech: e.tech,
	}
}

func (e *EmployeeBuild) BuildCon() *Consultant {
	return &Consultant{
		name:   e.name,
		fields: e.fields,
	}
}
