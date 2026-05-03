package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Text is the line read from the file
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading the file contents: ", err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Converting price to float failed: ", err)
			file.Close()
			return
		}
		prices[lineIndex] = floatPrice
	}
	file.Close()
	job.InputPrices = prices

}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData() // load the data from the file into the job.InputPrices
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)

}
