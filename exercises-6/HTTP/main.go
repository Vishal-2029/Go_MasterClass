package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2/log"
)




type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var item Item
	if err := json.Unmarshal(body,&item); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "Item created successfully",
		"item":    item,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}



var mutex sync.Mutex
var items = []Item{}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	var itemList = append([]Item{}, items...)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemList)
}
 
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var item Item
	if err := json.Unmarshal(body, &item); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for i, existingItem := range items {
		if existingItem.ID == item.ID {
			items[i] = item
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, `{"error": "Item not found"}`, http.StatusNotFound)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var item Item
	if err := json.Unmarshal(body, &item); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for i, existingItem := range items {
		if existingItem.ID == item.ID {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, `{"error": "Item not found"}`, http.StatusNotFound)
}



func main() {

	http.HandleFunc("/create", CreateItem)
	http.HandleFunc("/Allitems", GetAllItems)
	http.HandleFunc("/update", UpdateItem)
	http.HandleFunc("/delete", deleteItem)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
