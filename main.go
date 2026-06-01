package main

import (
	"fmt"

	"gotest.com/price-calculator/filemanager"
	"gotest.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
	// fmt.Println(result)
	// fmt.Println("--------------------------------")
	// fmt.Println(result[0])
	// fmt.Println(result[0.07])
	// fmt.Println(result[0.1])
	// fmt.Println(result[0.15])
}
