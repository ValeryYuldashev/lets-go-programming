package main

import (
	"fmt"
	"log"

	"github.com/mdlayher/wifi"

	myWifi "example_mock/internal/wifi"
)

func main() {
	wifiClient, err := wifi.New()
	if err != nil {
		log.Fatal("Ошибка при создании wifiClient: ", err.Error())
	}

	wifiService := myWifi.New(wifiClient)

	addrs, err := wifiService.GetAddresses()
	if err != nil {
		log.Fatal("Ошибка при получении адресов: ", err.Error())
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
