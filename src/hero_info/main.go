package main

///////////////////////////////////////////////////////////////////////
// Dump defikingdoms Hero information in JSON
// Adapted from https://github.com/0rtis/dfk.git
//
// Usage: hero_info -hid <hero id> (default is 1)
//
// TODO: Stats and Visual Genes decoding
//
//  muquit@muquit.com - January-16-2022 14:08:11
///////////////////////////////////////////////////////////////////////

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/muquit/dfk_utils/pkg/hero"
)

const (
	herrorContractAddress = "0x5f753dcdf9b1ad9aabc1346614d1f4746fd6ce5c"
	mainNet = "https://api.harmony.one"
)

type HeroInfo struct {
	OwnerAddress common.Address
	HeroInfo hero.IHeroTypesHero
}

func logInfo(format string, a ...interface{}) {
    log.Printf(format, a...)
}

func main() {
	// -hid
	var hid int64
	flag.Int64Var(&hid, "hid", 1, "Hero Id")
	flag.Parse()

	heroId := big.NewInt(hid)

	// connect to harmony.one blockchain
	client, err := ethclient.Dial(mainNet)
	if err != nil {
		log.Fatal(err)
	}

	// create Hero Smart Contract address object
	address := common.HexToAddress(herrorContractAddress)
	instance, err := hero.NewHero(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// get Hero Owner infomation
	ownerAddress, err := instance.OwnerOf(&bind.CallOpts{}, heroId)
	if err != nil {
		log.Fatal(err)
	}

	// get the hero information
	h, err := instance.GetHero(&bind.CallOpts{}, heroId)
	if err != nil {
		log.Fatal(err)
	}
	abuf, _ := hero.AlphabetaizeGenes(h.Info.StatGenes)
	fmt.Println(abuf)
	sg, _ := hero.DecodeStatGenes(h.Info.StatGenes)
	fmt.Println(sg)
//	os.Exit(0)

	// convert the object JSON and print
	hinfo := HeroInfo{OwnerAddress: ownerAddress, HeroInfo: h}
	// decode genes
	byteArray, err := json.MarshalIndent(hinfo, "", "  ")
	fmt.Println(string(byteArray))

	os.Exit(0)
}