package currency

import (
	"fmt"

	cryptoEntity "github.com/LexBokun/transaction-service/pkg/api/enums/cryptocurrency/v1"
	fiatEntity "github.com/LexBokun/transaction-service/pkg/api/enums/fiatcurrency/v1"

	"github.com/samber/lo"
)

type Currency struct {
	Fiat   *fiatEntity.Currency
	Crypto *cryptoEntity.Currency
}

func NewFiat(currency fiatEntity.Currency) (Currency, error) {
	if !lo.Contains([]fiatEntity.Currency{
		fiatEntity.Currency_CURRENCY_RUB,
	}, currency) {
		return Currency{}, fmt.Errorf("unsupported fiat currency: %v", currency)
	}
	return Currency{Fiat: &currency}, nil
}

func NewCrypto(currency cryptoEntity.Currency) (Currency, error) {
	if !lo.Contains([]cryptoEntity.Currency{
		cryptoEntity.Currency_CURRENCY_BTC,
		cryptoEntity.Currency_CURRENCY_TON,
		cryptoEntity.Currency_CURRENCY_USDT,
	}, currency) {
		return Currency{}, fmt.Errorf("unsupported crypto currency: %v", currency)
	}
	return Currency{Crypto: &currency}, nil
}

func (c *Currency) String() string {
	if c.Fiat != nil {
		return c.Fiat.String()
	}
	if c.Crypto != nil {
		return c.Crypto.String()
	}
	return "UNKNOWN"
}

func (c *Currency) IsFiat() bool {
	return c.Fiat != nil
}

func (c *Currency) IsCrypto() bool {
	return c.Crypto != nil
}

func (c *Currency) Validate() error {
	if (c.Fiat == nil && c.Crypto == nil) || (c.Fiat != nil && c.Crypto != nil) {
		return fmt.Errorf("invalid currency: must set exactly one of Fiat or Crypto")
	}
	return nil
}
