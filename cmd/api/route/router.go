package route

import (
	"echo"
	"net/http"

	"github.com/osmarjr94/capital-gain/cmd/api/internal/models"
)

func RegisterRoutes(e *echo.Echo, controller *OperationController) {
	e.POST("/operation", func(c echo.Context) error {
		var operation []models.Operation

		if err := c.Bind(&operation); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao decodificar o corpo da solicitação"})
		}

		tax, err := controller.HandleOperation(operation)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao processar a operação"})
		}

		return c.JSON(http.StatusOK, map[string]int{"tax": tax})
	})
}
