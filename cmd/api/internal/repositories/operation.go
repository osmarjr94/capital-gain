package repositories

import (
	"errors"
	"sync"

	"github.com/osmarjr94/capital-gain/cmd/api/models"
)

type OperationRepository struct {
	operations []models.Operation
	mutex      sync.Mutex
}

func NewOperationRepository() *OperationRepository {
	return &OperationRepository{}
}

func (r *OperationRepository) SaveOperation(operation models.Operation) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.operations = append(r.operations, operation)
	return nil
}

func (r *OperationRepository) GetWeightedAverage() (float64, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var totalQuantity int
	var totalCost float64

	for _, op := range r.operations {
		if op.Operation == "buy" {
			totalQuantity += op.Quantity
			totalCost += op.UnitCost * float64(op.Quantity)
		}
	}

	if totalQuantity == 0 {
		return 0, errors.New("no buy operations found")
	}

	return totalCost / float64(totalQuantity), nil
}
