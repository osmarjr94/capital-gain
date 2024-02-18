package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/osmarjr94/capital-gain/cmd/api/internal/models"
)

func main() {
	// Read input from stdin
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// Unmarshal JSON input
	var operations []models.Operation
	if err := json.Unmarshal(bytes, &operations); err != nil {
		log.Fatalf("Failed to unmarshal input: %v", err)
	}

	// Repository
	repository := &OperationRepository{}

	// Service
	service := NewOperationService(repository)

	// Controller
	controller := NewOperationController(service)

	// Process operations
	for _, op := range operations {
		if err := controller.HandleOperation(op); err != nil {
			log.Fatalf("Error processing operation: %v", err)
		}
	}
}
