package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64) // create a map with the key as the tax rate and the value as the price with tax
	for _, taxRate := range taxRates {
		priceWithTax := make([]float64, len(prices)) // create a slice with the length of the prices array
		for priceIndex, price := range prices {
			priceWithTax[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = priceWithTax
	}
	fmt.Println(result)
	fmt.Println("--------------------------------")
	fmt.Println(result[0])
	fmt.Println(result[0.07])
	fmt.Println(result[0.1])
	fmt.Println(result[0.15])
}
