package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		Prices:  []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)
	for _, price := range job.Prices {
		result[fmt.Sprintf("%.0f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)

}
