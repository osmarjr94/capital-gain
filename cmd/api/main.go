package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/osmarjr94/capital-gain/cmd/api/internal/models"
)

func main() {

	repository := NewOperationRepository()

	service := NewOperationService(repository)

	controller := NewOperationController(service)

	bytes, err := io.ReadAll(io.Reader(os.Stdin))
	if err != nil {
		log.Fatalf("Erro ao ler a entrada padrão: %v", err)
	}

	var operations []models.Operation
	if err := json.Unmarshal(bytes, &operations); err != nil {
		log.Fatalf("Erro ao decodificar as operações do JSON: %v", err)
	}

	for _, op := range operations {
		tax, err := controller.HandleOperation(op)
		if err != nil {
			log.Fatalf("Erro ao processar a operação: %v", err)
		}
		fmt.Printf("Imposto calculado para a operação %+v: %d\n", op, tax)
	}
}
