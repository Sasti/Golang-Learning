package shop

import (
	"math"
)

// BasketItem holds a reference to the Article and the amount for that article.
type BasketItem struct {
	*Article
	amount uint8
}

// CoupongType is relevant for distinguisihing the different couping types.
type CoupongType uint8

// Definition of the different available Voucher types
const (
	PresentVoucher CoupongType = 1
	Coupong        CoupongType = 2
)

// Voucher is a model for the different coupongs used in babymarkt.de
type Voucher struct {
	CoupongType
	amount float32
}

// The Basket struct contains all relevant objects which make up the users basket.
type Basket struct {
	Items    []BasketItem
	Vouchers []Voucher
}

// CalculateBasketTotal calculates the total sum of all products and also takes into account the present Vouchers.
func (b *Basket) CalculateBasketTotal() float32 {
	var total float32

	for _, item := range b.Items {
		total += item.CalculateBrutto() * float32(item.amount)
	}

	for _, voucher := range b.Vouchers {
		switch voucher.CoupongType {
		case PresentVoucher:
			total -= voucher.amount
		case Coupong:
			total -= voucher.amount * (100.0 - 0.19)
		}
	}

	total = float32(roundPlus(float64(total), 2))

	return total
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func roundPlus(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return round(f*shift) / shift
}
