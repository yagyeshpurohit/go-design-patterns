package decorator

import "go-design-patterns/common"

type BasicAnnualIncome struct {
	AnnualIncome IAnnualIncome
}

func (bai *BasicAnnualIncome) ComputeNetIncome(income float64) float64 {
	tax := computeTax(income)
	return income - tax
}

func computeTax(income float64) float64 {
	var totalTax float64 = 0
	var taxSlabLowerBound float64
	//base condition
	if income <= float64(common.TaxSlabIncomeLowerBound_5_PERCENT) {
		taxSlabLowerBound = 0
		//fmt.Printf("Tax for slab %d - %d: %d \n", int64(income), int64(taxSlabLowerBound), int64(totalTax))
		return 0
	}
	if income > float64(common.TaxSlabIncomeLowerBound_30_PERCENT) {
		taxSlabLowerBound = float64(common.TaxSlabIncomeLowerBound_30_PERCENT)
		totalTax = (income-taxSlabLowerBound)*float64(common.TaxSlabRate_30_PERCENT) + computeTax(taxSlabLowerBound)
	} else if income > float64(common.TaxSlabIncomeLowerBound_20_PERCENT) {
		taxSlabLowerBound = float64(common.TaxSlabIncomeLowerBound_20_PERCENT)
		totalTax = (income-taxSlabLowerBound)*float64(common.TaxSlabRate_20_PERCENT) + computeTax(taxSlabLowerBound)
	} else if income > float64(common.TaxSlabIncomeLowerBound_10_PERCENT) {
		taxSlabLowerBound = float64(common.TaxSlabIncomeLowerBound_10_PERCENT)
		totalTax = (income-taxSlabLowerBound)*float64(common.TaxSlabRate_10_PERCENT) + computeTax(taxSlabLowerBound)
	} else if income > float64(common.TaxSlabIncomeLowerBound_5_PERCENT) {
		taxSlabLowerBound = float64(common.TaxSlabIncomeLowerBound_5_PERCENT)
		totalTax = (income-taxSlabLowerBound)*float64(common.TaxSlabRate_5_PERCENT) + computeTax(taxSlabLowerBound)
	}
	//fmt.Printf("Tax upto slab %d - %d: %d \n", int64(income), int64(taxSlabLowerBound), int64(totalTax))
	return totalTax
}
