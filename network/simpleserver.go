package network

import (
	"log"
	"net/http"
	"github.com/IpLocation/handlers"
	"github.com/IpLocation/configuration"
	"github.com/IpLocation/models"
	"encoding/json"
)

type SimpleServer struct {
	ipLocationHandler handlers.IpLocation
}

func (this *SimpleServer) Run() error {
	apiVersion := configuration.ConfigInstance().ApiVersion()

	findCountryApi, err := configuration.ConfigInstance().FindCountryApi()
	if err != nil {
		log.Fatal(err)
		return err
	}

	serverPort, err  := configuration.ConfigInstance().ServerPort()
	if err != nil {
		log.Fatal(err)
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/" + apiVersion + "/" +findCountryApi, this.handler)
	http.ListenAndServe(":" + serverPort, limit(mux))

	return nil
}

func setError(writer *http.ResponseWriter, errorCode uint32) {
	var httpError models.HttpError
	httpError.Error = errorCode

	jsnHttpError, err := json.Marshal(httpError)
	if err != nil {
		http.Error(*writer, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	(*writer).Write(jsnHttpError)
}

func (this *SimpleServer) handler(writer http.ResponseWriter, reader *http.Request) {

	ips, ok := reader.URL.Query()["ip"]
	writer.Header().Set("Content-Type", "application/json")

	if !ok || len(ips[0]) < 1 {
		setError(&writer, 400)
		return
	}

	ip := ips[0]
	location, err := this.ipLocationHandler.Handle(ip)

	if _, ok := err.(*models.NotFoundError); ok {
		setError(&writer, 404)
		return

	} else if _, ok := err.(*models.InternalError); ok {
		setError(&writer, 500)
		return
	}

	jsonLocation, err := json.Marshal(location)
	if err != nil {
		setError(&writer, 500)
		return
	}

	writer.Write(jsonLocation)
}