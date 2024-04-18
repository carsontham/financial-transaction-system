package bootstrap

import (
	repository "financial-transaction-system/app/repository/postgresql"
	"financial-transaction-system/app/usecase"
)

func Run() {
	s := NewServer(":3000")
	repo := repository.NewFinancialTransactionRepository(GetDBConnection)
	service := usecase.NewService(repo)
	s.SetUpRoutes(service)
	s.RunServer()
}
