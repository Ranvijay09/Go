package prices

import (
	"bufio"
	"encoding/json"
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
		return
	}

	resultMap := make(map[string]float64, len(job.InputPrices))
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		resultMap[fmt.Sprintf("%.2f", price)] = taxIncludedPrice
	}
	job.TaxIncludedPrices = resultMap
	WriteJSONToFile(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), *job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{},
	}
}

func WriteJSONToFile(path string, data TaxIncludedPriceJob) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}
