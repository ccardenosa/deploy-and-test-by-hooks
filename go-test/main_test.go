package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/ccardenosa/deploy-and-test-by-hooks/go-test/backend"
)

func TestFrontend(t *testing.T) {

	t.Log("Start Frontend server")
	go backend.StartBackend(beConfig)
	time.Sleep(5)

	resp, err := http.Get("http://localhost:28081/languages")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Error("Not OK Response status:", resp.Status)
	} else {
		t.Log("Response status:", resp.Status)
	}
}
