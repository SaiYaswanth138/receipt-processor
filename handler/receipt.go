package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"receipt-processor/model"
	"receipt-processor/service"
	"receipt-processor/storage"
	"github.com/google/uuid"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var receipt model.Receipt
	_ = json.Unmarshal(body, &receipt)

	id := uuid.New().String()
	points := service.CalculatePoints(receipt)

	storage.Receipts[id] = points

	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/receipts/"):len("/receipts/")+36]
	points, ok := storage.Receipts[id]
	if !ok {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
