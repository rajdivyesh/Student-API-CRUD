/*
		RESTful API: Representational State Transfer Application Programming Interface
	this is a type of Web API which follows a set of architectural principles and constraints.
	Such as:
	1. Stateless:
	2. Resource Based
	3. HTTP Methods
	4. Uniform Inference
	5. Client-Server
*/

/*
This is a simple RESTful API and we use "net/http" package.
In the below ex, we'll create an API that handles basic CRUD operation.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Item struct {
	ID   int    `jason:"id"`
	Name string `jason:"name"`
}

var (
	items      []Item
	itemsMutex sync.Mutex
)

func main11() {
	items = []Item{
		{ID: 1, Name: "C++"},
		{ID: 2, Name: "Pyhton"},
		{ID: 3, Name: "Java"},
	}
	http.HandleFunc("/items", listItems)
	http.HandleFunc("/items/add", addItems)
	http.HandleFunc("/items/update", updateItems)
	http.HandleFunc("/items/delete", deleteItems)

	fmt.Println("Server is running on the port 6060")
	log.Fatal(http.ListenAndServe(":6060", nil))
}

func listItems(w http.ResponseWriter, r *http.Request) {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	// now we will add items into JSON format
	response, err := json.Marshal(items)
	if err != nil {
		http.Error(w, "Failed to Marsoel JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func addItems(w http.ResponseWriter, r *http.Request) {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()
	var newItem Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newItem); err != nil {
		http.Error(w, "Failed to decode request body!", http.StatusBadRequest)
		return
	}
	// now we generate a unique Id for the Item
	newItem.ID = len(items) + 1
	// now we add the new item into the data store for future use
	items = append(items, newItem)
	// Let's respond with the newly creates item
	response, err := json.Marshal(newItem)
	if err != nil {
		http.Error(w, "Failed to Marshal JSON response!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func updateItems(w http.ResponseWriter, r *http.Request) {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()
	var updatedItem Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedItem); err != nil {
		http.Error(w, "Failed to decode request body!", http.StatusBadRequest)
		return
	}
	// now we are going to search for the item in the data store by ID
	var found bool
	for i, item := range items {
		if item.ID == updatedItem.ID {
			// now it's the time to update the item in the data store
			items[i] = updatedItem
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "Item not Found!", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteItems(w http.ResponseWriter, r *http.Request) {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	// Now we parse the request body to get the Id of the item to delete it
	var itemToDelete struct {
		ID int `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&itemToDelete); err != nil {
		http.Error(w, "Failed to decode request body!", http.StatusBadRequest)
		return
	}
	var found bool
	for i, item := range items {
		if item.ID == itemToDelete.ID {
			// now we can remove the item from the data store
			items = append(items[:i], items[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "Item not Found!", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

/*
	This code is a basic example of a RESTful API implemented in Go.
	The API manages a list of items, and it can perform Create, Read, Update, and Delete (CRUD) operations.
	It listens for HTTP requests on port 8080 and has four endpoints:
	`/items` to list all items, `/items/add` to add a new item, `/items/update` to update an item, and `/items/delete` to delete an item.
	The items are represented by a struct called `Item`, which has an ID and a Name.
	To ensure safe access to the list of items, a mutex called `itemsMutex` is used for synchronization.
	The code demonstrates how to handle HTTP requests, decode JSON data, and respond with JSON-formatted data.
	It provides a foundation for building more complex RESTful APIs in Go.

*/
