package main

import (
	"github.com/talbe/src/github.com/IpLocation/network"
)

func main() {
	var server network.SimpleQueryStringServer
	server.Run()
}