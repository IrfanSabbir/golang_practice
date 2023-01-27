package interfaceExample

import (
	"fmt"
	"log"
)

type Employee interface {
	PrintName(name string)
	PrintSalary(salary int, tax int) int
}

type Emp int

func (e Emp) PrintName(name string) {
	fmt.Println("name of employee ", name)
}

func (e Emp) PrintSalary(salary int, tax int) int {
	value := (salary * tax) / 100
	return value
}

func InterfaceExample() {
	log.Println()
	fmt.Println("Interface Examples.................\n")

	var e1 Employee
	e1 = Emp(1)
	e1.PrintName("Abdul based")
	fmt.Println("Salary is ", e1.PrintSalary(2000, 10))

	log.Println()
	return
}
