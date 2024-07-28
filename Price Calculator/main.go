package main

import (
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.10, 0.02, 0.7}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
