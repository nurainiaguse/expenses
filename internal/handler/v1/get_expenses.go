package v1

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/nurainiaguse/expenses/api/v1/models"
	"github.com/nurainiaguse/expenses/api/v1/operations"
)

// NewGetExpensesHandler is the implementation of Get Expenses Handler
func NewGetExpensesHandler(params operations.GetExpensesParams) middleware.Responder {
	// dateFrom := swag.StringValue(params.DateFrom)
	// dateTo := swag.StringValue(params.DateTo)

	count := int64(1)
	expense := models.Expense{
		Amount: 12.0,
		ID:     "1234",
	}

	s := make([]*models.Expense, 1)
	s[0] = &expense

	expenses := models.Expenses{
		Count:    &count,
		Expenses: s,
	}
	return operations.NewGetExpensesOK().WithPayload(&expenses)
}
