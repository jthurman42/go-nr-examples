package main

import (
	"io"
	"net/http"

	newrelic "github.com/newrelic/go-agent"
)

func external(w http.ResponseWriter, r *http.Request) {
	url := "http://example.com/"
	txn, _ := w.(newrelic.Transaction)
	// This demonstrates an external segment where only the URL is known. If
	// an http.Request is accessible then `StartExternalSegment` is
	// recommended. See the implementation of `NewRoundTripper` for an
	// example.
	es := newrelic.ExternalSegment{
		StartTime: newrelic.StartSegmentNow(txn),
		URL:       url,
	}
	defer logIfError(es.End())

	resp, err := http.Get(url)
	if nil != err {
		_, e2 := io.WriteString(w, err.Error())
		logIfError(e2)
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(w, resp.Body)
	logIfError(err)
}

func roundtripper(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	txn, _ := w.(newrelic.Transaction)
	client.Transport = newrelic.NewRoundTripper(txn, nil)
	resp, err := client.Get("http://example.com/")
	if nil != err {
		_, e2 := io.WriteString(w, err.Error())
		logIfError(e2)
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(w, resp.Body)
	logIfError(err)
}
