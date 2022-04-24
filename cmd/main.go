package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "https://google.com"

	check(url)
}

func check(url string) {
	resp, err := http.Get(url)
	for err != nil {
		fmt.Println("No network connection")
		fmt.Println(time.Now())
		time.Sleep(1 * time.Minute)
		resp, err = http.Get(url)
	}

	defer resp.Body.Close()

}
