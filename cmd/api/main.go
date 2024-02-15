package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Operation representa uma operação do mercado financeiro.
type Operation struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

// TaxResult representa o resultado do imposto pago para uma operação.
type TaxResult struct {
	Tax int `json:"tax"`
}

// OperationRepository é responsável por lidar com o armazenamento ou recuperação de operações.
type OperationRepository struct{}

// GetOperations lê as operações da entrada padrão e retorna uma lista de operações.
func (r *OperationRepository) GetOperations() ([]Operation, error) {
	var operations []Operation

	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil || line == "" {
			break
		}

		var ops []Operation
		err = json.Unmarshal([]byte(line), &ops)
		if err != nil {
			return nil, err
		}

		operations = append(operations, ops...)
	}

	return operations, nil
}

// OperationService é responsável por processar as operações e calcular o imposto.
type OperationService struct{}

// CalculateTax calcula o imposto pago para cada operação.
func (s *OperationService) CalculateTax(operations []Operation) []TaxResult {
	var taxResults []TaxResult

	var boughtShares int
	for _, op := range operations {
		var tax int
		if op.Operation == "sell" {
			if op.Quantity > boughtShares {
				tax = (boughtShares * int(op.UnitCost) * 100) / 1000
				boughtShares = 0
			} else {
				tax = (op.Quantity * int(op.UnitCost) * 100) / 1000
				boughtShares -= op.Quantity
			}
		} else {
			boughtShares += op.Quantity
		}

		taxResults = append(taxResults, TaxResult{Tax: tax})
	}

	return taxResults
}

// OperationController é responsável por receber a entrada e retornar os resultados do imposto.
type OperationController struct {
	repo    *OperationRepository
	service *OperationService
}

// NewOperationController cria uma nova instância de OperationController.
func NewOperationController(repo *OperationRepository, service *OperationService) *OperationController {
	return &OperationController{
		repo:    repo,
		service: service,
	}
}

// ProcessOperations processa as operações recebidas e retorna os resultados do imposto.
func (c *OperationController) ProcessOperations() {
	operations, err := c.repo.GetOperations()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting operations: %v\n", err)
		os.Exit(1)
	}

	taxResults := c.service.CalculateTax(operations)

	for _, result := range taxResults {
		taxResultJSON, err := json.Marshal([]TaxResult{result})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshalling JSON: %v\n", err)
			continue
		}
		fmt.Println(string(taxResultJSON))
	}
}

func main() {
	repo := &OperationRepository{}
	service := &OperationService{}
	controller := NewOperationController(repo, service)
	controller.ProcessOperations()
}
