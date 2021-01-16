package v1

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/nurainiaguse/expenses/api/v1/models"
	"github.com/nurainiaguse/expenses/api/v1/operations"
)

// NewGetExpenseByIDHandler is the implementation of Get Expenses By ID Handler
func NewGetExpenseByIDHandler(params operations.GetExpenseByIDParams) middleware.Responder {

	expense := models.Expense{
		Amount: 12.0,
		ID:     params.ExpenseID,
	}

	return operations.NewGetExpenseByIDOK().WithPayload(&expense)
}
