package entity

import (
	"github.com/LexBokun/transaction-service/internal/pkg/lists/currency"
	"time"

	"google.golang.org/genproto/googleapis/type/decimal"
)

type Money struct {
	Amount      decimal.Decimal // Введённая сумма
	SuccessPaid decimal.Decimal // Сумма + комиссия
	Fee         decimal.Decimal // Комиссия

	Currency currency.Currency // валюта Фиат/Крипта
}

type Operation struct {
	Account   string
	Money     Money
	CreatedAt time.Time
	UpdatedAt time.Time
}
