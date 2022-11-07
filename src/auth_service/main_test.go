package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	go main()
	time.Sleep(time.Second) // Leave time for service to start
	resp, err := http.Get("http://localhost:8084/auth-service/hello")
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}
	expected := "Hello, World!"
	actual, _ := io.ReadAll(resp.Body)
	if expected != string(actual) {
		fmt.Println("expected", expected)
		fmt.Println("actual:", string(actual))
		t.Fail()
	}
}
