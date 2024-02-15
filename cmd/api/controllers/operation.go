package controllers

import (
	"encoding/json"
	"fmt"
	"os"

	"capital-gain-api/cmd/api/models"
)

type OperationProcessor interface {
	ProcessOperations(input []string) ([]models.TaxResult, error)
}

type OperationController struct{}

func (c *OperationController) ProcessOperations(input []string) []models.TaxResult {
	var taxResults []models.TaxResult

	for _, line := range input {
		var operations []models.Operation
		err := json.Unmarshal([]byte(line), &operations)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
			continue
		}

		taxResults = append(taxResults, calculateTax(operations)...)
	}

	return taxResults
}

func calculateTax(operations []models.Operation) []models.TaxResult {
	var results []models.TaxResult
	var boughtShares int

	for _, op := range operations {
		if op.Operation == "buy" {
			boughtShares += op.Quantity
			results = append(results, models.TaxResult{Tax: 0})
		} else if op.Operation == "sell" {
			if op.Quantity > boughtShares {
				continue
			}

			tax := (op.Quantity * op.UnitCost * 100) / 1000
			if op.Quantity > boughtShares {
				tax = (boughtShares * op.UnitCost * 100) / 1000
				boughtShares = 0
			} else {
				boughtShares -= op.Quantity
			}
			results = append(results, models.TaxResult{Tax: tax})
		}
	}

	return results
}

func operationController() {
	var input []string

	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil || line == "" {
			break
		}
		input = append(input, line)
	}

	processor := OperationController{}
	taxResults := processor.ProcessOperations(input)

	taxResultsJSON, err := json.Marshal(taxResults)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(taxResultsJSON))
}
