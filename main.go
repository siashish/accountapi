package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"accountapi/models"

	"github.com/gorilla/mux"
	"github.com/siashish/go/helper"
)

func create(w http.ResponseWriter, r *http.Request) {

	// Set header
	w.Header().Set("Content-Type", "application/json")

	var account models.AccountData

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&account)

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), account)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func fetch(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	var account models.AccountData
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	account_id, _ := primitive.ObjectIDFromHex(params["account_id"])

	filter := bson.M{"ID": account_id}
	err := collection.FindOne(context.TODO(), filter).Decode(&account)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func delete(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	account_id, err := primitive.ObjectIDFromHex(params["account_id"])

	// prepare filter.
	filter := bson.M{"ID": account_id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

func main() {
	// Init Router
	r := mux.Router()

	r.HandleFunc("/v1/organisation/accounts", create()).Method("POST")
	r.HandleFunc("/v1/organisation/accounts/{account_id}", fetch()).Method("GET")
	r.HandleFunc("/v1/organisation/accounts/{account_id}?version={version}", delete()).Method("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
