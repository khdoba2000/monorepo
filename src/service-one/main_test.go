package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestSrvice(t *testing.T) {
	go main()
	time.Sleep(time.Second) // Leave time for service to stat
	resp, err := http.Get("http://localhost:8080/one/hello")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	expected := "Hello, World!"
	actual, _ := ioutil.ReadAll(resp.Body)
	if expected != string(actual) {
		fmt.Println("expected:", expected)
		fmt.Println("actual:", string(actual))
		t.Fail()
	}

}
