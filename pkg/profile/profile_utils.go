package profile

// Helper methods for Profile Info
// Contract ABI and address is taken from https://github.com/0rtis/dfk
// Python project
//  muquit@muquit.com - February-06-2022 18:14:46
import (
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)
const (
	profileContractHexAddress = "0xabD4741948374b1f5DD5Dd7599AC1f85A34cAcDD"
)

type ProfileInfo struct {
	Id      *big.Int `json:"id"`
	Owner   common.Address `json:"owner_address"`
	Name    string `json:"name"`
	Created uint64 `json:"created_on_epoch"`
	CreatedTimeLocal string `json:"created_on_time_local"`
	PicId   uint8 `json:"pic_id"`
	HeroId  *big.Int `json:"hero_id"`
}

func GetProfileContractAddress() (common.Address){
	return common.HexToAddress(profileContractHexAddress)
}

func GetProfileInfo(client *ethclient.Client, profileHexAddr string)(profileInfo ProfileInfo, err error) {
	var pinfo ProfileInfo
	profileContractAddress := common.HexToAddress(profileContractHexAddress)
	instance, err := NewProfile(profileContractAddress, client)
	if err != nil {
		return pinfo, errors.New("Could not create Profile Object from Smart Contract")
	}
	profileAddress := common.HexToAddress(profileHexAddr)
	p, err := instance.GetProfileByAddress(&bind.CallOpts{}, profileAddress)
	if err != nil {
		return pinfo, errors.New("Could not Get Profile by address")
	}

	// map to our Struct for Pretty Print JSON
	pinfo.Id = p.Id
	pinfo.Owner = p.Owner
	pinfo.Name = p.Name
	pinfo.Created = p.Created
	t := time.Unix(int64(p.Created),0)
	pinfo.CreatedTimeLocal = t.String()
	pinfo.PicId = p.PicId
	pinfo.HeroId = p.HeroId
	return pinfo, nil
}