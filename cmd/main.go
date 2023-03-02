package main

import (
	"fmt"
	"time"

	"github.com/lanpaiva/multithreading/controller"
)

func main() {
	cep := "1111111"

	api1 := "https://cdn.apicep.com/file/apicep/" + cep + ".json"
	api2 := "http://viacep.com.br/ws/" + cep + "/json/"

	ch1 := make(chan string)
	ch2 := make(chan string)

	go controller.Fetch(api1, ch1)
	go controller.Fetch(api2, ch2)

	select {
	case api1 := <-ch1:
		fmt.Println("api 1 accepted", api1)
	case api2 := <-ch2:
		fmt.Println("api 2 accepted", api2)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}
