package handlers

import (
	"testing"
	"github.com/IpLocation/models"
)

func TestHandleSanity(t *testing.T) {
	var ipLocationHandler IpLocation

	location, err := ipLocationHandler.Handle("1.1.1.1")

	if err != nil{
		t.Errorf("Handle failed")
	}

	if location.City != "Tel-Aviv" || location.Country != "Israel"{
		t.Errorf("Handle return bad value")
	}
}

func TestHandleNotFound(t *testing.T) {
	var ipLocationHandler IpLocation

	_, err := ipLocationHandler.Handle("1.1.1.10")

	if _, ok := err.(*models.NotFoundError); !ok {
		t.Errorf("Should be NotFoundError here")
	}
}