package repositories

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/osmarjr94/capital-gain/cmd/api/models"
)

type OperationRepository struct {
	db *sql.DB
	mu sync.RWMutex
}

func NewOperationRepository(db *sql.DB) (*OperationRepository, error) {
	err := createOperationsTable(db)
	if err != nil {
		return nil, err
	}

	return &OperationRepository{db: db}, nil
}

func createOperationsTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS operations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			operation TEXT,
			unit_cost REAL,
			quantity INTEGER
		)
	`)
	return err
}

func (r *OperationRepository) SaveOperation(operation []models.Operation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec("INSERT INTO operations (operation, unit_cost, quantity) VALUES (?, ?, ?)", operation.Operation, operation.UnitCost, operation.Quantity)
	return err
}

func (r *OperationRepository) GetWeightedAverage() (float64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var sum float64
	var totalQuantity int

	rows, err := r.db.Query("SELECT unit_cost, quantity FROM operations WHERE operation = 'buy'")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var unitCost float64
		var quantity int
		err := rows.Scan(&unitCost, &quantity)
		if err != nil {
			return 0, err
		}
		sum += unitCost * float64(quantity)
		totalQuantity += quantity
	}

	if totalQuantity == 0 {
		return 0, errors.New("não há operações de compra")
	}

	return sum / float64(totalQuantity), nil
}
