package common

type TaxSlab float64

const (
	TaxSlabRate_5_PERCENT  TaxSlab = 0.05
	TaxSlabRate_10_PERCENT TaxSlab = 0.10
	TaxSlabRate_20_PERCENT TaxSlab = 0.20
	TaxSlabRate_30_PERCENT TaxSlab = 0.30
)

//var (
//	TaxSlab_name = map[float64]string{
//		0.05: "5_PERCENT",
//		0.10: "10_PERCENT",
//		0.20: "20_PERCENT",
//		0.30: "30_PERCENT",
//	}
//	TaxSlab_value = map[string]float64{
//		"5_PERCENT":  0.05,
//		"10_PERCENT": 0.10,
//		"20_PERCENT": 0.20,
//		"30_PERCENT": 0.30,
//	}
//)

type TaxSlabIncomeLowerBound float64

const (
	TaxSlabIncomeLowerBound_5_PERCENT  TaxSlabIncomeLowerBound = 500000
	TaxSlabIncomeLowerBound_10_PERCENT TaxSlabIncomeLowerBound = 1000000
	TaxSlabIncomeLowerBound_20_PERCENT TaxSlabIncomeLowerBound = 2000000
	TaxSlabIncomeLowerBound_30_PERCENT TaxSlabIncomeLowerBound = 2500000
)

const Eighty80CExemption = 150000

const Eighty80CCDNPSExemption = 50000
