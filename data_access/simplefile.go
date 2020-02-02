package data_access

import (
	"sync"
	"github.com/talbe/src/github.com/IpLocation/models"
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
)

type SimpleFile struct{
	locations map[string]models.Location
}

var simpleFile *SimpleFile
var once sync.Once

func SimpleFileInstance() *SimpleFile  {
	once.Do(func() {
		simpleFile = &SimpleFile{}
	})
	return simpleFile
}

func (this *SimpleFile) GetLocation(ip string) models.Location {
	file, err := os.Open("/home/tal/dev/go/src/github.com/talbe/src/github.com/IpLocation/datastores/datastore2.txt")
	if err != nil {
		log.Fatal(err)
	}

	this.locations = make(map[string]models.Location)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		raw_string := strings.Trim(scanner.Text(), "()")

		fmt.Println(raw_string)

		values := strings.Split(raw_string, ",")

		fmt.Println(values[0])
		fmt.Println(values[1])
		fmt.Println(values[2])

		log.Println("adding to locations the ip %s the counter %s and the city %s", values[0], values[2], values[1])
		this.locations[values[0]] = models.Location{Country:values[2], City: values[1]}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return this.locations[ip]
}