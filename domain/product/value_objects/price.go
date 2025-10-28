package value_objects

import "fmt"

type Price struct {
	amount   float64
	currency string
}

func NewPrice(amount float64, currency string) Price {
	if amount < 0 {
		panic("amount must be greater than zero")
	}
	return Price{amount: amount, currency: currency}
}

func (p Price) String() string {
	return fmt.Sprintf("%.2f %s", p.amount, p.currency)
}

func (p Price) EqualCurrency(currency string) bool {
	return p.currency == currency
}
