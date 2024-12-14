package decorator

import "go-design-patterns/common"

// Eighty80CDecorator 80C Exemption
type Eighty80CDecorator struct {
	AnnualIncome IAnnualIncome
}

func (ec *Eighty80CDecorator) ComputeNetIncome(income float64) float64 {
	taxableIncome := income - float64(common.Eighty80CExemption)
	return ec.AnnualIncome.ComputeNetIncome(taxableIncome)
}

// Eighty80CCDNPSDecorator NPS Exemption
type Eighty80CCDNPSDecorator struct {
	AnnualIncome IAnnualIncome
}

func (nps *Eighty80CCDNPSDecorator) ComputeNetIncome(income float64) float64 {
	taxableIncome := income - float64(common.Eighty80CCDNPSExemption)
	return nps.AnnualIncome.ComputeNetIncome(taxableIncome)
}
