package main

import "fmt"

func main() {
	taxRates := []float64{0, 10, 20}

	prices := []float64{10, 20, 30}

	resultMap := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxIncludedPrices := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrices[priceIndex] = price + ((price / 100) * taxRate)
		}
		resultMap[taxRate] = taxIncludedPrices
	}
	fmt.Println(resultMap)
}
