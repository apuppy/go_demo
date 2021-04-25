package main

import (
	"fmt"
)

// interface declaration
type Income interface {
	calculate() int
	source() string
}

// struct declaration
type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

// interface implementation for FixedBilling
func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

// interface implementation for TimeAndMaterial
func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income from %s = %d \n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d \n", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "project 1", biddedAmount: 200}
	project2 := FixedBilling{projectName: "project 2", biddedAmount: 500}
	project3 := TimeAndMaterial{projectName: "project 3", noOfHours: 8, hourlyRate: 25}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}
