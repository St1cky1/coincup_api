package main

import (
	coincap "coincap_api_client/coincap_client"
	"fmt"
	"log"
	"time"
)

func main() {
	coinCapClient, err := coincap.NewClient(time.Second * 10)

	if err != nil {
		log.Fatal(err)
	}

	assets, err := coinCapClient.GetAssets()

	if err != nil {
		log.Fatal(err)
	}

	for _, asset := range assets {
		fmt.Println(asset.Info())
	}

	// вывод отдельной криптовалюты по названию
	// bitcoin, err := coinCapClient.GetAsset("bitcoin")
	// if err != nil{
	// 	log.Fatal(err)
	// }

	// fmt.Println(bitcoin.Info())
}
