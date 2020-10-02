package checkout

import "testing"

func TestCreateBasketWillReturnAnEmptyNewBasket(t *testing.T) {
	basket := CreateBasket()
	if len(basket.Promotions) != 2 {
		t.Errorf("Basket has wrong number of baskets")
		return
	}
	if len(basket.Items) != 0 {
		t.Errorf("Basket is not empty")
		return
	}
	basket2 := CreateBasket()
	if basket.ID == basket2.ID {
		t.Errorf("Basket is not new")
		return
	}
}
