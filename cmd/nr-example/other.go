package main

import (
	"io"
	"math/rand"
	"net/http"
	"time"

	newrelic "github.com/newrelic/go-agent"
)

func background(w http.ResponseWriter, r *http.Request) {
	// Transactions started without an http.Request are classified as
	// background transactions.
	txn := app.StartTransaction("background", nil, nil)
	defer logIfError(txn.End())

	_, err := io.WriteString(w, "background transaction")
	logIfError(err)
	time.Sleep(150 * time.Millisecond)
}

// ignore shows you how to
func ignore(w http.ResponseWriter, r *http.Request) {
	// flip a coin
	if rand.Intn(2) == 0 {
		if txn, ok := w.(newrelic.Transaction); ok {
			logIfError(txn.Ignore())
		}
		_, err := io.WriteString(w, "ignoring the transaction")
		logIfError(err)
	} else {
		_, err := io.WriteString(w, "not ignoring the transaction")
		logIfError(err)
	}
}
