package network

import (
	"log"
	"net/http"
	"github.com/talbe/src/github.com/IpLocation/handlers"
	"encoding/json"
)

type SimpleQueryStringServer struct {
	ip_location_handler handlers.IpLocation
}

func (this *SimpleQueryStringServer) Run() {
	http.HandleFunc("/v1/find-country", this.handler)
	http.ListenAndServe(":8080", nil)
}

func (this *SimpleQueryStringServer) handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]
	location := this.ip_location_handler.Handle(key)

	x, err := json.Marshal(location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Url Param 'key' is: " + string(key))
	log.Println("Url Param 'key' is: " + r.RemoteAddr)

	w.Header().Set("Content-Type", "application/json")
	w.Write(x)
}