package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	newrelic "github.com/newrelic/go-agent"
)

var (
	Version = "undefined"
	app     newrelic.Application
)

// index prints the index page content
func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "hello world")
	logIfError(err)
}

func main() {
	cfg := newrelic.NewConfig("Example App", mustGetEnv("NEW_RELIC_LICENSE_KEY"))
	cfg.Logger = newrelic.NewDebugLogger(os.Stdout)

	var err error
	app, err = newrelic.NewApplication(cfg)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", index))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/notice_error", noticeError))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/notice_error_with_attributes", noticeErrorWithAttributes))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/add_attribute", addAttribute))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/ignore", ignore))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/segments", segments))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/mysql", mysql))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/external", external))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/roundtripper", roundtripper))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/custommetric", customMetric))
	http.HandleFunc("/background", background)

	err = http.ListenAndServe(":8000", nil)
	logIfError(err)
}
