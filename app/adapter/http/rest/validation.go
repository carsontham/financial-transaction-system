package rest

import (
	"github.com/go-playground/validator/v10"
	"log"
	"strconv"
)

func NewCustomValidator() (*validator.Validate, error) {
	v := validator.New()
	if err := SetUpCustomValidations(v); err != nil {
		log.Println("error setting custom validator")
		return nil, err
	}
	return v, nil
}

func SetUpCustomValidations(v *validator.Validate) error {
	if err := v.RegisterValidation("valid_balance", validBalance); err != nil {
		return err
	}
	if err := v.RegisterValidation("valid_amount", validAmount); err != nil {
		return err
	}
	return nil
}

func validBalance(fl validator.FieldLevel) bool {
	balanceString := fl.Field().String()
	_, err := strconv.ParseFloat(balanceString, 64)
	if err != nil {
		log.Println("invalid balance - does not contain only numbers ")
		return false
	}
	log.Println("valid balance")
	return true
}

func validAmount(fl validator.FieldLevel) bool {
	balanceString := fl.Field().String()
	balance, _ := strconv.ParseFloat(balanceString, 64)
	if balance < 0 {
		log.Println("invalid amount - < 0")
		return false
	}
	log.Println("valid amount >= 0")
	return true
}
