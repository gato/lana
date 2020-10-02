package merchandise

import "testing"

func TestProductMapIsInitialized(t *testing.T) {
	if len(products) != 3 {
		t.Errorf("Products Map is not properly initialized")
	}
	if products[PEN].Code != PEN {
		t.Errorf("Lana Pen not found on map")
	}
	if products[TSHIRT].Code != TSHIRT {
		t.Errorf("Lana T-Shirt not found on map")
	}
	if products[MUG].Code != MUG {
		t.Errorf("Lana Coffee Mug not found on map")
	}
}

func TestGetProducts(t *testing.T) {
	if GetProduct(PEN) != products[PEN] {
		t.Errorf("Lana Pen not found!")
	}
	if GetProduct(TSHIRT) != products[TSHIRT] {
		t.Errorf("Lana T-Shirt not found ")
	}
	if GetProduct(MUG) != products[MUG] {
		t.Errorf("Lana Coffee Mug not found")
	}
}
