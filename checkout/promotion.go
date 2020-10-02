package checkout

import (
	"fmt"
	"github.com/gato/lana/merchandise"
)

// Discount - model, one discount line
type Discount struct {
	Description string
	Amount      float64
}

// Promotion - interface that apply to items and generates Discounts
type Promotion interface {
	Apply(Items map[string]Item) ([]Discount, error)
}

// BuyXGetY - buy 2 get 1 free promotion type (BuyQuantity must be greater than GetFreeQuantity)
type BuyXGetY struct {
	BuyQuantity     int64
	GetFreeQuantity int64
	Code            string
}

// BulkPercentageDiscount - Buy 3 or more to get a 25% on price per unit promotion type
type BulkPercentageDiscount struct {
	BuyQuantity        int64
	DiscountPercentage int64
	Code               string
}

// Apply - Buy X get Y
func (promotion BuyXGetY) Apply(Items map[string]Item) (discounts []Discount, err error) {
	item, ok := Items[promotion.Code]
	if !ok || item.Count < promotion.BuyQuantity {
		// No items of type CODE or not enough of them
		return
	}
	// Discount Y items every X items bought
	m := item.Count / promotion.BuyQuantity
	d := item.Product.Price * float64(promotion.GetFreeQuantity*m)
	discounts = append(discounts, Discount{
		Description: fmt.Sprintf("Buy %d %s and get %d Free", promotion.BuyQuantity, merchandise.GetProduct(promotion.Code).Name, promotion.GetFreeQuantity),
		Amount:      d,
	})
	return
}

// Apply - Buy x or more get y% off
func (promotion BulkPercentageDiscount) Apply(Items map[string]Item) (discounts []Discount, err error) {
	item, ok := Items[promotion.Code]
	if !ok || item.Count < promotion.BuyQuantity {
		// No items of type CODE or not enough of them
		return
	}
	p := float64(promotion.DiscountPercentage) / 100
	d := item.Product.Price * p * float64(item.Count)
	discounts = append(discounts, Discount{
		Description: fmt.Sprintf("Buy %d or more %s get %d%% off", promotion.BuyQuantity, merchandise.GetProduct(promotion.Code).Name, promotion.DiscountPercentage),
		Amount:      d,
	})
	return
}

// PenBuy2Get1 - Buy 2 Pens get 1 Free Promotion
var PenBuy2Get1 = BuyXGetY{Code: merchandise.PEN, BuyQuantity: 2, GetFreeQuantity: 1}

// TshirtBuy3Get25OFF - Buy 3 or more shirts get 25% off
var TshirtBuy3Get25OFF = BulkPercentageDiscount{Code: merchandise.TSHIRT, BuyQuantity: 3, DiscountPercentage: 25}
