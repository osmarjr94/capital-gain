package controllers

/*import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/osmarjr94/capital-gain/cmd/api/models"
)

type OperationController struct{}

func (c *OperationController) CalculateTaxes(operations []models.Operation) []models.TaxResult {
	var taxResults []models.TaxResult

	var boughtShares int
	var weightedAverage float64

	for _, op := range operations {
		if op.Operation == "buy" {
			boughtShares += op.Quantity
			weightedAverage = ((float64(boughtShares-op.Quantity) * weightedAverage) + (float64(op.Quantity) * op.UnitCost)) / float64(boughtShares)
			taxResults = append(taxResults, models.TaxResult{Tax: 0})
		} else if op.Operation == "sell" {
			if op.Quantity > boughtShares {
				continue
			}

			var tax int
			totalCost := op.UnitCost * float64(op.Quantity)
			if totalCost <= 20000 {
				taxResults = append(taxResults, models.TaxResult{Tax: 0})
				continue
			}

			if op.UnitCost > weightedAverage {
				profit := totalCost - (weightedAverage * float64(op.Quantity))
				tax = int(profit * 0.20)
			}
			taxResults = append(taxResults, models.TaxResult{Tax: tax})

			boughtShares -= op.Quantity
		}
	}

	return taxResults
}

func operationScanner() {
	var input []string

	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil || line == "" {
			break
		}
		input = append(input, line)
	}

	oc := OperationController{}
	for _, line := range input {
		var operations []models.Operation
		if err := json.Unmarshal([]byte(line), &operations); err != nil {
			fmt.Println(os.Stderr, "Error parsing JSON: %v\n", err)
			continue
		}

		taxResults := oc.CalculateTaxes(operations)

		taxResultsJSON, err := json.Marshal(taxResults)
		if err != nil {
			fmt.Println(os.Stderr, "Error marshalling JSON: %v\n", err)
			continue
		}

		fmt.Println(string(taxResultsJSON))
	}
}
*/

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/osmarjr94/capital-gain/cmd/api/models"
)

type OperationController struct{}

func (c *OperationController) CalculateTaxes(operations []models.Operation) *[]models.TaxResult {
	var taxResults []models.TaxResult

	var boughtShares int
	var weightedAverage float64

	for _, op := range operations {
		if op.Operation == "buy" {
			boughtShares += op.Quantity
			weightedAverage = ((float64(boughtShares-op.Quantity) * weightedAverage) + (float64(op.Quantity) * op.UnitCost)) / float64(boughtShares)
			taxResults = append(taxResults, models.TaxResult{Tax: 0})
		} else if op.Operation == "sell" {
			if op.Quantity > boughtShares {
				continue
			}

			var tax int
			totalCost := op.UnitCost * float64(op.Quantity)
			if totalCost <= 20000 {
				taxResults = append(taxResults, models.TaxResult{Tax: 0})
				continue
			}

			if op.UnitCost > weightedAverage {
				profit := totalCost - (weightedAverage * float64(op.Quantity))
				tax = int(profit * 0.20)
			}
			taxResults = append(taxResults, models.TaxResult{Tax: tax})

			boughtShares -= op.Quantity
		}
	}

	return &taxResults
}

func operationController() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var operations []models.Operation
		if err := json.Unmarshal([]byte(line), &operations); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
			continue
		}

		oc := OperationController{}
		taxResults := oc.CalculateTaxes(operations)

		taxResultsJSON, err := json.Marshal(taxResults)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshalling JSON: %v\n", err)
			continue
		}

		fmt.Println(string(taxResultsJSON))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading standard input: %v\n", err)
		os.Exit(1)
	}
}
