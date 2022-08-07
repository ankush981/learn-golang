package main

import "fmt"

// Let's imagine an organization with different types of income streams
// Each income type will have the follwing two methods
type Income interface {
	calculate() int
	source() string
}

// And now, one of the income stream is defined
type FixedBillingProject struct {
	projectName string
	bidAmount   int
}

// Another income stream, that belonging to a Time & Materials project
type TimeAndMaterialProject struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

// Now, let's write receiver functions so that each of the income streams becomes
// an Income type by satisfying that interface
func (fb FixedBillingProject) calculate() int {
	return fb.bidAmount
}

func (fb FixedBillingProject) source() string {
	return fb.projectName
}

func (tm TimeAndMaterialProject) calculate() int {
	return tm.hourlyRate * tm.noOfHours
}

func (tm TimeAndMaterialProject) source() string {
	return tm.projectName
}

// Next, a utility function to calculate the organization's total income
func netIncome(incomes []Income) int {
	total := 0
	for _, income := range incomes {
		total += income.calculate()
	}
	return total
}

func main() {
	project1 := FixedBillingProject{
		projectName: "Fixed Billing Project 1",
		bidAmount:   200000,
	}
	project2 := FixedBillingProject{
		projectName: "Fixed Billing Project 2",
		bidAmount:   200000,
	}
	project3 := TimeAndMaterialProject{
		projectName: "T&M Project 1",
		noOfHours:   340,
		hourlyRate:  80,
	}

	// Having implemented the Income interface, the projects
	// have also become types of Income
	incomeStreams := []Income{project1, project2, project3}
	fmt.Println("Total income: ", netIncome(incomeStreams))
}
