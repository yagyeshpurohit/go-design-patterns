package main

import (
	"fmt"
	"go-design-patterns/structural/decorator"
)

func main() {

	incomeHavingNPSAnd80C := &decorator.Eighty80CCDNPSDecorator{
		AnnualIncome: &decorator.Eighty80CDecorator{
			AnnualIncome: &decorator.BasicAnnualIncome{},
		},
	}

	incomePlain := &decorator.BasicAnnualIncome{}

	fmt.Printf("Plain income : %f \n", incomePlain.ComputeNetIncome(2600000))
	fmt.Printf("Income having NPS and 80C Exemption: %f \n", incomeHavingNPSAnd80C.ComputeNetIncome(2600000))
}
