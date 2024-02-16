package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/osmarjr94/capital-gain/cmd/api/models"
)

type OperationService struct {
	repo *OperationRepository
}

func NewOperationService(repo *OperationRepository) *OperationService {
	return &OperationService{repo: repo}
}

func (s *OperationService) ProcessOperations(operations []models.Operation) ([]models.TaxResult, error) {
	// Implemente a lógica para processar as operações e calcular os impostos.
	var taxResults []models.TaxResult

	// Sua lógica de processamento de operações aqui...

	return taxResults, nil
}

func main() {
	repo := &OperationRepository{} // Instancie seu repositório com as dependências necessárias.
	service := NewOperationService(repo)

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

		taxResults, err := service.ProcessOperations(operations)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error processing operations: %v\n", err)
			continue
		}

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
