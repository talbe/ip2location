package handlers

import (
	"github.com/talbe/src/github.com/IpLocation/data_access"
	"github.com/talbe/src/github.com/IpLocation/models"
)

type IpLocation struct {
	x int;
}

func (this *IpLocation) Handle(ip string) models.Location {
	return data_access.SimpleFileInstance().GetLocation(ip)

}