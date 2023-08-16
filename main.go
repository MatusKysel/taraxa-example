package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7777/")
	if err != nil {
		log.Panic(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Panic(err)
	}

	for i := header.Number.Int64(); i >= 0; i-- {
		log.Println(i)

		block, err := client.BlockByNumber(context.Background(), big.NewInt(i))
		if err != nil {
			log.Panic(err)
		}

		log.Println(block.Body())
	}
}
