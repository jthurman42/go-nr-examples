package main

import (
	"io"
	"net/http"
	"time"

	newrelic "github.com/newrelic/go-agent"
)

func segments(w http.ResponseWriter, r *http.Request) {
	txn, _ := w.(newrelic.Transaction)

	func() {
		defer logIfError(newrelic.StartSegment(txn, "f1").End())

		func() {
			defer logIfError(newrelic.StartSegment(txn, "f2").End())

			_, err := io.WriteString(w, "segments!")
			logIfError(err)
			time.Sleep(10 * time.Millisecond)
		}()
		time.Sleep(15 * time.Millisecond)
	}()
	time.Sleep(20 * time.Millisecond)
}
