package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type PaymentMethod interface {
	Pay(float32) string
}

func GetPayMethod(method string) (PaymentMethod, error) {
	switch method {
	case "cash":
		return new(PayInCash), nil
	case "card":
		return new(PayWithCard), nil
	default:
		return nil, errors.New(fmt.Sprintf("%s: no such pay method", method))
	}
}

type PayInCash struct{}

func (pic *PayInCash) Pay(amount float32) string {
	return fmt.Sprintf("$%.2f paid in cash", amount)
}

type PayWithCard struct{}

func (pvc *PayWithCard) Pay(amount float32) string {
	return fmt.Sprintf("$%.2f paid via debit card", amount)
}

func example_factory() {
	if payment, err := GetPayMethod("cash"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(payment.Pay(150))
	}

	if payment, err := GetPayMethod("card"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(payment.Pay(150.00))
	}

	if payment, err := GetPayMethod("credit card"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(payment.Pay(150.00))
	}
}
