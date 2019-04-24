package main

import (
	"io"
	"net/http"

	newrelic "github.com/newrelic/go-agent"
)

func addAttribute(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "adding attributes")
	logIfError(err)

	if txn, ok := w.(newrelic.Transaction); ok {
		logIfError(txn.AddAttribute("myString", "hello"))
		logIfError(txn.AddAttribute("myInt", 123))
	}
}

func customMetric(w http.ResponseWriter, r *http.Request) {
	for _, vals := range r.Header {
		for _, v := range vals {
			// This custom metric will have the name
			// "Custom/HeaderLength" in the New Relic UI.
			err := app.RecordCustomMetric("HeaderLength", float64(len(v)))
			logIfError(err)
		}
	}

	_, err := io.WriteString(w, "custom metric recorded")
	logIfError(err)
}
