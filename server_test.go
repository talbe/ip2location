package main

import (
	"testing"
	"github.com/talbe/src/github.com/IpLocation/network"
	"github.com/talbe/src/github.com/IpLocation/models"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func runServer() {
	var server network.SimpleServer
	server.Run()
}

func TestSanity(t *testing.T) {
	go runServer()

	resp, err := http.Get("http://localhost:8080/v1/find-country?ip=1.1.1.1")

	if err != nil {
		t.Errorf("Call failed")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var responseLocation models.Location
	err = json.Unmarshal(body, &responseLocation)
	if err != nil {
		t.Errorf("Bad json received %s", err)
	}

	if responseLocation.Country != "Israel" || responseLocation.City != "Tel-Aviv" {
		t.Errorf("Bad item received")
	}
}

func TestNotExisting(t *testing.T) {
	go runServer()

	resp, err := http.Get("http://localhost:8080/v1/find-country?ip=1.1.1.0")

	if err != nil {
		t.Errorf("Call failed")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var responseError models.HttpError
	err = json.Unmarshal(body, &responseError)
	if err != nil {
		t.Errorf("Bad json received %s", err)
	}

	if responseError.Error != 404 {
		t.Errorf("Bad item received")
	}
}

func TestBadInput(t *testing.T) {
	go runServer()

	resp, err := http.Get("http://localhost:8080/v1/find-country?badParam=1.1.1.0")

	if err != nil {
		t.Errorf("Call failed")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var responseError models.HttpError
	err = json.Unmarshal(body, &responseError)
	if err != nil {
		t.Errorf("Bad json received %s", err)
	}

	if responseError.Error != 400 {
		t.Errorf("Bad item received %d" , responseError.Error)
	}
}

func TestRateLimit(t *testing.T) {

	go runServer()

	resp, err := http.Get("http://localhost:8080/v1/find-country?ip=1.1.1.1")
	for i := 0; i < 5; i++ {
		resp, err = http.Get("http://localhost:8080/v1/find-country?ip=1.1.1.1")
	}

	if err != nil {
		t.Errorf("Call failed")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var responseError models.HttpError
	err = json.Unmarshal(body, &responseError)
	if err != nil {
		t.Errorf("Bad json received %s", err)
	}

	if responseError.Error != 429{
		t.Errorf("Bad item received %d" , responseError.Error)
	}
}