package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/muquit/dfk_utils/pkg/land"
)

const (
	landContractAddress = "0xD5f5bE1037e457727e011ADE9Ca54d21c21a3F8A"
	mainNet = "https://api.harmony.one"
)

func main() {
	// -lid
	var lid int64
	flag.Int64Var(&lid, "lid", 0, "Land Id. If not specified all lands will be displayed")
	flag.Parse()

	landId := big.NewInt(lid)

	// connect to harmony.one blockchain
	client, err := ethclient.Dial(mainNet)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(landContractAddress)

	instance, err := land.NewLand(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// lands
	if big.NewInt(0).Cmp(landId) == 0 {
		landData, err := instance.GetAllLands(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}
		bytearray, err := json.MarshalIndent(landData, "", "   ")
		fmt.Println(string(bytearray))
	} else {
		landData, err := instance.GetLand(&bind.CallOpts{}, landId)
		if err != nil {
			log.Fatal(err)
		}
		bytearray, err := json.MarshalIndent(landData, "", "   ")
		fmt.Println(string(bytearray))
	}


}