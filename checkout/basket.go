package checkout

import (
	"github.com/gato/lana/merchandise"
	"github.com/google/uuid"
)

// Item - model, one entry in the Basket, what Product and how many are in
type Item struct {
	Product merchandise.Product
	Count   int64
}

// Basket - model
type Basket struct {
	ID         string
	Items      map[string]Item
	Promotions []Promotion
}

// CreateBasket - creates an empty basket with default promotions already applied
func CreateBasket() (basket Basket) {
	uuid := uuid.Must(uuid.NewRandom())
	basket.ID = uuid.String()
	basket.Items = make(map[string]Item)
	basket.Promotions = make([]Promotion, 2)
	basket.Promotions[0] = PenBuy2Get1
	basket.Promotions[1] = TshirtBuy3Get25OFF
	return
}

// BasketStore
