package bootstrap

import "financial-transaction-system/app/adapter/http/handlers"

func (s *Server) SetUpRoutes() {
	// TODO:
	s.router.Get("/accounts/{account_id}", handlers.GetAccountByID())
	s.router.Post("/accounts", handlers.CreateNewAccount())
	s.router.Get("/transactions", handlers.CreateNewTransaction())

}
