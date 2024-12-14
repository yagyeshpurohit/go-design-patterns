package decorator

type IAnnualIncome interface {
	ComputeNetIncome(float64) float64
}
