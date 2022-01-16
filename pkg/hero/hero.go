// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hero

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ICrystalTypesHeroCrystal is an auto generated low-level Go binding around an user-defined struct.
type ICrystalTypesHeroCrystal struct {
	Owner          common.Address
	SummonerId     *big.Int
	AssistantId    *big.Int
	Generation     uint16
	CreatedBlock   *big.Int
	HeroId         *big.Int
	SummonerTears  uint8
	AssistantTears uint8
	BonusItem      common.Address
	MaxSummons     uint32
	FirstName      uint32
	LastName       uint32
	ShinyStyle     uint8
}

// IHeroTypesHero is an auto generated low-level Go binding around an user-defined struct.
type IHeroTypesHero struct {
	Id                  *big.Int
	SummoningInfo       IHeroTypesSummoningInfo
	Info                IHeroTypesHeroInfo
	State               IHeroTypesHeroState
	Stats               IHeroTypesHeroStats
	PrimaryStatGrowth   IHeroTypesHeroStatGrowth
	SecondaryStatGrowth IHeroTypesHeroStatGrowth
	Professions         IHeroTypesHeroProfessions
}

// IHeroTypesHeroInfo is an auto generated low-level Go binding around an user-defined struct.
type IHeroTypesHeroInfo struct {
	StatGenes   *big.Int
	VisualGenes *big.Int
	Rarity      uint8
	Shiny       bool
	Generation  uint16
	FirstName   uint32
	LastName    uint32
	ShinyStyle  uint8
	Class       uint8
	SubClass    uint8
}

// IHeroTypesHeroProfessions is an auto generated low-level Go binding around an user-defined struct.
type IHeroTypesHeroProfessions struct {
	Mining    uint16
	Gardening uint16
	Foraging  uint16
	Fishing   uint16
}

// IHeroTypesHeroStatGrowth is an auto generated low-level Go binding around an user-defined struct.
type IHeroTypesHeroStatGrowth struct {
	Strength     uint16
	Intelligence uint16
	Wisdom       uint16
	Luck         uint16
	Agility      uint16
	Vitality     uint16
	Endurance    uint16
	Dexterity    uint16
	HpSm         uint16
	HpRg         uint16
	HpLg         uint16
	MpSm         uint16
	MpRg         uint16
	MpLg         uint16
}

// IHeroTypesHeroState is an auto generated low-level Go binding around an user-defined struct.
type IHeroTypesHeroState struct {
	StaminaFullAt *big.Int
	HpFullAt      *big.Int
	MpFullAt      *big.Int
	Level         uint16
	Xp            uint64
	CurrentQuest  common.Address
	Sp            uint8
	Status        uint8
}

// IHeroTypesHeroStats is an auto generated low-level Go binding around an user-defined struct.
type IHeroTypesHeroStats struct {
	Strength     uint16
	Intelligence uint16
	Wisdom       uint16
	Luck         uint16
	Agility      uint16
	Vitality     uint16
	Endurance    uint16
	Dexterity    uint16
	Hp           uint16
	Mp           uint16
	Stamina      uint16
}

// IHeroTypesSummoningInfo is an auto generated low-level Go binding around an user-defined struct.
type IHeroTypesSummoningInfo struct {
	SummonedTime   *big.Int
	NextSummonTime *big.Int
	SummonerId     *big.Int
	AssistantId    *big.Int
	Summons        uint32
	MaxSummons     uint32
}

// HeroMetaData contains all meta data concerning the Hero contract.
var HeroMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"heroId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"summonerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assistantId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"statGenes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"visualGenes\",\"type\":\"uint256\"}],\"name\":\"HeroSummoned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HERO_MODERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_statGenes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_visualGenes\",\"type\":\"uint256\"},{\"internalType\":\"enumIHeroTypes.Rarity\",\"name\":\"_rarity\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_shiny\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"summonerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assistantId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"generation\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"createdBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"heroId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"summonerTears\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"assistantTears\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bonusItem\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxSummons\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"firstName\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastName\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"shinyStyle\",\"type\":\"uint8\"}],\"internalType\":\"structICrystalTypes.HeroCrystal\",\"name\":\"_crystal\",\"type\":\"tuple\"}],\"name\":\"createHero\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getHero\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"summonedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextSummonTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"summonerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assistantId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"summons\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxSummons\",\"type\":\"uint32\"}],\"internalType\":\"structIHeroTypes.SummoningInfo\",\"name\":\"summoningInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"statGenes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"visualGenes\",\"type\":\"uint256\"},{\"internalType\":\"enumIHeroTypes.Rarity\",\"name\":\"rarity\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"shiny\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"generation\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"firstName\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastName\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"shinyStyle\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"class\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"subClass\",\"type\":\"uint8\"}],\"internalType\":\"structIHeroTypes.HeroInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"staminaFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hpFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mpFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"level\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"xp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"currentQuest\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sp\",\"type\":\"uint8\"},{\"internalType\":\"enumIHeroTypes.HeroStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIHeroTypes.HeroState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"stamina\",\"type\":\"uint16\"}],\"internalType\":\"structIHeroTypes.HeroStats\",\"name\":\"stats\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpLg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpLg\",\"type\":\"uint16\"}],\"internalType\":\"structIHeroTypes.HeroStatGrowth\",\"name\":\"primaryStatGrowth\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpLg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpLg\",\"type\":\"uint16\"}],\"internalType\":\"structIHeroTypes.HeroStatGrowth\",\"name\":\"secondaryStatGrowth\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"mining\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gardening\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"foraging\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"fishing\",\"type\":\"uint16\"}],\"internalType\":\"structIHeroTypes.HeroProfessions\",\"name\":\"professions\",\"type\":\"tuple\"}],\"internalType\":\"structIHeroTypes.Hero\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserHeroes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_statScienceAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseTokenURI\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_statScienceAddress\",\"type\":\"address\"}],\"name\":\"setStatScienceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"summonedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextSummonTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"summonerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assistantId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"summons\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxSummons\",\"type\":\"uint32\"}],\"internalType\":\"structIHeroTypes.SummoningInfo\",\"name\":\"summoningInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"statGenes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"visualGenes\",\"type\":\"uint256\"},{\"internalType\":\"enumIHeroTypes.Rarity\",\"name\":\"rarity\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"shiny\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"generation\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"firstName\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastName\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"shinyStyle\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"class\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"subClass\",\"type\":\"uint8\"}],\"internalType\":\"structIHeroTypes.HeroInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"staminaFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hpFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mpFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"level\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"xp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"currentQuest\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sp\",\"type\":\"uint8\"},{\"internalType\":\"enumIHeroTypes.HeroStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIHeroTypes.HeroState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"stamina\",\"type\":\"uint16\"}],\"internalType\":\"structIHeroTypes.HeroStats\",\"name\":\"stats\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpLg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpLg\",\"type\":\"uint16\"}],\"internalType\":\"structIHeroTypes.HeroStatGrowth\",\"name\":\"primaryStatGrowth\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpLg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpLg\",\"type\":\"uint16\"}],\"internalType\":\"structIHeroTypes.HeroStatGrowth\",\"name\":\"secondaryStatGrowth\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"mining\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gardening\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"foraging\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"fishing\",\"type\":\"uint16\"}],\"internalType\":\"structIHeroTypes.HeroProfessions\",\"name\":\"professions\",\"type\":\"tuple\"}],\"internalType\":\"structIHeroTypes.Hero\",\"name\":\"_hero\",\"type\":\"tuple\"}],\"name\":\"updateHero\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userHeroes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HeroABI is the input ABI used to generate the binding from.
// Deprecated: Use HeroMetaData.ABI instead.
var HeroABI = HeroMetaData.ABI

// Hero is an auto generated Go binding around an Ethereum contract.
type Hero struct {
	HeroCaller     // Read-only binding to the contract
	HeroTransactor // Write-only binding to the contract
	HeroFilterer   // Log filterer for contract events
}

// HeroCaller is an auto generated read-only Go binding around an Ethereum contract.
type HeroCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeroTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HeroTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeroFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HeroFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeroSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HeroSession struct {
	Contract     *Hero             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeroCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HeroCallerSession struct {
	Contract *HeroCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HeroTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HeroTransactorSession struct {
	Contract     *HeroTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeroRaw is an auto generated low-level Go binding around an Ethereum contract.
type HeroRaw struct {
	Contract *Hero // Generic contract binding to access the raw methods on
}

// HeroCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HeroCallerRaw struct {
	Contract *HeroCaller // Generic read-only contract binding to access the raw methods on
}

// HeroTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HeroTransactorRaw struct {
	Contract *HeroTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHero creates a new instance of Hero, bound to a specific deployed contract.
func NewHero(address common.Address, backend bind.ContractBackend) (*Hero, error) {
	contract, err := bindHero(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hero{HeroCaller: HeroCaller{contract: contract}, HeroTransactor: HeroTransactor{contract: contract}, HeroFilterer: HeroFilterer{contract: contract}}, nil
}

// NewHeroCaller creates a new read-only instance of Hero, bound to a specific deployed contract.
func NewHeroCaller(address common.Address, caller bind.ContractCaller) (*HeroCaller, error) {
	contract, err := bindHero(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HeroCaller{contract: contract}, nil
}

// NewHeroTransactor creates a new write-only instance of Hero, bound to a specific deployed contract.
func NewHeroTransactor(address common.Address, transactor bind.ContractTransactor) (*HeroTransactor, error) {
	contract, err := bindHero(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HeroTransactor{contract: contract}, nil
}

// NewHeroFilterer creates a new log filterer instance of Hero, bound to a specific deployed contract.
func NewHeroFilterer(address common.Address, filterer bind.ContractFilterer) (*HeroFilterer, error) {
	contract, err := bindHero(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HeroFilterer{contract: contract}, nil
}

// bindHero binds a generic wrapper to an already deployed contract.
func bindHero(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HeroABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hero *HeroRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hero.Contract.HeroCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hero *HeroRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hero.Contract.HeroTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hero *HeroRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hero.Contract.HeroTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hero *HeroCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hero.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hero *HeroTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hero.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hero *HeroTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hero.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Hero *HeroCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Hero *HeroSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Hero.Contract.DEFAULTADMINROLE(&_Hero.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Hero *HeroCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Hero.Contract.DEFAULTADMINROLE(&_Hero.CallOpts)
}

// HEROMODERATORROLE is a free data retrieval call binding the contract method 0x39ab52d5.
//
// Solidity: function HERO_MODERATOR_ROLE() view returns(bytes32)
func (_Hero *HeroCaller) HEROMODERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "HERO_MODERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HEROMODERATORROLE is a free data retrieval call binding the contract method 0x39ab52d5.
//
// Solidity: function HERO_MODERATOR_ROLE() view returns(bytes32)
func (_Hero *HeroSession) HEROMODERATORROLE() ([32]byte, error) {
	return _Hero.Contract.HEROMODERATORROLE(&_Hero.CallOpts)
}

// HEROMODERATORROLE is a free data retrieval call binding the contract method 0x39ab52d5.
//
// Solidity: function HERO_MODERATOR_ROLE() view returns(bytes32)
func (_Hero *HeroCallerSession) HEROMODERATORROLE() ([32]byte, error) {
	return _Hero.Contract.HEROMODERATORROLE(&_Hero.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Hero *HeroCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Hero *HeroSession) MINTERROLE() ([32]byte, error) {
	return _Hero.Contract.MINTERROLE(&_Hero.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Hero *HeroCallerSession) MINTERROLE() ([32]byte, error) {
	return _Hero.Contract.MINTERROLE(&_Hero.CallOpts)
}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_Hero *HeroCaller) MODERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "MODERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_Hero *HeroSession) MODERATORROLE() ([32]byte, error) {
	return _Hero.Contract.MODERATORROLE(&_Hero.CallOpts)
}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_Hero *HeroCallerSession) MODERATORROLE() ([32]byte, error) {
	return _Hero.Contract.MODERATORROLE(&_Hero.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Hero *HeroCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Hero *HeroSession) PAUSERROLE() ([32]byte, error) {
	return _Hero.Contract.PAUSERROLE(&_Hero.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Hero *HeroCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Hero.Contract.PAUSERROLE(&_Hero.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Hero *HeroCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Hero *HeroSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Hero.Contract.BalanceOf(&_Hero.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Hero *HeroCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Hero.Contract.BalanceOf(&_Hero.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Hero *HeroCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Hero *HeroSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Hero.Contract.GetApproved(&_Hero.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Hero *HeroCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Hero.Contract.GetApproved(&_Hero.CallOpts, tokenId)
}

// GetHero is a free data retrieval call binding the contract method 0x21d80111.
//
// Solidity: function getHero(uint256 _id) view returns((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)))
func (_Hero *HeroCaller) GetHero(opts *bind.CallOpts, _id *big.Int) (IHeroTypesHero, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "getHero", _id)

	if err != nil {
		return *new(IHeroTypesHero), err
	}

	out0 := *abi.ConvertType(out[0], new(IHeroTypesHero)).(*IHeroTypesHero)

	return out0, err

}

// GetHero is a free data retrieval call binding the contract method 0x21d80111.
//
// Solidity: function getHero(uint256 _id) view returns((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)))
func (_Hero *HeroSession) GetHero(_id *big.Int) (IHeroTypesHero, error) {
	return _Hero.Contract.GetHero(&_Hero.CallOpts, _id)
}

// GetHero is a free data retrieval call binding the contract method 0x21d80111.
//
// Solidity: function getHero(uint256 _id) view returns((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)))
func (_Hero *HeroCallerSession) GetHero(_id *big.Int) (IHeroTypesHero, error) {
	return _Hero.Contract.GetHero(&_Hero.CallOpts, _id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Hero *HeroCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Hero *HeroSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Hero.Contract.GetRoleAdmin(&_Hero.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Hero *HeroCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Hero.Contract.GetRoleAdmin(&_Hero.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Hero *HeroCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Hero *HeroSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Hero.Contract.GetRoleMember(&_Hero.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Hero *HeroCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Hero.Contract.GetRoleMember(&_Hero.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Hero *HeroCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Hero *HeroSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Hero.Contract.GetRoleMemberCount(&_Hero.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Hero *HeroCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Hero.Contract.GetRoleMemberCount(&_Hero.CallOpts, role)
}

// GetUserHeroes is a free data retrieval call binding the contract method 0xcdf7d856.
//
// Solidity: function getUserHeroes(address _address) view returns(uint256[])
func (_Hero *HeroCaller) GetUserHeroes(opts *bind.CallOpts, _address common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "getUserHeroes", _address)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserHeroes is a free data retrieval call binding the contract method 0xcdf7d856.
//
// Solidity: function getUserHeroes(address _address) view returns(uint256[])
func (_Hero *HeroSession) GetUserHeroes(_address common.Address) ([]*big.Int, error) {
	return _Hero.Contract.GetUserHeroes(&_Hero.CallOpts, _address)
}

// GetUserHeroes is a free data retrieval call binding the contract method 0xcdf7d856.
//
// Solidity: function getUserHeroes(address _address) view returns(uint256[])
func (_Hero *HeroCallerSession) GetUserHeroes(_address common.Address) ([]*big.Int, error) {
	return _Hero.Contract.GetUserHeroes(&_Hero.CallOpts, _address)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Hero *HeroCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Hero *HeroSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Hero.Contract.HasRole(&_Hero.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Hero *HeroCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Hero.Contract.HasRole(&_Hero.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Hero *HeroCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Hero *HeroSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Hero.Contract.IsApprovedForAll(&_Hero.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Hero *HeroCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Hero.Contract.IsApprovedForAll(&_Hero.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hero *HeroCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hero *HeroSession) Name() (string, error) {
	return _Hero.Contract.Name(&_Hero.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hero *HeroCallerSession) Name() (string, error) {
	return _Hero.Contract.Name(&_Hero.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Hero *HeroCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Hero *HeroSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Hero.Contract.OwnerOf(&_Hero.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Hero *HeroCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Hero.Contract.OwnerOf(&_Hero.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Hero *HeroCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Hero *HeroSession) Paused() (bool, error) {
	return _Hero.Contract.Paused(&_Hero.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Hero *HeroCallerSession) Paused() (bool, error) {
	return _Hero.Contract.Paused(&_Hero.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Hero *HeroCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Hero *HeroSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Hero.Contract.SupportsInterface(&_Hero.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Hero *HeroCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Hero.Contract.SupportsInterface(&_Hero.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hero *HeroCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hero *HeroSession) Symbol() (string, error) {
	return _Hero.Contract.Symbol(&_Hero.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hero *HeroCallerSession) Symbol() (string, error) {
	return _Hero.Contract.Symbol(&_Hero.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Hero *HeroCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Hero *HeroSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Hero.Contract.TokenByIndex(&_Hero.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Hero *HeroCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Hero.Contract.TokenByIndex(&_Hero.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Hero *HeroCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Hero *HeroSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Hero.Contract.TokenOfOwnerByIndex(&_Hero.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Hero *HeroCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Hero.Contract.TokenOfOwnerByIndex(&_Hero.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Hero *HeroCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Hero *HeroSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Hero.Contract.TokenURI(&_Hero.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Hero *HeroCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Hero.Contract.TokenURI(&_Hero.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hero *HeroCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hero *HeroSession) TotalSupply() (*big.Int, error) {
	return _Hero.Contract.TotalSupply(&_Hero.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hero *HeroCallerSession) TotalSupply() (*big.Int, error) {
	return _Hero.Contract.TotalSupply(&_Hero.CallOpts)
}

// UserHeroes is a free data retrieval call binding the contract method 0x2d55f737.
//
// Solidity: function userHeroes(address , uint256 ) view returns(uint256)
func (_Hero *HeroCaller) UserHeroes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hero.contract.Call(opts, &out, "userHeroes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserHeroes is a free data retrieval call binding the contract method 0x2d55f737.
//
// Solidity: function userHeroes(address , uint256 ) view returns(uint256)
func (_Hero *HeroSession) UserHeroes(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Hero.Contract.UserHeroes(&_Hero.CallOpts, arg0, arg1)
}

// UserHeroes is a free data retrieval call binding the contract method 0x2d55f737.
//
// Solidity: function userHeroes(address , uint256 ) view returns(uint256)
func (_Hero *HeroCallerSession) UserHeroes(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Hero.Contract.UserHeroes(&_Hero.CallOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Hero *HeroTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Hero *HeroSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hero.Contract.Approve(&_Hero.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Hero *HeroTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hero.Contract.Approve(&_Hero.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Hero *HeroTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Hero *HeroSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Hero.Contract.Burn(&_Hero.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Hero *HeroTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Hero.Contract.Burn(&_Hero.TransactOpts, tokenId)
}

// CreateHero is a paid mutator transaction binding the contract method 0xf6bde81d.
//
// Solidity: function createHero(uint256 _statGenes, uint256 _visualGenes, uint8 _rarity, bool _shiny, (address,uint256,uint256,uint16,uint256,uint256,uint8,uint8,address,uint32,uint32,uint32,uint8) _crystal) returns(uint256)
func (_Hero *HeroTransactor) CreateHero(opts *bind.TransactOpts, _statGenes *big.Int, _visualGenes *big.Int, _rarity uint8, _shiny bool, _crystal ICrystalTypesHeroCrystal) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "createHero", _statGenes, _visualGenes, _rarity, _shiny, _crystal)
}

// CreateHero is a paid mutator transaction binding the contract method 0xf6bde81d.
//
// Solidity: function createHero(uint256 _statGenes, uint256 _visualGenes, uint8 _rarity, bool _shiny, (address,uint256,uint256,uint16,uint256,uint256,uint8,uint8,address,uint32,uint32,uint32,uint8) _crystal) returns(uint256)
func (_Hero *HeroSession) CreateHero(_statGenes *big.Int, _visualGenes *big.Int, _rarity uint8, _shiny bool, _crystal ICrystalTypesHeroCrystal) (*types.Transaction, error) {
	return _Hero.Contract.CreateHero(&_Hero.TransactOpts, _statGenes, _visualGenes, _rarity, _shiny, _crystal)
}

// CreateHero is a paid mutator transaction binding the contract method 0xf6bde81d.
//
// Solidity: function createHero(uint256 _statGenes, uint256 _visualGenes, uint8 _rarity, bool _shiny, (address,uint256,uint256,uint16,uint256,uint256,uint8,uint8,address,uint32,uint32,uint32,uint8) _crystal) returns(uint256)
func (_Hero *HeroTransactorSession) CreateHero(_statGenes *big.Int, _visualGenes *big.Int, _rarity uint8, _shiny bool, _crystal ICrystalTypesHeroCrystal) (*types.Transaction, error) {
	return _Hero.Contract.CreateHero(&_Hero.TransactOpts, _statGenes, _visualGenes, _rarity, _shiny, _crystal)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Hero *HeroTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Hero *HeroSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Hero.Contract.GrantRole(&_Hero.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Hero *HeroTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Hero.Contract.GrantRole(&_Hero.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x5c6d8da1.
//
// Solidity: function initialize(string _name, string _symbol, string _url, address _statScienceAddress) returns()
func (_Hero *HeroTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _url string, _statScienceAddress common.Address) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "initialize", _name, _symbol, _url, _statScienceAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x5c6d8da1.
//
// Solidity: function initialize(string _name, string _symbol, string _url, address _statScienceAddress) returns()
func (_Hero *HeroSession) Initialize(_name string, _symbol string, _url string, _statScienceAddress common.Address) (*types.Transaction, error) {
	return _Hero.Contract.Initialize(&_Hero.TransactOpts, _name, _symbol, _url, _statScienceAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x5c6d8da1.
//
// Solidity: function initialize(string _name, string _symbol, string _url, address _statScienceAddress) returns()
func (_Hero *HeroTransactorSession) Initialize(_name string, _symbol string, _url string, _statScienceAddress common.Address) (*types.Transaction, error) {
	return _Hero.Contract.Initialize(&_Hero.TransactOpts, _name, _symbol, _url, _statScienceAddress)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa6487c53.
//
// Solidity: function initialize(string name, string symbol, string baseTokenURI) returns()
func (_Hero *HeroTransactor) Initialize0(opts *bind.TransactOpts, name string, symbol string, baseTokenURI string) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "initialize0", name, symbol, baseTokenURI)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa6487c53.
//
// Solidity: function initialize(string name, string symbol, string baseTokenURI) returns()
func (_Hero *HeroSession) Initialize0(name string, symbol string, baseTokenURI string) (*types.Transaction, error) {
	return _Hero.Contract.Initialize0(&_Hero.TransactOpts, name, symbol, baseTokenURI)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xa6487c53.
//
// Solidity: function initialize(string name, string symbol, string baseTokenURI) returns()
func (_Hero *HeroTransactorSession) Initialize0(name string, symbol string, baseTokenURI string) (*types.Transaction, error) {
	return _Hero.Contract.Initialize0(&_Hero.TransactOpts, name, symbol, baseTokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_Hero *HeroTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_Hero *HeroSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Hero.Contract.Mint(&_Hero.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_Hero *HeroTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Hero.Contract.Mint(&_Hero.TransactOpts, to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Hero *HeroTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Hero *HeroSession) Pause() (*types.Transaction, error) {
	return _Hero.Contract.Pause(&_Hero.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Hero *HeroTransactorSession) Pause() (*types.Transaction, error) {
	return _Hero.Contract.Pause(&_Hero.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Hero *HeroTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Hero *HeroSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Hero.Contract.RenounceRole(&_Hero.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Hero *HeroTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Hero.Contract.RenounceRole(&_Hero.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Hero *HeroTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Hero *HeroSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Hero.Contract.RevokeRole(&_Hero.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Hero *HeroTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Hero.Contract.RevokeRole(&_Hero.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Hero *HeroTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Hero *HeroSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hero.Contract.SafeTransferFrom(&_Hero.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Hero *HeroTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hero.Contract.SafeTransferFrom(&_Hero.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Hero *HeroTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Hero *HeroSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Hero.Contract.SafeTransferFrom0(&_Hero.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Hero *HeroTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Hero.Contract.SafeTransferFrom0(&_Hero.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Hero *HeroTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Hero *HeroSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Hero.Contract.SetApprovalForAll(&_Hero.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Hero *HeroTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Hero.Contract.SetApprovalForAll(&_Hero.TransactOpts, operator, approved)
}

// SetStatScienceAddress is a paid mutator transaction binding the contract method 0x5454368b.
//
// Solidity: function setStatScienceAddress(address _statScienceAddress) returns()
func (_Hero *HeroTransactor) SetStatScienceAddress(opts *bind.TransactOpts, _statScienceAddress common.Address) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "setStatScienceAddress", _statScienceAddress)
}

// SetStatScienceAddress is a paid mutator transaction binding the contract method 0x5454368b.
//
// Solidity: function setStatScienceAddress(address _statScienceAddress) returns()
func (_Hero *HeroSession) SetStatScienceAddress(_statScienceAddress common.Address) (*types.Transaction, error) {
	return _Hero.Contract.SetStatScienceAddress(&_Hero.TransactOpts, _statScienceAddress)
}

// SetStatScienceAddress is a paid mutator transaction binding the contract method 0x5454368b.
//
// Solidity: function setStatScienceAddress(address _statScienceAddress) returns()
func (_Hero *HeroTransactorSession) SetStatScienceAddress(_statScienceAddress common.Address) (*types.Transaction, error) {
	return _Hero.Contract.SetStatScienceAddress(&_Hero.TransactOpts, _statScienceAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Hero *HeroTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Hero *HeroSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hero.Contract.TransferFrom(&_Hero.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Hero *HeroTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hero.Contract.TransferFrom(&_Hero.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Hero *HeroTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Hero *HeroSession) Unpause() (*types.Transaction, error) {
	return _Hero.Contract.Unpause(&_Hero.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Hero *HeroTransactorSession) Unpause() (*types.Transaction, error) {
	return _Hero.Contract.Unpause(&_Hero.TransactOpts)
}

// UpdateHero is a paid mutator transaction binding the contract method 0xb0064103.
//
// Solidity: function updateHero((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)) _hero) returns()
func (_Hero *HeroTransactor) UpdateHero(opts *bind.TransactOpts, _hero IHeroTypesHero) (*types.Transaction, error) {
	return _Hero.contract.Transact(opts, "updateHero", _hero)
}

// UpdateHero is a paid mutator transaction binding the contract method 0xb0064103.
//
// Solidity: function updateHero((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)) _hero) returns()
func (_Hero *HeroSession) UpdateHero(_hero IHeroTypesHero) (*types.Transaction, error) {
	return _Hero.Contract.UpdateHero(&_Hero.TransactOpts, _hero)
}

// UpdateHero is a paid mutator transaction binding the contract method 0xb0064103.
//
// Solidity: function updateHero((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)) _hero) returns()
func (_Hero *HeroTransactorSession) UpdateHero(_hero IHeroTypesHero) (*types.Transaction, error) {
	return _Hero.Contract.UpdateHero(&_Hero.TransactOpts, _hero)
}

// HeroApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Hero contract.
type HeroApprovalIterator struct {
	Event *HeroApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroApproval represents a Approval event raised by the Hero contract.
type HeroApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Hero *HeroFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*HeroApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Hero.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HeroApprovalIterator{contract: _Hero.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Hero *HeroFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HeroApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Hero.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroApproval)
				if err := _Hero.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Hero *HeroFilterer) ParseApproval(log types.Log) (*HeroApproval, error) {
	event := new(HeroApproval)
	if err := _Hero.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeroApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Hero contract.
type HeroApprovalForAllIterator struct {
	Event *HeroApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroApprovalForAll represents a ApprovalForAll event raised by the Hero contract.
type HeroApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Hero *HeroFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*HeroApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Hero.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &HeroApprovalForAllIterator{contract: _Hero.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Hero *HeroFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *HeroApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Hero.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroApprovalForAll)
				if err := _Hero.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Hero *HeroFilterer) ParseApprovalForAll(log types.Log) (*HeroApprovalForAll, error) {
	event := new(HeroApprovalForAll)
	if err := _Hero.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeroHeroSummonedIterator is returned from FilterHeroSummoned and is used to iterate over the raw logs and unpacked data for HeroSummoned events raised by the Hero contract.
type HeroHeroSummonedIterator struct {
	Event *HeroHeroSummoned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroHeroSummonedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroHeroSummoned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroHeroSummoned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroHeroSummonedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroHeroSummonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroHeroSummoned represents a HeroSummoned event raised by the Hero contract.
type HeroHeroSummoned struct {
	Owner       common.Address
	HeroId      *big.Int
	SummonerId  *big.Int
	AssistantId *big.Int
	StatGenes   *big.Int
	VisualGenes *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHeroSummoned is a free log retrieval operation binding the contract event 0xef277373d709abc6be9d4926ac62b54991f4de2b6f8718079ae3d735ded22340.
//
// Solidity: event HeroSummoned(address indexed owner, uint256 heroId, uint256 summonerId, uint256 assistantId, uint256 statGenes, uint256 visualGenes)
func (_Hero *HeroFilterer) FilterHeroSummoned(opts *bind.FilterOpts, owner []common.Address) (*HeroHeroSummonedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Hero.contract.FilterLogs(opts, "HeroSummoned", ownerRule)
	if err != nil {
		return nil, err
	}
	return &HeroHeroSummonedIterator{contract: _Hero.contract, event: "HeroSummoned", logs: logs, sub: sub}, nil
}

// WatchHeroSummoned is a free log subscription operation binding the contract event 0xef277373d709abc6be9d4926ac62b54991f4de2b6f8718079ae3d735ded22340.
//
// Solidity: event HeroSummoned(address indexed owner, uint256 heroId, uint256 summonerId, uint256 assistantId, uint256 statGenes, uint256 visualGenes)
func (_Hero *HeroFilterer) WatchHeroSummoned(opts *bind.WatchOpts, sink chan<- *HeroHeroSummoned, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Hero.contract.WatchLogs(opts, "HeroSummoned", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroHeroSummoned)
				if err := _Hero.contract.UnpackLog(event, "HeroSummoned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHeroSummoned is a log parse operation binding the contract event 0xef277373d709abc6be9d4926ac62b54991f4de2b6f8718079ae3d735ded22340.
//
// Solidity: event HeroSummoned(address indexed owner, uint256 heroId, uint256 summonerId, uint256 assistantId, uint256 statGenes, uint256 visualGenes)
func (_Hero *HeroFilterer) ParseHeroSummoned(log types.Log) (*HeroHeroSummoned, error) {
	event := new(HeroHeroSummoned)
	if err := _Hero.contract.UnpackLog(event, "HeroSummoned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeroPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Hero contract.
type HeroPausedIterator struct {
	Event *HeroPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroPaused represents a Paused event raised by the Hero contract.
type HeroPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Hero *HeroFilterer) FilterPaused(opts *bind.FilterOpts) (*HeroPausedIterator, error) {

	logs, sub, err := _Hero.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &HeroPausedIterator{contract: _Hero.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Hero *HeroFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HeroPaused) (event.Subscription, error) {

	logs, sub, err := _Hero.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroPaused)
				if err := _Hero.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Hero *HeroFilterer) ParsePaused(log types.Log) (*HeroPaused, error) {
	event := new(HeroPaused)
	if err := _Hero.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeroRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Hero contract.
type HeroRoleAdminChangedIterator struct {
	Event *HeroRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroRoleAdminChanged represents a RoleAdminChanged event raised by the Hero contract.
type HeroRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Hero *HeroFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*HeroRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Hero.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &HeroRoleAdminChangedIterator{contract: _Hero.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Hero *HeroFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *HeroRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Hero.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroRoleAdminChanged)
				if err := _Hero.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Hero *HeroFilterer) ParseRoleAdminChanged(log types.Log) (*HeroRoleAdminChanged, error) {
	event := new(HeroRoleAdminChanged)
	if err := _Hero.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeroRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Hero contract.
type HeroRoleGrantedIterator struct {
	Event *HeroRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroRoleGranted represents a RoleGranted event raised by the Hero contract.
type HeroRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Hero *HeroFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HeroRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hero.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HeroRoleGrantedIterator{contract: _Hero.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Hero *HeroFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *HeroRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hero.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroRoleGranted)
				if err := _Hero.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Hero *HeroFilterer) ParseRoleGranted(log types.Log) (*HeroRoleGranted, error) {
	event := new(HeroRoleGranted)
	if err := _Hero.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeroRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Hero contract.
type HeroRoleRevokedIterator struct {
	Event *HeroRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroRoleRevoked represents a RoleRevoked event raised by the Hero contract.
type HeroRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Hero *HeroFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HeroRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hero.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HeroRoleRevokedIterator{contract: _Hero.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Hero *HeroFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *HeroRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hero.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroRoleRevoked)
				if err := _Hero.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Hero *HeroFilterer) ParseRoleRevoked(log types.Log) (*HeroRoleRevoked, error) {
	event := new(HeroRoleRevoked)
	if err := _Hero.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeroTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Hero contract.
type HeroTransferIterator struct {
	Event *HeroTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroTransfer represents a Transfer event raised by the Hero contract.
type HeroTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Hero *HeroFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*HeroTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Hero.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HeroTransferIterator{contract: _Hero.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Hero *HeroFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HeroTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Hero.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroTransfer)
				if err := _Hero.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Hero *HeroFilterer) ParseTransfer(log types.Log) (*HeroTransfer, error) {
	event := new(HeroTransfer)
	if err := _Hero.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeroUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Hero contract.
type HeroUnpausedIterator struct {
	Event *HeroUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroUnpaused represents a Unpaused event raised by the Hero contract.
type HeroUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Hero *HeroFilterer) FilterUnpaused(opts *bind.FilterOpts) (*HeroUnpausedIterator, error) {

	logs, sub, err := _Hero.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &HeroUnpausedIterator{contract: _Hero.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Hero *HeroFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HeroUnpaused) (event.Subscription, error) {

	logs, sub, err := _Hero.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroUnpaused)
				if err := _Hero.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Hero *HeroFilterer) ParseUnpaused(log types.Log) (*HeroUnpaused, error) {
	event := new(HeroUnpaused)
	if err := _Hero.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
