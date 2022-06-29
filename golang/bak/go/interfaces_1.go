package main

import (
	"fmt"
	"time"
)

type TeamMember interface {
	PrintName()
	PrintDetails()
}

type Employee struct {
	FirstName, LastName string
	Dob                 time.Time
	JobTitle, Location  string
}

func (e Employee) PrintName() {
	fmt.Printf("\n%s %s\n", e.FirstName, e.LastName)
}

func (e Employee) PrintDetails() {
	fmt.Printf("Date of Birth: %s, Job: %s, Location: %s\n", e.Dob.String(), e.JobTitle, e.Location)
}

type Developer struct {
	Employee //type embedding for composition
	Skills   []string
}

type Emp1 struct {
	name string
}

func (e Emp1) PrintName() {
	fmt.Println("i m emp1")
}

type D struct {
	
}

func main() {
	steve := Developer{
		Employee{
			"Steve",
			"John",
			time.Date(1990, time.February, 17, 0, 0, 0, 0, time.UTC),
			"Software Engineer",
			"San Fancisco",
		},
		Emp1{
			"Goutham"
		}
		[]string{"Go", "Docker", "Kubernetes"},
	}
}
