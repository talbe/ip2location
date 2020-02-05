package handlers

import (
	"github.com/IpLocation/data_access"
	"github.com/IpLocation/models"
)

type IpLocation struct {
}

func (this *IpLocation) Handle(ip string) (models.Location, error) {
	return data_access.SimpleFileInstance().GetLocation(ip)

}