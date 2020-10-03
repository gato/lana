package checkout

import (
	"github.com/gato/lana/merchandise"
	"testing"
)

func TestBuy2PenGet1FreePromotion(t *testing.T) {
	items := make(map[string]item)
	items[merchandise.PEN] = item{Product: merchandise.GetProduct(merchandise.PEN), Count: 2}
	discount, err := PenBuy2Get1.Apply(items)
	if err != nil {
		t.Errorf("There was an error calculating PEN discounts")
		return
	}
	if len(discount) != 1 {
		t.Errorf("PEN Discount was not applied")
		return
	}
	expectedAmount := merchandise.GetProduct(merchandise.PEN).Price
	if discount[0].Amount != expectedAmount {
		t.Errorf("Discount should be %.2f but was %.2f", expectedAmount, discount[0].Amount)
	}
	expected := "Buy 2 Lana Pen and get 1 Free"
	if discount[0].Description != expected {
		t.Errorf("Discount description should be %s but was %s", expected, discount[0].Description)
	}
}
func TestBuy2PenGet1FreePromotionMulti(t *testing.T) {
	items := make(map[string]item)
	items[merchandise.PEN] = item{Product: merchandise.GetProduct(merchandise.PEN), Count: 20}
	discount, err := PenBuy2Get1.Apply(items)
	if err != nil {
		t.Errorf("There was an error calculating PEN discounts")
		return
	}
	if len(discount) != 1 {
		t.Errorf("PEN Discount was not applied")
		return
	}
	expectedAmount := merchandise.GetProduct(merchandise.PEN).Price * float64(10)
	if discount[0].Amount != expectedAmount {
		t.Errorf("Discount should be %.2f but was %.2f", expectedAmount, discount[0].Amount)
	}
	expected := "Buy 2 Lana Pen and get 1 Free"
	if discount[0].Description != expected {
		t.Errorf("Discount description should be %s but was %s", expected, discount[0].Description)
	}
}
func TestBuy2PenGet1FreeNotApplyIfNoPens(t *testing.T) {
	items := make(map[string]item)
	items[merchandise.TSHIRT] = item{Product: merchandise.GetProduct(merchandise.TSHIRT), Count: 2}
	discount, err := PenBuy2Get1.Apply(items)
	if err != nil {
		t.Errorf("There was an error calculating PEN discounts")
		return
	}
	if len(discount) != 0 {
		t.Errorf("PEN Discount was applied but there are no pens in the basket")
		return
	}
}

func TestBuy2PenGet1FreeNotApplyIfLessPenThanNeeded(t *testing.T) {
	items := make(map[string]item)
	items[merchandise.PEN] = item{Product: merchandise.GetProduct(merchandise.PEN), Count: 1}
	discount, err := PenBuy2Get1.Apply(items)
	if err != nil {
		t.Errorf("There was an error calculating PEN discounts")
		return
	}
	if len(discount) != 0 {
		t.Errorf("PEN Discount was applied but there are no enough pens in the basket")
		return
	}
}

func TestBuyXGetYCalculations(t *testing.T) {
	buyXGetY := BuyXGetY{Code: merchandise.PEN, BuyQuantity: 3, GetFreeQuantity: 2}
	items := make(map[string]item)
	items[merchandise.PEN] = item{Product: merchandise.GetProduct(merchandise.PEN), Count: 3}
	discount, err := buyXGetY.Apply(items)
	if err != nil {
		t.Errorf("There was an error calculating Custom discounts")
		return
	}
	if len(discount) != 1 {
		t.Errorf("Custom Discount was not applied")
		return
	}
	expectedAmount := merchandise.GetProduct(merchandise.PEN).Price * float64(2)
	if discount[0].Amount != expectedAmount {
		t.Errorf("Custom Discount should be %.2f but was %.2f", expectedAmount, discount[0].Amount)
	}

	expected := "Buy 3 Lana Pen and get 2 Free"
	if discount[0].Description != expected {
		t.Errorf("Discount description should be %s but was %s", expected, discount[0].Description)
	}
}

func TestBuy3TshirtsGet25OffPromotion(t *testing.T) {
	items := make(map[string]item)
	items[merchandise.TSHIRT] = item{Product: merchandise.GetProduct(merchandise.TSHIRT), Count: 3}
	discount, err := TshirtBuy3Get25OFF.Apply(items)
	if err != nil {
		t.Errorf("There was an error calculating Tshirt discounts")
		return
	}
	if len(discount) != 1 {
		t.Errorf("Tshirt Discount was not applied")
		return
	}
	expectedAmount := merchandise.GetProduct(merchandise.TSHIRT).Price * float64(3) * .25
	if discount[0].Amount != expectedAmount {
		t.Errorf("Discount should be %.2f but was %.2f", expectedAmount, discount[0].Amount)
	}
	expected := "Buy 3 or more Lana T-Shirt get 25% off"
	if discount[0].Description != expected {
		t.Errorf("Discount description should be %s but was %s", expected, discount[0].Description)
	}
}

func TestBuy3TshirtsGet25OffPromotionMulti(t *testing.T) {
	items := make(map[string]item)
	items[merchandise.TSHIRT] = item{Product: merchandise.GetProduct(merchandise.TSHIRT), Count: 8}
	discount, err := TshirtBuy3Get25OFF.Apply(items)
	if err != nil {
		t.Errorf("There was an error calculating Tshirt discounts")
		return
	}
	if len(discount) != 1 {
		t.Errorf("Tshirt Discount was not applied")
		return
	}
	expectedAmount := merchandise.GetProduct(merchandise.TSHIRT).Price * float64(8) * .25
	if discount[0].Amount != expectedAmount {
		t.Errorf("Discount should be %.2f but was %.2f", expectedAmount, discount[0].Amount)
	}
	expected := "Buy 3 or more Lana T-Shirt get 25% off"
	if discount[0].Description != expected {
		t.Errorf("Discount description should be %s but was %s", expected, discount[0].Description)
	}
}

func TestBuy3TshirtsGet25OffNotApplyIfNoTshirt(t *testing.T) {
	items := make(map[string]item)
	items[merchandise.PEN] = item{Product: merchandise.GetProduct(merchandise.PEN), Count: 3}
	discount, err := TshirtBuy3Get25OFF.Apply(items)
	if err != nil {
		t.Errorf("There was an error calculating Tshirt discounts")
		return
	}
	if len(discount) != 0 {
		t.Errorf("Tshirt Discount was applied but there are no Tshirt in the basket")
		return
	}
}

func TestBuy3TshirtsGet25OffNotApplyIfLessTshirt(t *testing.T) {
	items := make(map[string]item)
	items[merchandise.TSHIRT] = item{Product: merchandise.GetProduct(merchandise.TSHIRT), Count: 2}
	discount, err := TshirtBuy3Get25OFF.Apply(items)
	if err != nil {
		t.Errorf("There was an error calculating Tshirt discounts")
		return
	}
	if len(discount) != 0 {
		t.Errorf("Tshirt Discount was applied but there are no enough Tshirt in the basket")
		return
	}
}
