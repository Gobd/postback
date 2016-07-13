package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

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

// Logging to a file probably isn't the best way, would be better to do something like ELK (Elastic, Logstash, Kibana)
func makeLogger(file string) {
	dataLog, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic("Error opening logfile!")
	}
	log.SetOutput(dataLog)
}

// These two structs mirror the JSON struture we will get from Redis
type request struct {
	Endpoint endpointData      `json:"endpoint"`
	Data     map[string]string `json:"data"`
}

type endpointData struct {
	Method string `json:"method"`
	URL    string `json:"url"`
}

// This is used to replace the old (key={xxx}) with the new (key=key) in the URL
func braceReplace(key, val string) (string, string) {
	return key + "={" + key + "}", key + "=" + val
}

// Formats and returns a URL string that should recieve a get request
func getReq(decodedJSON request) string {
	removeEmptyBraceRegex := regexp.MustCompile("{[[:word:]]*}")
	formattedURL := decodedJSON.Endpoint.URL
	for key, val := range decodedJSON.Data {
		old, new := braceReplace(key, val)
		formattedURL = strings.Replace(formattedURL, old, new, 1)
	}
	// The empty quotes are what replace unmatched url {key}s, can change to anything or have function take in value to use
	formattedURL = removeEmptyBraceRegex.ReplaceAllString(formattedURL, "")
	return formattedURL
}

// Returns the URL string to post to, and a data map to send with it
func postReq(decodedJSON request) (string, map[string]string) {
	return decodedJSON.Endpoint.URL, decodedJSON.Data
}

// Checks the method in the request, then sends the request via the appropriate method
// Also logs delivery time, response code, response time, and response body
func sendRequest(data request, timeStart time.Time) {
	if data.Endpoint.Method == "GET" {
		getURL := getReq(data)
		resp, getErr := http.Get(getURL)
		if getErr != nil {
			log.Println("ERROR: ", time.Now(), "GET request error: ", getErr, "to URL: ", getURL)
		} else {
			// this needs to log delivery time, response code, response time, and response body logging
			log.Println(resp)
		}
	} else if data.Endpoint.Method == "POST" {
		postURL, postData := postReq(data)
		postDataJSON, _ := json.Marshal(postData)
		resp, postErr := http.Post(postURL, "application/json", bytes.NewBuffer(postDataJSON))
		if postErr != nil {
			log.Println("ERROR: ", time.Now(), "GET request error: ", postErr, "to URL: ", postURL, "with data: ", postData)
		} else {
			// this needs to log delivery time, response code, response time, and response body logging
			log.Println(resp)
		}
	} else {
		log.Println("WARN: ", time.Now(), "Unknown HTTP method in data: ", data)
	}
}

func main() {
	client := redisClient()
	// Change log file here
	makeLogger("go.log")

	_, errPing := client.Ping().Result()
	if errPing != nil {
		log.Println("ERROR: ", time.Now(), "Error connecting to Redis: ", errPing)
	}

	for {
		if popData, errPop := client.BLPop(0, "requests").Result(); errPop == nil {
			timeStart := time.Now()
			var decodedData request
			jsonErr := json.Unmarshal([]byte(popData[1]), &decodedData)
			if jsonErr == nil {
				sendRequest(decodedData, timeStart)
			} else {
				log.Println("ERROR: ", time.Now(), "Error decoding JSON from Redis: ", jsonErr)
			}
		} else {
			log.Println("ERROR: ", time.Now(), "Error popping data from Redis: ", errPop)
		}
	}
}
