package checkout

import (
	"fmt"
	"github.com/gato/lana/merchandise"
	"testing"
)

func TestNewBasket(t *testing.T) {
	basketMap = make(map[string]basket)
	basket := NewBasket()
	if basket.GetID() == "" {
		t.Errorf("Basket ID should not be empty")
		return
	}
	items, err := basket.GetItems()
	if err != nil {
		t.Errorf("GetCount returned an error %s", err.Error())
		return
	}
	if len(items) != 0 {
		t.Errorf("Initial Get Items should have no items")
		return
	}
	if len(basketMap) != 1 {
		t.Errorf("Basket was not added to map")
		return
	}
	basket2 := NewBasket()
	if len(basketMap) != 2 {
		t.Errorf("Basket2 was not added to map")
		return
	}
	if basket2.GetID() == basket.GetID() {
		t.Errorf("Basket2 should not be equal to Basket")
	}
}

func TestGetBasket(t *testing.T) {
	basketMap = make(map[string]basket)
	basket := NewBasket()
	basket2, err := GetBasket(basket.GetID())
	if err != nil {
		t.Errorf("GetBasket returned an error %s", err.Error())
		return
	}
	if basket2.GetID() != basket.GetID() {
		t.Errorf("Basket2 should be equal to Basket")
	}
}

func TestGetBasketNotFound(t *testing.T) {
	basketMap = make(map[string]basket)
	b := NewBasket()
	basketMap = make(map[string]basket)
	_, err := GetBasket(b.GetID())
	if err == nil {
		t.Errorf("GetBasket should have returned an error")
		return
	}
	expected := "Basket not found"
	if err.Error() != expected {
		t.Errorf("Wrong error expected %s but got %s", expected, err.Error())
	}
}

func TestAddItem(t *testing.T) {
	basketMap = make(map[string]basket)
	b := NewBasket()
	count, err := b.AddItem(ProductItem{Product: merchandise.PEN, Count: 2})
	if err != nil {
		t.Errorf("AddItem returned an error %s", err.Error())
		return
	}
	if count != 2 {
		t.Errorf("wrong number of items expected 2 got %d", count)
		return
	}
	count, err = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 3})
	if err != nil {
		t.Errorf("AddItem returned an error %s", err.Error())
		return
	}
	if count != 5 {
		t.Errorf("wrong number of items expected 5 got %d", count)
		return
	}
}

func TestMixedAddItem(t *testing.T) {
	basketMap = make(map[string]basket)
	b := NewBasket()
	count, err := b.AddItem(ProductItem{Product: merchandise.PEN, Count: 2})
	if err != nil {
		t.Errorf("AddItem returned an error %s", err.Error())
		return
	}
	if count != 2 {
		t.Errorf("wrong number of items expected 2 got %d", count)
		return
	}
	count, err = b.AddItem(ProductItem{Product: merchandise.TSHIRT, Count: 3})
	if err != nil {
		t.Errorf("AddItem returned an error %s", err.Error())
		return
	}
	if count != 3 {
		t.Errorf("wrong number of items expected 3 got %d", count)
		return
	}
}

func TestAddItemError(t *testing.T) {
	basketMap = make(map[string]basket)
	b := NewBasket()
	count, err := b.AddItem(ProductItem{Product: merchandise.PEN, Count: 2})
	if err != nil {
		t.Errorf("AddItem returned an error %s", err.Error())
		return
	}
	if count != 2 {
		t.Errorf("wrong number of items expected 2 got %d", count)
		return
	}
	// Clear merchandise map to force an error
	basketMap = make(map[string]basket)
	_, err = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 3})
	if err == nil {
		t.Errorf("An error was expected")
		return
	}
}

func findItem(items []ProductItem, code string) (item ProductItem, err error) {
	for _, i := range items {
		if i.Product == code {
			item = i
			return
		}
	}
	return
}

func TestGetItems(t *testing.T) {
	basketMap = make(map[string]basket)
	b := NewBasket()
	_, _ = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 2})
	items, err := b.GetItems()
	if err != nil {
		t.Errorf("GetItems returned an error %s", err.Error())
		return
	}
	if len(items) != 1 {
		t.Errorf("wrong number of items expected 1 got %d", len(items))
		return
	}
	_item, err := findItem(items, merchandise.PEN)
	if err != nil {
		t.Errorf("findItem returned an error %s", err.Error())
		return
	}
	if _item.Count != 2 {
		t.Errorf("wrong number of items expected 2 got %d", _item.Count)
		return
	}

	_, _ = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 3})
	items, err = b.GetItems()
	if len(items) != 1 {
		t.Errorf("wrong number of items expected 1 got %d", len(items))
		return
	}
	_item, err = findItem(items, merchandise.PEN)
	if err != nil {
		t.Errorf("findItem returned an error %s", err.Error())
		return
	}
	if _item.Count != 5 {
		t.Errorf("wrong number of items expected 5 got %d", _item.Count)
		return
	}
}

func TestGetItemsError(t *testing.T) {
	basketMap = make(map[string]basket)
	b := NewBasket()
	_, _ = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 2})
	items, err := b.GetItems()
	if err != nil {
		t.Errorf("GetCount returned an error %s", err.Error())
		return
	}
	if len(items) != 1 {
		t.Errorf("wrong number of items expected 1 got %d", len(items))
		return
	}
	// Clear merchandise map to force an error
	basketMap = make(map[string]basket)
	_, err = b.GetItems()
	if err == nil {
		t.Errorf("An error was expected")
		return
	}
}

func findBasket(baskets []Basket, id string) (Basket, error) {
	for _, b := range baskets {
		if b.GetID() == id {
			return b, nil
		}
	}
	return nil, fmt.Errorf("Basket not found")
}

func TestListBaskets(t *testing.T) {
	basketMap = make(map[string]basket)
	baskets := ListBaskets()
	if len(baskets) != 0 {
		t.Errorf("wrong number of baskets expected 0 got %d", len(baskets))
		return
	}
	b1 := NewBasket()
	_, _ = b1.AddItem(ProductItem{Product: merchandise.PEN, Count: 2})
	b2 := NewBasket()
	_, _ = b2.AddItem(ProductItem{Product: merchandise.TSHIRT, Count: 3})

	baskets = ListBaskets()
	if len(baskets) != 2 {
		t.Errorf("wrong number of baskets expected 2 got %d", len(baskets))
		return
	}
	// Validate items returned are the ones that were added
	_, err := findBasket(baskets, b1.GetID())
	if err != nil {
		t.Errorf("first basket not found")
		return
	}
	_, err = findBasket(baskets, b2.GetID())
	if err != nil {
		t.Errorf("second basket not found")
		return
	}
}

func TestGetTotal(t *testing.T) {
	// Items: PEN, TSHIRT, MUG
	// Total: 32.50€
	b := NewBasket()
	_, _ = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.TSHIRT, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.MUG, Count: 1})
	total, err := b.GetTotal()
	if err != nil {
		t.Errorf("GetTotal returned an error %s", err.Error())
		return
	}
	expected := 32.5
	if total != expected {
		t.Errorf("invalid total expected %.2f got %.2f", expected, total)
		return
	}
	// Items: PEN, TSHIRT, PEN
	// Total: 25.00€
	b = NewBasket()
	_, _ = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.TSHIRT, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 1})
	total, _ = b.GetTotal()
	expected = 25
	if total != expected {
		t.Errorf("invalid total expected %.2f got %.2f", expected, total)
		return
	}

	// Items: TSHIRT, TSHIRT, TSHIRT, PEN, TSHIRT
	// Total: 65.00€
	b = NewBasket()
	_, _ = b.AddItem(ProductItem{Product: merchandise.TSHIRT, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.TSHIRT, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.TSHIRT, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.TSHIRT, Count: 1})
	total, _ = b.GetTotal()
	expected = 65
	if total != expected {
		t.Errorf("invalid total expected %.2f got %.2f", expected, total)
		return
	}

	// Items: PEN, TSHIRT, PEN, PEN, MUG, TSHIRT, TSHIRT
	// Total: 62.50€
	b = NewBasket()
	_, _ = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.TSHIRT, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.MUG, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.TSHIRT, Count: 1})
	_, _ = b.AddItem(ProductItem{Product: merchandise.TSHIRT, Count: 1})
	total, _ = b.GetTotal()
	expected = 62.50
	if total != expected {
		t.Errorf("invalid total expected %.2f got %.2f", expected, total)
		return
	}
}

func TestGetTotalBasketNotFoundError(t *testing.T) {
	b := NewBasket()
	_, _ = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 1})
	// force error
	basketMap = make(map[string]basket)
	_, err := b.GetTotal()
	if err == nil {
		t.Errorf("GetBasket should have returned an error")
		return
	}
	expected := "Basket not found"
	if err.Error() != expected {
		t.Errorf("Wrong error expected %s but got %s", expected, err.Error())
	}
}

type FailingPromo struct{}

func (promotion FailingPromo) Apply(Items map[string]item) (discounts []Discount, err error) {
	err = fmt.Errorf("some random error")
	return
}

func TestGetTotalPromotionError(t *testing.T) {
	basketMap = make(map[string]basket)
	b := NewBasket()
	// Here I'm adding a failing promotion into basket internals
	// This can't be done from the outside as users will only see
	// the public interfase
	// Only for the sake of test coverage. no real value here
	// get the internal basket representation (in a not thread safe way)
	internalBasket, _ := basketMap[b.GetID()]
	// add new failing promotion
	internalBasket.promotions = append(basketMap[b.GetID()].promotions, FailingPromo{})
	// replace basket
	basketMap[b.GetID()] = internalBasket
	_, _ = b.AddItem(ProductItem{Product: merchandise.PEN, Count: 1})
	_, err := b.GetTotal()
	if err == nil {
		t.Errorf("GetBasket should have returned an error")
		return
	}
	expected := "some random error"
	if err.Error() != expected {
		t.Errorf("Wrong error expected %s but got %s", expected, err.Error())
	}
}

func TestDeleteBasket(t *testing.T) {
	basketMap = make(map[string]basket)
	basket := NewBasket()
	baskets := ListBaskets()
	if len(baskets) != 1 {
		t.Errorf("wrong number of baskets expected 1 got %d", len(baskets))
		return
	}
	err := DeleteBasket(basket.GetID())
	if err != nil {
		t.Errorf("DeleteBasket returned an error %s", err.Error())
		return
	}
	baskets = ListBaskets()
	if len(baskets) != 0 {
		t.Errorf("wrong number of baskets expected 0 got %d", len(baskets))
		return
	}
	_, err = GetBasket(basket.GetID())
	if err == nil {
		t.Errorf("GetBasket should have returned an error")
		return
	}
	expected := "Basket not found"
	if err.Error() != expected {
		t.Errorf("Wrong error expected %s but got %s", expected, err.Error())
	}
}

func TestDeleteBasketNotFoundError(t *testing.T) {
	basketMap = make(map[string]basket)
	err := DeleteBasket("1234")
	if err == nil {
		t.Errorf("DeleteBasket should have returned an error")
		return
	}
	expected := "Basket not found"
	if err.Error() != expected {
		t.Errorf("Wrong error expected %s but got %s", expected, err.Error())
	}
}
