package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type normalEmployee struct{
	empId int
	basicPay int
}

type excellentEmployee struct {
	emId int
	basicPay int
	bonus int
}

func (n normalEmployee)CalculateSalary() int {
	return n.basicPay
}

func (e excellentEmployee)CalculateSalary() int {
	return e.basicPay + e.bonus
}

func totalExpense(s []SalaryCalculator) int {
	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
	}
	fmt.Printf("雇员总开支为: %d$", expense)

	return expense
}

func main(){
	jst := excellentEmployee{emId:12345, basicPay:8000, bonus:5000}
	heshan := normalEmployee{empId:23456, basicPay:8000}
	longqin := excellentEmployee{emId:34567, basicPay:8000, bonus:3000}
	employees := []SalaryCalculator{jst, heshan, longqin}
	totalExpense(employees)
}
