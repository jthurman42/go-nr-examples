package main

import (
	"errors"
	"io"
	"net/http"

	newrelic "github.com/newrelic/go-agent"
)

func noticeError(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "noticing an error")
	logIfError(err)

	if txn, ok := w.(newrelic.Transaction); ok {
		err = txn.NoticeError(errors.New("my error message"))
		logIfError(err)
	}
}

func noticeErrorWithAttributes(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "noticing an error")
	logIfError(err)

	if txn, ok := w.(newrelic.Transaction); ok {
		err = txn.NoticeError(newrelic.Error{
			Message: "uh oh. something went very wrong",
			Class:   "errors are aggregated by class",
			Attributes: map[string]interface{}{
				"important_number": 97232,
				"relevant_string":  "zap",
			},
		})
		logIfError(err)
	}
}
