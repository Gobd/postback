package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"gopkg.in/redis.v4"
)

func redisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6400",
		Password: "",
		DB:       0,
	})
	return client
}

func makeLogger(file string) {
	errorLog, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic("error opening file")
	}
	log.SetOutput(errorLog)
}

type request struct {
	Endpoint endpointData        `json:"endpoint"`
	Data     []map[string]string `json:"data"`
}

type endpointData struct {
	Method string `json:"method"`
	URL    string `json:"url"`
}

// hopefully this is good mock data, intended to represent what Redis will give me when I BLPop data
const input = `
{
	"endpoint": {
	"method":"GET",
	"url":"http://sample_domain_endpoint.com/data?key={key}&value={value}&foo={bar}"
},
	"data": [
	{
		"key":"Azureus",
		"value":"Dendrobates"
	},
	{
		"key":"Phyllobates",
		"value":"Terribilis"
	}
	]
}`

func braceReplace(key, val string) (string, string) {
	return key + "={" + key + "}", key + "=" + val
}

func getReq(decodedJSON request) []string {
	var getSlice []string
	for _, dataSet := range decodedJSON.Data {
		formattedURL := decodedJSON.Endpoint.URL
		for key, val := range dataSet {
			old, new := braceReplace(key, val)
			formattedURL = strings.Replace(formattedURL, old, new, 1)
		}
		getSlice = append(getSlice, formattedURL)
	}
	return getSlice
}

func postReq(decodedJSON request) {

}

func main() {
	client := redisClient()
	makeLogger("go.log")
	var decodedJSON request
	jsonErr := json.Unmarshal([]byte(input), &decodedJSON)
	getURL := getReq(decodedJSON)
	if jsonErr != nil {
		panic(jsonErr)
	}
	// fmt.Println("METHOD ", decodedJSON.Endpoint.Method)
	// fmt.Println("URL ", decodedJSON.Endpoint.URL)
	_, errPing := client.Ping().Result()
	if errPing != nil {
		log.Println("Connected to database.")
	}
}
