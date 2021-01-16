package server

import (
	"log"

	"github.com/go-openapi/loads"
	v1API "github.com/nurainiaguse/expenses/api/v1"
	"github.com/nurainiaguse/expenses/api/v1/operations"
	v1Handler "github.com/nurainiaguse/expenses/internal/handler/v1"
)

// NewServer creates an instance of a server
func NewServer() *v1API.Server {
	swaggerSpec, err := loads.Analyzed(v1API.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewExpensesAPI(swaggerSpec)
	server := v1API.NewServer(api)

	api.GetExpensesHandler = operations.GetExpensesHandlerFunc(v1Handler.NewGetExpensesHandler)
	api.GetExpenseByIDHandler = operations.GetExpenseByIDHandlerFunc(v1Handler.NewGetExpenseByIDHandler)

	return server
}
