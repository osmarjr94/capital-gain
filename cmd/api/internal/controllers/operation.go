package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/osmarjr94/capital-gain/cmd/api/models"
)
type OperationController struct {
	service *OperationService
}

func NewOperationController(service &OperationService) *OperationController {
	return &OperationController{service: service}
}

func (c *OperationController) HandleOperation(operation models.Operation) (int, error) {
	return oc.service.ProcessOperation(operation)
}