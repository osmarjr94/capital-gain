package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"

	"github.com/osmarjr94/capital-gain/cmd/api/models"
)

type OperationService struct {
	repository *OperationRepository
}

func NewOperationService(repository &OperationRepository) *OperationService {
	return &OperationService{repository: repository}
}

func (os *OperationService) ProcessOperation(operation []models.Operation) (int, error) {
	if operation.Operation == "buy" {
		return 0, os.repository.SaveOperation(operation)
	}

	if operation.Operation != "sell" {
		return 0, nil
	}

	totalCost := operation.UnitCost * float64(operation.Quantity)
	if totalCost <= 20000 {
		return 0, os.repository.SaveOperation(operation)
	}

	weightedAverage, err := os.repository.GetWeightedAverage()
	if err != nil {
		return 0, err
	}

	profit := totalCost - (float64(operation.Quantity) * weightedAverage)
	if profit <= 0 {
		return 0, os.repository.SaveOperation(operation)
	}

	tax := int(math.Ceil(profit * 0.2))
	return tax, os.repository.SaveOperation(operation)
}
