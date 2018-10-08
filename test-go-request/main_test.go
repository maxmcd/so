package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteApp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	baseURL = ts.URL + "/"

	DeleteApp("platform", "hostname", "header")

}
