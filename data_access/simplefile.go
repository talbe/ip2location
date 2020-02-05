package data_access

import (
	"sync"
	"github.com/IpLocation/models"
	"os"
	"log"
	"bufio"
	"strings"
	"github.com/IpLocation/configuration"
)

type SimpleFile struct{
	locations map[string]models.Location
	locationsInitialized bool
}

var simpleFile *SimpleFile
var once sync.Once

func SimpleFileInstance() *SimpleFile  {
	once.Do(func() {
		simpleFile = &SimpleFile{locationsInitialized: false}
	})
	return simpleFile
}

func (this *SimpleFile) loadLocations() error {
	dataStorePath, err := configuration.ConfigInstance().DataStorePath()
	if err != nil {
		log.Fatal(err)
		return &models.InternalError{}
	}

	file, err := os.Open(dataStorePath)
	if err != nil {
		log.Fatal(err)
		return &models.InternalError{}
	}
	defer file.Close()

	this.locations = make(map[string]models.Location)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rawString := strings.Trim(scanner.Text(), "()")
		values := strings.Split(rawString, ",")

		log.Println("Adding to locations the ip %s the counter %s and the city %s", values[0], values[2], values[1])
		this.locations[values[0]] = models.Location{Country:values[2], City: values[1]}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return &models.InternalError{}
	}

	return nil
}

func (this *SimpleFile) GetLocation(ip string) (models.Location, error) {

	// In case we have already loaded the locations - do not load again.
	if ! this.locationsInitialized{
		err := this.loadLocations()
		if err != nil {
			return models.Location{}, err
		}

		this.locationsInitialized = true
	}

	location, ok := this.locations[ip]

	if !ok {
		return models.Location{}, &models.NotFoundError{}
	}

	return location, nil
}