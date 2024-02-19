package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osmarjr94/capital-gain/cmd/api/models"
)

type OperationController struct {
	service *OperationServiceInterface
}

func NewOperationController(service &OperationServiceInterface) *OperationController {
	return &OperationController{service: service}
}

func (oc *OperationController) HandleOperation(c echo.Context) error {
	var operation Operation

	if err := c.Bind(&operation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao decodificar o corpo da solicitação"})
	}

	tax, err := oc.service.ProcessOperation(operation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao processar a operação"})
	}

	return c.JSON(http.StatusOK, TaxResult{Tax: tax})
}