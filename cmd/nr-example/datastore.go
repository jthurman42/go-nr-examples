package main

import (
	"io"
	"net/http"
	"time"

	newrelic "github.com/newrelic/go-agent"
)

func mysql(w http.ResponseWriter, r *http.Request) {
	txn, _ := w.(newrelic.Transaction)

	s := newrelic.DatastoreSegment{
		StartTime:          newrelic.StartSegmentNow(txn),
		Product:            newrelic.DatastoreMySQL,
		Collection:         "users",
		Operation:          "INSERT",
		ParameterizedQuery: "INSERT INTO users (name, age) VALUES ($1, $2)",
		QueryParameters: map[string]interface{}{
			"name": "Dracula",
			"age":  439,
		},
		Host:         "mysql-server-1",
		PortPathOrID: "3306",
		DatabaseName: "my_database",
	}

	defer logIfError(s.End())

	time.Sleep(20 * time.Millisecond)
	_, err := io.WriteString(w, `performing fake query "INSERT * from users"`)
	logIfError(err)
}
