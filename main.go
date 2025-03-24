package main

import (
	"net/http"
	"receipt-processor/handler"
)

func main() {
	http.HandleFunc("/receipts/process", handler.ProcessReceipt)
	http.HandleFunc("/receipts/", handler.GetPoints)

	http.ListenAndServe(":8080", nil)
}
