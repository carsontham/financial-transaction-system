package bootstrap

import (
	"financial-transaction-system/app/adapter/http/handlers"
	"financial-transaction-system/app/usecase"
)

func (s *Server) SetUpRoutes(service *usecase.Service) {
	// TODO:
	s.router.Get("/accounts/{account_id}", handlers.GetAccountByID(service))
	s.router.Get("/accounts", handlers.GetAccountByID(service))
	s.router.Get("/transactions", handlers.CreateNewTransaction())

}