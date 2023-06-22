package main

import (
	"fmt"
	"net/http"
	"time"
)

var healthCheckUrls = []string{
	"https://kljhasdfhljsad.com/status.php",
	"https://flo.com.tr/status.php",
	"https://ninewest.com.tr/status.php",
	"https://instreet.com.tr/status.php",
	"https://lumberjack.com.tr/status.php",
}

func main() {
	// channel := make(chan string) // unbuffered channel
	channel := make(chan string, len(healthCheckUrls))
	defer close(channel)

	for _, uri := range healthCheckUrls {
		go healthChecker(uri, channel)
	}

	for i := 0; i < len(healthCheckUrls); i++ {
		fmt.Println(<-channel)
	}

	time.Sleep(1000 * time.Millisecond)
}

func healthChecker(uri string, currentChannel chan string) {
	_, responseError := http.Get(uri)
	if responseError != nil {
		currentChannel <- fmt.Sprintf("%s seemas down!\n", uri)
	} else {
		currentChannel <- fmt.Sprintf("%s seemas healthy!\n", uri)
	}
}
