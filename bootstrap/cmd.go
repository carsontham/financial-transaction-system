package bootstrap

import (
	"financial-transaction-system/app/adapter/http/rest"
	repository "financial-transaction-system/app/repository/postgresql"
	"financial-transaction-system/app/usecase"
	"log"
)

func Run() {
	s := NewServer(":3000")
	repo := repository.NewFinancialTransactionRepository(GetDBConnection)
	service := usecase.NewService(repo)
	validator, err := rest.NewCustomValidator()
	if err != nil {
		log.Println("failed building custom validator")
	}
	s.SetUpRoutes(service, validator)
	s.RunServer()
}
