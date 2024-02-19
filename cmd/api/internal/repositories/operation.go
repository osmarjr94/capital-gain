package repositories

import (
	"github.com/osmarjr94/capital-gain/cmd/api/models"
)

type OperationRepository struct {
	// Aqui você pode adicionar qualquer dependência necessária, como uma conexão com o banco de dados
}

func NewOperationRepository() *OperationRepository {
	return &OperationRepository{}
}

func (or *OperationRepository) SaveOperation(operation []models.Operation) error {
	// Implemente a lógica para salvar a operação no banco de dados
	return nil
}

func (or *OperationRepository) GetWeightedAverage() (float64, error) {
	// Implemente a lógica para obter a média ponderada do banco de dados
	return 0, nil
}
