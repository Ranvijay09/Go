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

func (job *TaxIncludedPriceJob) LoadPrices() error {
	file, err := os.Open("prices.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		convertedPrice, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return err
		}
		job.InputPrices = append(job.InputPrices, convertedPrice)
	}
	return scanner.Err()
}

func (job *TaxIncludedPriceJob) Process() {
	err := job.LoadPrices()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(job.InputPrices)

	resultMap := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		resultMap[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(resultMap)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{},
	}
}
