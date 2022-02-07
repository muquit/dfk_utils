package main

// print defikingdom's profile information from blockchain
//  muquit@muquit.com - February-06-2022 15:59:57
import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/muquit/dfk_utils/pkg/profile"
)

const (
	profileContractAddress = "0xabD4741948374b1f5DD5Dd7599AC1f85A34cAcDD"
	mainNet = "https://api.harmony.one"
)

func main() {
	// -addr
	var profileHexAddr string
	flag.StringVar(&profileHexAddr,"addr","","Profile Hex Address")
	flag.Parse()

	if len(profileHexAddr) == 0 {
		flag.Usage()
		os.Exit(0)		
	}
	// connect to harmony.one blockchain
	client, err := ethclient.Dial(mainNet)
	if err != nil {
		log.Fatal(err)
	}

	pinfo, err := profile.GetProfileInfo(client, profileHexAddr)
	if err != nil {
		log.Fatal(err)
	}
	bytesArray, err := json.MarshalIndent(pinfo, "", "   ")
	fmt.Println(string(bytesArray))
}