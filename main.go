package main

import (
	"fmt"

	"gotest.com/price-calculator/filemanager"
	"gotest.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool) // create a channel to send a signal to the main thread that the job is done
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.0f.json", taxRate*100))
		//cmd := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		go priceJob.Process(doneChans[index]) // process the price job in a goroutine

		// if err != nil {
		// 	fmt.Println("Error processing price job: ", err)
		// 	return
		// }
	}

	// wait for all the jobs to be done
	for _, doneChan := range doneChans {
		<-doneChan
	}
	// fmt.Println(result)
	// fmt.Println("--------------------------------")
	// fmt.Println(result[0])
	// fmt.Println(result[0.07])
	// fmt.Println(result[0.1])
	// fmt.Println(result[0.15])
}
