package checkout

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func getRouter() *gin.Engine {
	r := gin.Default()
	apiv1 := r.Group("/api/v1/")
	AddRoutes(apiv1)
	return r
}

func TestHandleGetByID(t *testing.T) {
	basket := NewBasket()
	r := getRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/basket/"+basket.GetID(), nil)
	r.ServeHTTP(w, req)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("HandleGetByID wrong http status expected %d got %d", expected, w.Code)
		return
	}
	expectedBody := fmt.Sprintf("{\"amount\":0,\"id\":\"%s\",\"items\":[]}", basket.GetID())
	if w.Body.String() != expectedBody {
		t.Errorf("HandleGetByID wrong response body expected %s got %s", expectedBody, w.Body.String())
		return
	}
}
func TestHandleGetByIDNotFound(t *testing.T) {
	r := getRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/basket/123", nil)
	r.ServeHTTP(w, req)

	expected := http.StatusNotFound
	if w.Code != expected {
		t.Errorf("HandleGetByID wrong http status expected %d got %d", expected, w.Code)
		return
	}
	expectedBody := "Basket not found\n"
	if w.Body.String() != expectedBody {
		t.Errorf("HandleGetByID wrong response body expected %s got %s", expectedBody, w.Body.String())
		return
	}
}

func TestHandleCreateEmtpyBasket(t *testing.T) {
	r := getRouter()
	basketMap = make(map[string]basket)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/basket/", nil)
	r.ServeHTTP(w, req)

	expected := http.StatusCreated
	if w.Code != expected {
		t.Errorf("HandleCreateEmtpyBasket wrong http status expected %d got %d", expected, w.Code)
		return
	}

	baskets, _ := ListBaskets()
	if len(baskets) != 1 {
		t.Errorf("HandleCreateEmtpyBasket haven't created a basket")
		return
	}

	expectedBody := fmt.Sprintf("{\"id\":\"%s\"}", baskets[0].GetID())
	if w.Body.String() != expectedBody {
		t.Errorf("HandleCreateEmtpyBasket wrong response body expected %s got %s", expectedBody, w.Body.String())
		return
	}
}

func TestHandleDeleteBasket(t *testing.T) {
	basket := NewBasket()
	r := getRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/basket/"+basket.GetID(), nil)
	r.ServeHTTP(w, req)

	expected := http.StatusNoContent
	if w.Code != expected {
		t.Errorf("HandleGetByID wrong http status expected %d got %d", expected, w.Code)
		return
	}
}

func TestHandleDeleteBasketNotFound(t *testing.T) {
	r := getRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/basket/123", nil)
	r.ServeHTTP(w, req)

	expected := http.StatusNotFound
	if w.Code != expected {
		t.Errorf("HandleDeleteBasket wrong http status expected %d got %d", expected, w.Code)
		return
	}
	expectedBody := "Basket not found\n"
	if w.Body.String() != expectedBody {
		t.Errorf("HandleDeleteBasket wrong response body expected %s got %s", expectedBody, w.Body.String())
		return
	}
}

func TestHandleAddProduct(t *testing.T) {
	basket := NewBasket()
	r := getRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/basket/"+basket.GetID(), strings.NewReader("{\"product\":\"MUG\",\"count\":1}"))
	r.ServeHTTP(w, req)

	expected := http.StatusCreated
	if w.Code != expected {
		t.Errorf("HandleAddProduct wrong http status expected %d got %d", expected, w.Code)
		return
	}
	expectedBody := "{\"count\":1}"
	if w.Body.String() != expectedBody {
		t.Errorf("HandleAddProduct wrong response body expected %s got %s", expectedBody, w.Body.String())
		return
	}
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/basket/"+basket.GetID(), strings.NewReader("{\"product\":\"MUG\",\"count\":1}"))
	r.ServeHTTP(w, req)

	expected = http.StatusOK
	if w.Code != expected {
		t.Errorf("HandleAddProduct wrong http status expected %d got %d", expected, w.Code)
		return
	}
	expectedBody = "{\"count\":2}"
	if w.Body.String() != expectedBody {
		t.Errorf("HandleAddProduct wrong response body expected %s got %s", expectedBody, w.Body.String())
		return
	}
}

func TestHandleAddProductErrorInvalidProduct(t *testing.T) {
	basket := NewBasket()
	r := getRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/basket/"+basket.GetID(), strings.NewReader("{\"product\":\"Rocket Fuel\",\"count\":1}"))
	r.ServeHTTP(w, req)

	expected := http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("HandleAddProduct wrong http status expected %d got %d", expected, w.Code)
		return
	}
	expectedBody := "Invalid product\n"
	if w.Body.String() != expectedBody {
		t.Errorf("HandleAddProduct wrong response body expected %s got %s", expectedBody, w.Body.String())
		return
	}
}

func TestHandleAddProductErrorBasketNotFound(t *testing.T) {
	r := getRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/basket/1111", strings.NewReader("{\"product\":\"PEN\",\"count\":1}"))
	r.ServeHTTP(w, req)

	expected := http.StatusNotFound
	if w.Code != expected {
		t.Errorf("HandleAddProduct wrong http status expected %d got %d", expected, w.Code)
		return
	}
	expectedBody := "Basket not found\n"
	if w.Body.String() != expectedBody {
		t.Errorf("HandleAddProduct wrong response body expected %s got %s", expectedBody, w.Body.String())
		return
	}
}

func TestHandleAddProductErrorBadPayload(t *testing.T) {
	basket := NewBasket()
	r := getRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/basket/"+basket.GetID(), strings.NewReader("{\"lala\":\"lolo\""))
	r.ServeHTTP(w, req)

	expected := http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("HandleAddProduct wrong http status expected %d got %d", expected, w.Code)
		return
	}
}
