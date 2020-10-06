package checkout

import (
	// "fmt"
	"github.com/gato/lana/merchandise"
	"github.com/gin-gonic/gin"
	"net/http"
)

func abort(c *gin.Context, err error) {
	status := http.StatusInternalServerError
	// TODO canonalize errors
	if err.Error() == "Basket not found" {
		status = http.StatusNotFound
	}
	http.Error(c.Writer, err.Error(), status)
}

// HandleGetByID - http handler for getting a Basket by Id
func HandleGetByID(c *gin.Context, id string) {
	b, err := GetBasket(id)
	if err != nil {
		abort(c, err)
		return
	}
	// TODO handle error
	_items, _ := b.GetItems()
	amount, _ := b.GetTotal()

	c.JSON(http.StatusOK, gin.H{
		"id":     b.GetID(),
		"items":  _items,
		"amount": amount,
	})
}

// HandleCreateEmtpyBasket - http handler for creating a new basket
func HandleCreateEmtpyBasket(c *gin.Context) {
	b := NewBasket()
	id := b.GetID()
	// TODO: build using url tools
	location := c.Request.Host + c.Request.RequestURI + id
	c.Header("Location", location)
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// HandleDeleteBasket - http handler to delete a basket
func HandleDeleteBasket(c *gin.Context, id string) {
	err := DeleteBasket(id)
	if err != nil {
		abort(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// HandleGetAllBaskets - return all baskets in server
// no pagination so use with caution!
func HandleGetAllBaskets(c *gin.Context) {
	baskets := ListBaskets()
	ids := make([]string, len(baskets))
	for i, v := range baskets {
		ids[i] = v.GetID()
	}
	c.JSON(http.StatusOK, ids)
}

// HandleAddProduct - http handler to delete a basket
func HandleAddProduct(c *gin.Context, id string, _item ProductItem) {
	// Validate product
	if !merchandise.IsValidProduct(_item.Product) {
		http.Error(c.Writer, "Invalid product", http.StatusBadRequest)
		return
	}
	b, err := GetBasket(id)
	if err != nil {
		abort(c, err)
		return
	}
	count, err := b.AddItem(_item)
	if err != nil {
		// This error is extremely weird and can't be easily tested
		// basically the only way AddItem will fail is if basket do not exists
		// but that is also checked in GetBasket
		// only way this could happen is if someone deletes the basket between
		// getting the basket and addItem
		abort(c, err)
		return
	}
	status := http.StatusCreated
	if count != _item.Count {
		// item was already present return OK insted of created
		status = http.StatusOK
	}
	c.JSON(status, gin.H{"count": count})
}
