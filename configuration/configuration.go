package configuration

import (
	"sync"
	"os"
	"errors"
	"log"
	"strconv"
)

type Configuration struct {
	api_version string
	server_port string
	find_country_api string
	data_store_path  string
}

var simpleFile *Configuration
var once sync.Once
func ConfigInstance() *Configuration  {
	once.Do(func() {
		simpleFile = &Configuration{}
	})
	return simpleFile
}

func (this *Configuration) ApiVersion() string {
	version := os.Getenv("API_VERSION")
	if version == "" {
		version = "v1"
	}

	return version
}

func (this *Configuration) DataStorePath() (string, error) {
	data_store_path := os.Getenv("DATA_STORE_PATH")

	if data_store_path == "" {
		error_str := "No datstore path specified " + data_store_path
		return "", errors.New(error_str)
	}

	return data_store_path, nil
}

func (this *Configuration) FindCountryApi() (string, error) {
	find_country_api := os.Getenv("FIND_COUNTRY_API")
	if find_country_api == "" {
		return "", errors.New("FIND_COUNTRY_API environment variable not exist")
	}

	return find_country_api, nil
}

func (this *Configuration) ServerPort() (string, error) {
	server_port := os.Getenv("SERVER_PORT")
	if server_port == "" {
		return "", errors.New("No datstore path specified")
	}

	return server_port, nil
}

func (this *Configuration) VisitorMinutesToLive() (uint32, error) {
	visitorMinutesToLive := os.Getenv("VISITOR_MINUTES_TO_LIVE")
	if visitorMinutesToLive == "" {
		return 0, errors.New("No visitor minutes to live specified")
	}

	visitorMinutesToLiveInt,err := strconv.Atoi(visitorMinutesToLive)
	if err != nil {
		log.Fatal("Bad VISITOR_MINUTES_TO_LIVE value. should be a number")
		return 0, err
	}

	return uint32(visitorMinutesToLiveInt), nil
}

func (this *Configuration) TokenIncreaseRate() (uint32, error) {
	tokenIncreaseRate := os.Getenv("TOKEN_INCREASE_RATE")
	if tokenIncreaseRate == "" {
		return 0, errors.New("No visitor minutes to live specified")
	}

	tokenIncreaseRateInt,err := strconv.Atoi(tokenIncreaseRate)
	if err != nil {
		log.Fatal("Bad VISITOR_MINUTES_TO_LIVE value. should be a number")
		return 0, err
	}

	return uint32(tokenIncreaseRateInt), nil
}

func (this *Configuration) BucketSize() (uint32, error) {
	bucketSize := os.Getenv("BUCKET_SIZE")
	if bucketSize == "" {
		return 0, errors.New("No visitor minutes to live specified")
	}

	bucketSizeInt,err := strconv.Atoi(bucketSize)
	if err != nil {
		log.Fatal("Bad VISITOR_MINUTES_TO_LIVE value. should be a number")
		return 0, err
	}

	return uint32(bucketSizeInt), nil
}