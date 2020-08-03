package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupAPI()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Pong", w.Body.String())

	router.ServeHTTP(w, req)

}
func TestOrdersFunc(t *testing.T){
	router := setupAPI()
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET","/api/orders",nil)

	if err!=nil{
		panic(err)
	}

	router.ServeHTTP(w, req)
}

func TestPostOrder(t *testing.T) {
	router := setupAPI()
	w := httptest.NewRecorder()

	order := &Order{
		OrderId: 101,
		CustomerName: "rt",
		OrderReview: "good",
	}
	jsonOrder, _ := json.Marshal(order)

	req, err := http.NewRequest("POST","/api/PostOrder",bytes.NewBuffer(jsonOrder))

	response := httptest.NewRecorder()
	if err != nil{
		panic(err)
	}
	assert.Equal(t, 200, response.Code, "OK response is expected")

	router.ServeHTTP(w, req)
}