package main

import (
	"fmt"
	"log"
	"os"

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

var errLog *os.File
var logger *log.Logger

func startLogger() {
	errLog, err := os.OpenFile(errLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	logger = log.New(errLog, "applog: ", log.Lshortfile|log.LstdFlags)
	return logger
}

func main() {
	fmt.Println("hey")
	client := redisClient()
	logFile := setupLogger()

}
