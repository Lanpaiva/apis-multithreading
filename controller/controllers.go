package controller

import (
	"fmt"
	"net/http"
	"time"
)

func Fetch(url string, ch chan<- string) {
	start := time.Now()
	client := http.Client{Timeout: time.Second}
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	resp.Body.Close()

	elapsed := time.Since(start).Seconds()
	time := fmt.Sprintf("response was %.2fs: ", elapsed)

	ch <- fmt.Sprint(time, url)

}
