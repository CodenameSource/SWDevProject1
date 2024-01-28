package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	"webapp/db"
	kafkamessenger "webapp/kafkaMessenger"
)

var mutex sync.Mutex

func Init(dsn string) (*mux.Router, error) {
	router := mux.NewRouter()

	err := db.Init(dsn)

	router.HandleFunc("/api/addItem", addItemHandler).Methods("POST")
	router.HandleFunc("/api/getItems", getItemsHandler).Methods("GET")
	router.HandleFunc("/api/removeItem", removeItemHandler).Methods("DELETE")
	router.HandleFunc("/api/updatePrice", updatePriceHandler).Methods("GET")
	router.HandleFunc("/api/updatePrices", updatePricesHandler).Methods("GET")

	return router, err
}

func addItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem db.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = db.CreateItem(newItem.URL, 0)

	if err != nil {
		http.Error(w, "Error adding item", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getItemsHandler(w http.ResponseWriter, r *http.Request) {
	productsList, err := db.GetItems()

	if err != nil {
		http.Error(w, "Error getting items", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productsList)
}

func updatePriceHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}
	mutex.Lock()

	product, err := db.GetItem(url)

	if err != nil {
		http.Error(w, "Error refreshing item", http.StatusBadRequest)
		return
	}

	kafkamessenger.SendRefreshEvent(product.URL)

	defer mutex.Unlock()

	w.WriteHeader(http.StatusOK)
}

func updatePricesHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()

	products, err := db.GetItems()

	if err != nil {
		http.Error(w, "Error refreshing items", http.StatusBadRequest)
		return
	}

	var product db.Item

	fmt.Println("Sending refresh events")

	for _, product = range products {
		fmt.Println("Waiting for refresh event")
		fmt.Println("Now refreshing url: " + product.URL)
		err = kafkamessenger.SendRefreshEvent(product.URL)
		if err != nil {
			http.Error(w, "Error refreshing items", http.StatusBadRequest)
			return
		}
		url, price, err := kafkamessenger.ListenForRefreshPrice()
		fmt.Println("Refreshed url: " + url + " with price: " + fmt.Sprintf("%f", price))
		if err != nil {
			http.Error(w, "Error refreshing items", http.StatusBadRequest)
			return
		}
		db.EditItem(url, price)
	}

	defer mutex.Unlock()

	fmt.Println("Successfully refreshed all items")

	w.WriteHeader(http.StatusOK)
}

func removeItemHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}

	mutex.Lock()

	fmt.Println("Removing item with url: " + url)

	err := db.DeleteItem(url)

	if err != nil {
		http.Error(w, "Error deleting item", http.StatusBadRequest)
		return
	}
	mutex.Unlock()

	w.WriteHeader(http.StatusOK)
}
