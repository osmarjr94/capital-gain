package repositories

import (
	"github.com/osmarjr94/capital-gain/cmd/api/models"
)

type OperationRepository struct {
	// Aqui você pode adicionar campos relevantes para o seu repositório, como uma conexão com o banco de dados, por exemplo.
}

func (r *OperationRepository) Save(operations []models.Operation) error {
	// Implemente a lógica para salvar as operações, se necessário.
	return nil
}
