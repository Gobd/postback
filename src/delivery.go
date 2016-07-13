package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func makeLogger(file string) {
	errorLog, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic("error opening file")
	}
	log.SetOutput(errorLog)
}

// these structs mirror the JSON struture that Redis will give us
type request struct {
	Endpoint endpointData      `json:"endpoint"`
	Data     map[string]string `json:"data"`
}

type endpointData struct {
	Method string `json:"method"`
	URL    string `json:"url"`
}

func braceReplace(key, val string) (string, string) {
	return key + "={" + key + "}", key + "=" + val
}

// returns url that should recieve get requests
func getReq(decodedJSON request) string {
	removeEmptyBraceRegex := regexp.MustCompile("{[[:word:]]*}")
	formattedURL := decodedJSON.Endpoint.URL
	for key, val := range decodedJSON.Data {
		old, new := braceReplace(key, val)
		formattedURL = strings.Replace(formattedURL, old, new, 1)
	}
	formattedURL = removeEmptyBraceRegex.ReplaceAllString(formattedURL, "")
	return formattedURL
}

// returns the url to post to, and the data to send with it
func postReq(decodedJSON request) (string, map[string]string) {
	return decodedJSON.Endpoint.URL, decodedJSON.Data
}

// checks the method in the request, then sends the request via the appropriate method
// also logs time the request took, and pertinent information from the http response
func sendRequest(data request, timeStart time.Time) {
	if data.Endpoint.Method == "GET" {
		getURL := getReq(data)
		resp, err := http.Get(getURL)
		fmt.Println("get resp", resp)
		fmt.Println("get err", err)
	} else if data.Endpoint.Method == "POST" {
		postURL, postData := postReq(data)
		postDataJSON, _ := json.Marshal(postData)
		resp, err := http.Post(postURL, "application/json", bytes.NewBuffer(postDataJSON))
		fmt.Println("post resp", resp)
		fmt.Println("post err", err)
	}
}

func main() {
	client := redisClient()
	makeLogger("go.log")

	_, errPing := client.Ping().Result()
	if errPing != nil {
		log.Println("Connected to database.")
	}

	for {
		if str, errPop := client.BLPop(0, "requests").Result(); errPop == nil {
			timeStart := time.Now()
			var decodedData request
			jsonErr := json.Unmarshal([]byte(str[1]), &decodedData)
			if jsonErr != nil {
				panic(jsonErr)
			}
			sendRequest(decodedData, timeStart)
		} else {
			panic("Error popping")
		}
	}
}
