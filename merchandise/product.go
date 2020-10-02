package merchandise

// Product - model, Lana's awesome merchandise item (PEN, TSHIRT, MUG)
// Code         | Name              |  Price
// -----------------------------------------------
// PEN          | Lana Pen          |   5.00€
// TSHIRT       | Lana T-Shirt      |  20.00€
// MUG          | Lana Coffee Mug   |   7.50€
type Product struct {
	Code  string
	Name  string
	Price float64
}

// PEN constant for lookup
const PEN string = "PEN"

// TSHIRT constant for lookup
const TSHIRT string = "TSHIRT"

// MUG constant for lookup
const MUG string = "MUG"

var products = map[string]Product{
	PEN:    Product{Code: PEN, Name: "Lana Pen", Price: 5.00},
	TSHIRT: Product{Code: TSHIRT, Name: "Lana T-Shirt", Price: 20.00},
	MUG:    Product{Code: MUG, Name: "Lana Coffee Mug", Price: 7.50},
}

// GetProduct - dummy function to simulate access to some product "persistance"
func GetProduct(prod string) Product {
	return products[prod]
}
