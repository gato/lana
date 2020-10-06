package checkout

import (
	"fmt"
	"github.com/gato/lana/merchandise"
	"github.com/google/uuid"
	"sync"
)

type item struct {
	Product merchandise.Product
	Count   int64
}

// ProductItem - DTO for basket entries
type ProductItem struct {
	Product string `json:"product"`
	Count   int64  `json:"count"`
}

type basket struct {
	id         string
	items      map[string]item
	promotions []Promotion
	lock       *sync.RWMutex
}

func (basket *basket) getItems() []ProductItem {
	basket.lock.RLock()
	defer basket.lock.RUnlock()
	items := make([]ProductItem, len(basket.items))
	i := 0
	for _, item := range basket.items {
		items[i] = ProductItem{Product: item.Product.Code, Count: item.Count}
		i++
	}
	return items
}

// Mutex to syncronize accest to fake datastore (a simple map)
var basketLock = sync.RWMutex{}

// poor man's datatstore
var basketMap = make(map[string]basket)

// Basket - interface to access minimum needed basket functionanlity without exporting
// internal implementation
type Basket interface {
	GetID() string
	GetItems() ([]ProductItem, error)
	AddItem(ProductItem) (int64, error)
	GetTotal() (float64, error)
}

func createBasket() (basket basket) {
	uuid := uuid.Must(uuid.NewRandom())

	basket.id = uuid.String()
	basket.items = make(map[string]item)
	basket.promotions = make([]Promotion, 2)
	basket.lock = &sync.RWMutex{}
	basket.promotions[0] = PenBuy2Get1
	basket.promotions[1] = TshirtBuy3Get25OFF
	return
}

func getBasket(id string) (basket basket, ok bool) {
	basketLock.RLock()
	defer basketLock.RUnlock()
	basket, ok = basketMap[id]
	return
}

// BasketWrapper - Implements Basket Interface
type BasketWrapper struct {
	id string
}

// GetID - Get Basket identifier for future reference
func (b BasketWrapper) GetID() string {
	return b.id
}

// GetItems - Get Basket's item count
func (b BasketWrapper) GetItems() ([]ProductItem, error) {
	basket, ok := getBasket(b.id)
	if !ok {
		return nil, fmt.Errorf("Basket not found")
	}
	return basket.getItems(), nil
}

// AddItem - add "amount" items to basket
// if product exist it will add the amount
// if not will set
func (b BasketWrapper) AddItem(_item ProductItem) (int64, error) {
	basket, ok := getBasket(b.id)
	if !ok {
		return 0, fmt.Errorf("Basket not found")
	}
	// ADD item
	basket.lock.Lock()
	defer basket.lock.Unlock()
	i, ok := basket.items[_item.Product]
	if !ok {
		i = item{Product: merchandise.GetProduct(_item.Product), Count: 0}
	}
	i.Count = i.Count + _item.Count
	basket.items[_item.Product] = i

	return i.Count, nil
}

// GetTotal - calculate amount to be paid for the basket
func (b BasketWrapper) GetTotal() (float64, error) {
	// TODO GET basket
	basket, ok := getBasket(b.id)
	if !ok {
		return 0, fmt.Errorf("Basket not found")
	}
	var total float64 = 0
	basket.lock.RLock()
	defer basket.lock.RUnlock()
	// sumarize products
	for _, item := range basket.items {
		total += (item.Product.Price * float64(item.Count))
	}
	// calculate discounts
	for _, promo := range basket.promotions {
		discounts, err := promo.Apply(basket.items)
		if err != nil {
			return 0, err
		}
		for _, discount := range discounts {
			total -= discount.Amount
		}
	}
	// RETURN total
	return total, nil
}

// NewBasket - creates a new basket and returns a BasketWrapper to it
func NewBasket() Basket {
	basket := createBasket()
	// Get write lock (no need to check for existance as we asume uuids are unique)
	basketLock.Lock()
	defer basketLock.Unlock()
	// ADD to map
	basketMap[basket.id] = basket
	return BasketWrapper{id: basket.id}
}

// GetBasket - Get basket by id
func GetBasket(id string) (Basket, error) {
	basket, ok := getBasket(id)
	if !ok {
		return nil, fmt.Errorf("Basket not found")
	}
	return BasketWrapper{id: basket.id}, nil
}

// ListBaskets - Get Baskets ids with item count
func ListBaskets() []Basket {
	basketLock.RLock()
	defer basketLock.RUnlock()
	list := make([]Basket, 0)
	for _, basket := range basketMap {
		list = append(list, BasketWrapper{id: basket.id})
	}
	return list
}

// DeleteBasket - Remove a Basket from storage
func DeleteBasket(id string) error {
	basketLock.RLock()
	defer basketLock.RUnlock()
	_, ok := basketMap[id]
	if !ok {
		return fmt.Errorf("Basket not found")
	}
	delete(basketMap, id)
	return nil
}
