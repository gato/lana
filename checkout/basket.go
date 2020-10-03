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

type basket struct {
	id         string
	items      map[string]item
	promotions []Promotion
	lock       *sync.RWMutex
}

func (basket *basket) getCount() int64 {
	var count int64 = 0
	basket.lock.RLock()
	defer basket.lock.RUnlock()
	for _, item := range basket.items {
		count += item.Count
	}
	return count
}

// Mutex to syncronize accest to fake datastore (a simple map)
var basketLock = sync.RWMutex{}

// poor man's datatstore
var basketMap = make(map[string]basket)

// Basket - interface to access minimum needed basket functionanlity without exporting
// internal implementation
type Basket interface {
	GetID() string
	GetCount() (int64, error)
	AddItem(product string, amount int64) (int64, error)
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

// GetCount - Get Basket's item count
func (b BasketWrapper) GetCount() (int64, error) {
	basket, ok := getBasket(b.id)
	if !ok {
		return 0, fmt.Errorf("Basket not found")
	}
	return basket.getCount(), nil
}

// AddItem - add "amount" items to basket
// if product exist it will add the amount
// if not will set
func (b BasketWrapper) AddItem(product string, amount int64) (int64, error) {
	basket, ok := getBasket(b.id)
	if !ok {
		return 0, fmt.Errorf("Basket not found")
	}
	// ADD item
	basket.lock.Lock()
	defer basket.lock.Unlock()
	i, ok := basket.items[product]
	if !ok {
		i = item{Product: merchandise.GetProduct(product), Count: 0}
	}
	i.Count = i.Count + amount
	basket.items[product] = i

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
func ListBaskets() ([]Basket, error) {
	basketLock.RLock()
	defer basketLock.RUnlock()
	list := make([]Basket, 0)
	for _, basket := range basketMap {
		list = append(list, BasketWrapper{id: basket.id})
	}
	return list, nil
}
