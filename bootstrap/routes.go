package bootstrap

import (
	"financial-transaction-system/app/adapter/http/handlers"
	"financial-transaction-system/app/usecase"
	"github.com/go-playground/validator/v10"
)

func (s *Server) SetUpRoutes(service *usecase.Service, validator *validator.Validate) {
	// TODO:
	s.router.Get("/accounts/{account_id:\\d+}", handlers.GetAccountByID(service))
	s.router.Post("/accounts", handlers.CreateNewAccount(service, validator))
	s.router.Get("/transactions", handlers.CreateNewTransaction())

}
