// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tavern

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

// TavernMetaData contains all meta data concerning the Tavern contract.
var TavernMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_heroCoreAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_geneScienceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_jewelTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gaiaTearsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_statScienceAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crystalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"heroId\",\"type\":\"uint256\"}],\"name\":\"CrystalOpen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crystalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"summonerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assistantId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"generation\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"summonerTears\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"assistantTears\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bonusItem\",\"type\":\"address\"}],\"name\":\"CrystalSummoned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionHeroCore\",\"outputs\":[{\"internalType\":\"contractIHeroCore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCooldown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseSummonFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidAmount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"summonedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextSummonTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"summonerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assistantId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"summons\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxSummons\",\"type\":\"uint32\"}],\"internalType\":\"structIHeroTypes.SummoningInfo\",\"name\":\"summoningInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"statGenes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"visualGenes\",\"type\":\"uint256\"},{\"internalType\":\"enumIHeroTypes.Rarity\",\"name\":\"rarity\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"shiny\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"generation\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"firstName\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastName\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"shinyStyle\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"class\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"subClass\",\"type\":\"uint8\"}],\"internalType\":\"structIHeroTypes.HeroInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"staminaFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hpFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mpFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"level\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"xp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"currentQuest\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sp\",\"type\":\"uint8\"},{\"internalType\":\"enumIHeroTypes.HeroStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIHeroTypes.HeroState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"stamina\",\"type\":\"uint16\"}],\"internalType\":\"structIHeroTypes.HeroStats\",\"name\":\"stats\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpLg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpLg\",\"type\":\"uint16\"}],\"internalType\":\"structIHeroTypes.HeroStatGrowth\",\"name\":\"primaryStatGrowth\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpLg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpLg\",\"type\":\"uint16\"}],\"internalType\":\"structIHeroTypes.HeroStatGrowth\",\"name\":\"secondaryStatGrowth\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"mining\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gardening\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"foraging\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"fishing\",\"type\":\"uint16\"}],\"internalType\":\"structIHeroTypes.HeroProfessions\",\"name\":\"professions\",\"type\":\"tuple\"}],\"internalType\":\"structIHeroTypes.Hero\",\"name\":\"_hero\",\"type\":\"tuple\"}],\"name\":\"calculateSummoningCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownPerGen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_startingPrice\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_endingPrice\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"_duration\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_winner\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"crystals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"summonerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assistantId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"generation\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"createdBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"heroId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"summonerTears\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"assistantTears\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bonusItem\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxSummons\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"firstName\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastName\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"shinyStyle\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rarityRoll\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rarityMod\",\"type\":\"uint256\"}],\"name\":\"determineRarity\",\"outputs\":[{\"internalType\":\"enumIHeroTypes.Rarity\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"randomNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"digits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"extractNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_crystalId\",\"type\":\"uint256\"}],\"name\":\"getCrystal\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"summonerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assistantId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"generation\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"createdBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"heroId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"summonerTears\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"assistantTears\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bonusItem\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxSummons\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"firstName\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastName\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"shinyStyle\",\"type\":\"uint8\"}],\"internalType\":\"structICrystalTypes.HeroCrystal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserAuctions\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserCrystals\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increasePerGen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increasePerSummon\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isOnAuction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jewelToken\",\"outputs\":[{\"internalType\":\"contractIJewelToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newSummonCooldown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_crystalId\",\"type\":\"uint256\"}],\"name\":\"open\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_crystalId\",\"type\":\"uint256\"}],\"name\":\"rechargeCrystal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_feeAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_feePercents\",\"type\":\"uint256[]\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_geneScienceAddress\",\"type\":\"address\"}],\"name\":\"setGeneScienceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_statScienceAddress\",\"type\":\"address\"}],\"name\":\"setStatScienceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newSummonCooldown\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseCooldown\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cooldownPerGen\",\"type\":\"uint256\"}],\"name\":\"setSummonCooldowns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseSummonFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_increasePerSummon\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_increasePerGen\",\"type\":\"uint256\"}],\"name\":\"setSummonFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"statScience\",\"outputs\":[{\"internalType\":\"contractIStatScience\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summonerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assistantId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_summonerTears\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_assistantTears\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_bonusItem\",\"type\":\"address\"}],\"name\":\"summonCrystal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userAuctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userCrystals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"vrf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TavernABI is the input ABI used to generate the binding from.
// Deprecated: Use TavernMetaData.ABI instead.
var TavernABI = TavernMetaData.ABI

// Tavern is an auto generated Go binding around an Ethereum contract.
type Tavern struct {
	TavernCaller     // Read-only binding to the contract
	TavernTransactor // Write-only binding to the contract
	TavernFilterer   // Log filterer for contract events
}

// TavernCaller is an auto generated read-only Go binding around an Ethereum contract.
type TavernCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TavernTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TavernTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TavernFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TavernFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TavernSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TavernSession struct {
	Contract     *Tavern           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TavernCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TavernCallerSession struct {
	Contract *TavernCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TavernTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TavernTransactorSession struct {
	Contract     *TavernTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TavernRaw is an auto generated low-level Go binding around an Ethereum contract.
type TavernRaw struct {
	Contract *Tavern // Generic contract binding to access the raw methods on
}

// TavernCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TavernCallerRaw struct {
	Contract *TavernCaller // Generic read-only contract binding to access the raw methods on
}

// TavernTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TavernTransactorRaw struct {
	Contract *TavernTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTavern creates a new instance of Tavern, bound to a specific deployed contract.
func NewTavern(address common.Address, backend bind.ContractBackend) (*Tavern, error) {
	contract, err := bindTavern(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tavern{TavernCaller: TavernCaller{contract: contract}, TavernTransactor: TavernTransactor{contract: contract}, TavernFilterer: TavernFilterer{contract: contract}}, nil
}

// NewTavernCaller creates a new read-only instance of Tavern, bound to a specific deployed contract.
func NewTavernCaller(address common.Address, caller bind.ContractCaller) (*TavernCaller, error) {
	contract, err := bindTavern(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TavernCaller{contract: contract}, nil
}

// NewTavernTransactor creates a new write-only instance of Tavern, bound to a specific deployed contract.
func NewTavernTransactor(address common.Address, transactor bind.ContractTransactor) (*TavernTransactor, error) {
	contract, err := bindTavern(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TavernTransactor{contract: contract}, nil
}

// NewTavernFilterer creates a new log filterer instance of Tavern, bound to a specific deployed contract.
func NewTavernFilterer(address common.Address, filterer bind.ContractFilterer) (*TavernFilterer, error) {
	contract, err := bindTavern(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TavernFilterer{contract: contract}, nil
}

// bindTavern binds a generic wrapper to an already deployed contract.
func bindTavern(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TavernABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tavern *TavernRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tavern.Contract.TavernCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tavern *TavernRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tavern.Contract.TavernTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tavern *TavernRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tavern.Contract.TavernTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tavern *TavernCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tavern.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tavern *TavernTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tavern.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tavern *TavernTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tavern.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Tavern *TavernCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Tavern *TavernSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Tavern.Contract.DEFAULTADMINROLE(&_Tavern.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Tavern *TavernCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Tavern.Contract.DEFAULTADMINROLE(&_Tavern.CallOpts)
}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_Tavern *TavernCaller) MODERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "MODERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_Tavern *TavernSession) MODERATORROLE() ([32]byte, error) {
	return _Tavern.Contract.MODERATORROLE(&_Tavern.CallOpts)
}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_Tavern *TavernCallerSession) MODERATORROLE() ([32]byte, error) {
	return _Tavern.Contract.MODERATORROLE(&_Tavern.CallOpts)
}

// AuctionHeroCore is a free data retrieval call binding the contract method 0xe03825b7.
//
// Solidity: function auctionHeroCore() view returns(address)
func (_Tavern *TavernCaller) AuctionHeroCore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "auctionHeroCore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuctionHeroCore is a free data retrieval call binding the contract method 0xe03825b7.
//
// Solidity: function auctionHeroCore() view returns(address)
func (_Tavern *TavernSession) AuctionHeroCore() (common.Address, error) {
	return _Tavern.Contract.AuctionHeroCore(&_Tavern.CallOpts)
}

// AuctionHeroCore is a free data retrieval call binding the contract method 0xe03825b7.
//
// Solidity: function auctionHeroCore() view returns(address)
func (_Tavern *TavernCallerSession) AuctionHeroCore() (common.Address, error) {
	return _Tavern.Contract.AuctionHeroCore(&_Tavern.CallOpts)
}

// BaseCooldown is a free data retrieval call binding the contract method 0x93deec27.
//
// Solidity: function baseCooldown() view returns(uint256)
func (_Tavern *TavernCaller) BaseCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "baseCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseCooldown is a free data retrieval call binding the contract method 0x93deec27.
//
// Solidity: function baseCooldown() view returns(uint256)
func (_Tavern *TavernSession) BaseCooldown() (*big.Int, error) {
	return _Tavern.Contract.BaseCooldown(&_Tavern.CallOpts)
}

// BaseCooldown is a free data retrieval call binding the contract method 0x93deec27.
//
// Solidity: function baseCooldown() view returns(uint256)
func (_Tavern *TavernCallerSession) BaseCooldown() (*big.Int, error) {
	return _Tavern.Contract.BaseCooldown(&_Tavern.CallOpts)
}

// BaseSummonFee is a free data retrieval call binding the contract method 0x45d31394.
//
// Solidity: function baseSummonFee() view returns(uint256)
func (_Tavern *TavernCaller) BaseSummonFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "baseSummonFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseSummonFee is a free data retrieval call binding the contract method 0x45d31394.
//
// Solidity: function baseSummonFee() view returns(uint256)
func (_Tavern *TavernSession) BaseSummonFee() (*big.Int, error) {
	return _Tavern.Contract.BaseSummonFee(&_Tavern.CallOpts)
}

// BaseSummonFee is a free data retrieval call binding the contract method 0x45d31394.
//
// Solidity: function baseSummonFee() view returns(uint256)
func (_Tavern *TavernCallerSession) BaseSummonFee() (*big.Int, error) {
	return _Tavern.Contract.BaseSummonFee(&_Tavern.CallOpts)
}

// CalculateSummoningCost is a free data retrieval call binding the contract method 0x98d25152.
//
// Solidity: function calculateSummoningCost((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)) _hero) view returns(uint256)
func (_Tavern *TavernCaller) CalculateSummoningCost(opts *bind.CallOpts, _hero IHeroTypesHero) (*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "calculateSummoningCost", _hero)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSummoningCost is a free data retrieval call binding the contract method 0x98d25152.
//
// Solidity: function calculateSummoningCost((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)) _hero) view returns(uint256)
func (_Tavern *TavernSession) CalculateSummoningCost(_hero IHeroTypesHero) (*big.Int, error) {
	return _Tavern.Contract.CalculateSummoningCost(&_Tavern.CallOpts, _hero)
}

// CalculateSummoningCost is a free data retrieval call binding the contract method 0x98d25152.
//
// Solidity: function calculateSummoningCost((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)) _hero) view returns(uint256)
func (_Tavern *TavernCallerSession) CalculateSummoningCost(_hero IHeroTypesHero) (*big.Int, error) {
	return _Tavern.Contract.CalculateSummoningCost(&_Tavern.CallOpts, _hero)
}

// CooldownPerGen is a free data retrieval call binding the contract method 0xa8029920.
//
// Solidity: function cooldownPerGen() view returns(uint256)
func (_Tavern *TavernCaller) CooldownPerGen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "cooldownPerGen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownPerGen is a free data retrieval call binding the contract method 0xa8029920.
//
// Solidity: function cooldownPerGen() view returns(uint256)
func (_Tavern *TavernSession) CooldownPerGen() (*big.Int, error) {
	return _Tavern.Contract.CooldownPerGen(&_Tavern.CallOpts)
}

// CooldownPerGen is a free data retrieval call binding the contract method 0xa8029920.
//
// Solidity: function cooldownPerGen() view returns(uint256)
func (_Tavern *TavernCallerSession) CooldownPerGen() (*big.Int, error) {
	return _Tavern.Contract.CooldownPerGen(&_Tavern.CallOpts)
}

// Crystals is a free data retrieval call binding the contract method 0x31b3ab3f.
//
// Solidity: function crystals(uint256 ) view returns(address owner, uint256 summonerId, uint256 assistantId, uint16 generation, uint256 createdBlock, uint256 heroId, uint8 summonerTears, uint8 assistantTears, address bonusItem, uint32 maxSummons, uint32 firstName, uint32 lastName, uint8 shinyStyle)
func (_Tavern *TavernCaller) Crystals(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
}, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "crystals", arg0)

	outstruct := new(struct {
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
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SummonerId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AssistantId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Generation = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.CreatedBlock = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HeroId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.SummonerTears = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.AssistantTears = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.BonusItem = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.MaxSummons = *abi.ConvertType(out[9], new(uint32)).(*uint32)
	outstruct.FirstName = *abi.ConvertType(out[10], new(uint32)).(*uint32)
	outstruct.LastName = *abi.ConvertType(out[11], new(uint32)).(*uint32)
	outstruct.ShinyStyle = *abi.ConvertType(out[12], new(uint8)).(*uint8)

	return *outstruct, err

}

// Crystals is a free data retrieval call binding the contract method 0x31b3ab3f.
//
// Solidity: function crystals(uint256 ) view returns(address owner, uint256 summonerId, uint256 assistantId, uint16 generation, uint256 createdBlock, uint256 heroId, uint8 summonerTears, uint8 assistantTears, address bonusItem, uint32 maxSummons, uint32 firstName, uint32 lastName, uint8 shinyStyle)
func (_Tavern *TavernSession) Crystals(arg0 *big.Int) (struct {
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
}, error) {
	return _Tavern.Contract.Crystals(&_Tavern.CallOpts, arg0)
}

// Crystals is a free data retrieval call binding the contract method 0x31b3ab3f.
//
// Solidity: function crystals(uint256 ) view returns(address owner, uint256 summonerId, uint256 assistantId, uint16 generation, uint256 createdBlock, uint256 heroId, uint8 summonerTears, uint8 assistantTears, address bonusItem, uint32 maxSummons, uint32 firstName, uint32 lastName, uint8 shinyStyle)
func (_Tavern *TavernCallerSession) Crystals(arg0 *big.Int) (struct {
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
}, error) {
	return _Tavern.Contract.Crystals(&_Tavern.CallOpts, arg0)
}

// DetermineRarity is a free data retrieval call binding the contract method 0x9b4178c1.
//
// Solidity: function determineRarity(uint256 _rarityRoll, uint256 _rarityMod) pure returns(uint8)
func (_Tavern *TavernCaller) DetermineRarity(opts *bind.CallOpts, _rarityRoll *big.Int, _rarityMod *big.Int) (uint8, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "determineRarity", _rarityRoll, _rarityMod)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DetermineRarity is a free data retrieval call binding the contract method 0x9b4178c1.
//
// Solidity: function determineRarity(uint256 _rarityRoll, uint256 _rarityMod) pure returns(uint8)
func (_Tavern *TavernSession) DetermineRarity(_rarityRoll *big.Int, _rarityMod *big.Int) (uint8, error) {
	return _Tavern.Contract.DetermineRarity(&_Tavern.CallOpts, _rarityRoll, _rarityMod)
}

// DetermineRarity is a free data retrieval call binding the contract method 0x9b4178c1.
//
// Solidity: function determineRarity(uint256 _rarityRoll, uint256 _rarityMod) pure returns(uint8)
func (_Tavern *TavernCallerSession) DetermineRarity(_rarityRoll *big.Int, _rarityMod *big.Int) (uint8, error) {
	return _Tavern.Contract.DetermineRarity(&_Tavern.CallOpts, _rarityRoll, _rarityMod)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() view returns(bool)
func (_Tavern *TavernCaller) Enabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "enabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() view returns(bool)
func (_Tavern *TavernSession) Enabled() (bool, error) {
	return _Tavern.Contract.Enabled(&_Tavern.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() view returns(bool)
func (_Tavern *TavernCallerSession) Enabled() (bool, error) {
	return _Tavern.Contract.Enabled(&_Tavern.CallOpts)
}

// ExtractNumber is a free data retrieval call binding the contract method 0xb00a7b9e.
//
// Solidity: function extractNumber(uint256 randomNumber, uint256 digits, uint256 offset) pure returns(uint256 result)
func (_Tavern *TavernCaller) ExtractNumber(opts *bind.CallOpts, randomNumber *big.Int, digits *big.Int, offset *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "extractNumber", randomNumber, digits, offset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtractNumber is a free data retrieval call binding the contract method 0xb00a7b9e.
//
// Solidity: function extractNumber(uint256 randomNumber, uint256 digits, uint256 offset) pure returns(uint256 result)
func (_Tavern *TavernSession) ExtractNumber(randomNumber *big.Int, digits *big.Int, offset *big.Int) (*big.Int, error) {
	return _Tavern.Contract.ExtractNumber(&_Tavern.CallOpts, randomNumber, digits, offset)
}

// ExtractNumber is a free data retrieval call binding the contract method 0xb00a7b9e.
//
// Solidity: function extractNumber(uint256 randomNumber, uint256 digits, uint256 offset) pure returns(uint256 result)
func (_Tavern *TavernCallerSession) ExtractNumber(randomNumber *big.Int, digits *big.Int, offset *big.Int) (*big.Int, error) {
	return _Tavern.Contract.ExtractNumber(&_Tavern.CallOpts, randomNumber, digits, offset)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(uint256 auctionId, address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_Tavern *TavernCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	AuctionId     *big.Int
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "getAuction", _tokenId)

	outstruct := new(struct {
		AuctionId     *big.Int
		Seller        common.Address
		StartingPrice *big.Int
		EndingPrice   *big.Int
		Duration      *big.Int
		StartedAt     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AuctionId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StartingPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndingPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(uint256 auctionId, address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_Tavern *TavernSession) GetAuction(_tokenId *big.Int) (struct {
	AuctionId     *big.Int
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _Tavern.Contract.GetAuction(&_Tavern.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(uint256 auctionId, address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_Tavern *TavernCallerSession) GetAuction(_tokenId *big.Int) (struct {
	AuctionId     *big.Int
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _Tavern.Contract.GetAuction(&_Tavern.CallOpts, _tokenId)
}

// GetCrystal is a free data retrieval call binding the contract method 0x9c2f7a43.
//
// Solidity: function getCrystal(uint256 _crystalId) view returns((address,uint256,uint256,uint16,uint256,uint256,uint8,uint8,address,uint32,uint32,uint32,uint8))
func (_Tavern *TavernCaller) GetCrystal(opts *bind.CallOpts, _crystalId *big.Int) (ICrystalTypesHeroCrystal, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "getCrystal", _crystalId)

	if err != nil {
		return *new(ICrystalTypesHeroCrystal), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrystalTypesHeroCrystal)).(*ICrystalTypesHeroCrystal)

	return out0, err

}

// GetCrystal is a free data retrieval call binding the contract method 0x9c2f7a43.
//
// Solidity: function getCrystal(uint256 _crystalId) view returns((address,uint256,uint256,uint16,uint256,uint256,uint8,uint8,address,uint32,uint32,uint32,uint8))
func (_Tavern *TavernSession) GetCrystal(_crystalId *big.Int) (ICrystalTypesHeroCrystal, error) {
	return _Tavern.Contract.GetCrystal(&_Tavern.CallOpts, _crystalId)
}

// GetCrystal is a free data retrieval call binding the contract method 0x9c2f7a43.
//
// Solidity: function getCrystal(uint256 _crystalId) view returns((address,uint256,uint256,uint16,uint256,uint256,uint8,uint8,address,uint32,uint32,uint32,uint8))
func (_Tavern *TavernCallerSession) GetCrystal(_crystalId *big.Int) (ICrystalTypesHeroCrystal, error) {
	return _Tavern.Contract.GetCrystal(&_Tavern.CallOpts, _crystalId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_Tavern *TavernCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "getCurrentPrice", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_Tavern *TavernSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _Tavern.Contract.GetCurrentPrice(&_Tavern.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_Tavern *TavernCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _Tavern.Contract.GetCurrentPrice(&_Tavern.CallOpts, _tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Tavern *TavernCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Tavern *TavernSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Tavern.Contract.GetRoleAdmin(&_Tavern.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Tavern *TavernCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Tavern.Contract.GetRoleAdmin(&_Tavern.CallOpts, role)
}

// GetUserAuctions is a free data retrieval call binding the contract method 0xff3ad0b4.
//
// Solidity: function getUserAuctions(address _address) view returns(uint256[])
func (_Tavern *TavernCaller) GetUserAuctions(opts *bind.CallOpts, _address common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "getUserAuctions", _address)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserAuctions is a free data retrieval call binding the contract method 0xff3ad0b4.
//
// Solidity: function getUserAuctions(address _address) view returns(uint256[])
func (_Tavern *TavernSession) GetUserAuctions(_address common.Address) ([]*big.Int, error) {
	return _Tavern.Contract.GetUserAuctions(&_Tavern.CallOpts, _address)
}

// GetUserAuctions is a free data retrieval call binding the contract method 0xff3ad0b4.
//
// Solidity: function getUserAuctions(address _address) view returns(uint256[])
func (_Tavern *TavernCallerSession) GetUserAuctions(_address common.Address) ([]*big.Int, error) {
	return _Tavern.Contract.GetUserAuctions(&_Tavern.CallOpts, _address)
}

// GetUserCrystals is a free data retrieval call binding the contract method 0xbb987007.
//
// Solidity: function getUserCrystals(address _address) view returns(uint256[])
func (_Tavern *TavernCaller) GetUserCrystals(opts *bind.CallOpts, _address common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "getUserCrystals", _address)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserCrystals is a free data retrieval call binding the contract method 0xbb987007.
//
// Solidity: function getUserCrystals(address _address) view returns(uint256[])
func (_Tavern *TavernSession) GetUserCrystals(_address common.Address) ([]*big.Int, error) {
	return _Tavern.Contract.GetUserCrystals(&_Tavern.CallOpts, _address)
}

// GetUserCrystals is a free data retrieval call binding the contract method 0xbb987007.
//
// Solidity: function getUserCrystals(address _address) view returns(uint256[])
func (_Tavern *TavernCallerSession) GetUserCrystals(_address common.Address) ([]*big.Int, error) {
	return _Tavern.Contract.GetUserCrystals(&_Tavern.CallOpts, _address)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Tavern *TavernCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Tavern *TavernSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Tavern.Contract.HasRole(&_Tavern.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Tavern *TavernCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Tavern.Contract.HasRole(&_Tavern.CallOpts, role, account)
}

// IncreasePerGen is a free data retrieval call binding the contract method 0x20c20a07.
//
// Solidity: function increasePerGen() view returns(uint256)
func (_Tavern *TavernCaller) IncreasePerGen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "increasePerGen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreasePerGen is a free data retrieval call binding the contract method 0x20c20a07.
//
// Solidity: function increasePerGen() view returns(uint256)
func (_Tavern *TavernSession) IncreasePerGen() (*big.Int, error) {
	return _Tavern.Contract.IncreasePerGen(&_Tavern.CallOpts)
}

// IncreasePerGen is a free data retrieval call binding the contract method 0x20c20a07.
//
// Solidity: function increasePerGen() view returns(uint256)
func (_Tavern *TavernCallerSession) IncreasePerGen() (*big.Int, error) {
	return _Tavern.Contract.IncreasePerGen(&_Tavern.CallOpts)
}

// IncreasePerSummon is a free data retrieval call binding the contract method 0xe9dea449.
//
// Solidity: function increasePerSummon() view returns(uint256)
func (_Tavern *TavernCaller) IncreasePerSummon(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "increasePerSummon")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreasePerSummon is a free data retrieval call binding the contract method 0xe9dea449.
//
// Solidity: function increasePerSummon() view returns(uint256)
func (_Tavern *TavernSession) IncreasePerSummon() (*big.Int, error) {
	return _Tavern.Contract.IncreasePerSummon(&_Tavern.CallOpts)
}

// IncreasePerSummon is a free data retrieval call binding the contract method 0xe9dea449.
//
// Solidity: function increasePerSummon() view returns(uint256)
func (_Tavern *TavernCallerSession) IncreasePerSummon() (*big.Int, error) {
	return _Tavern.Contract.IncreasePerSummon(&_Tavern.CallOpts)
}

// JewelToken is a free data retrieval call binding the contract method 0x237cc032.
//
// Solidity: function jewelToken() view returns(address)
func (_Tavern *TavernCaller) JewelToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "jewelToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JewelToken is a free data retrieval call binding the contract method 0x237cc032.
//
// Solidity: function jewelToken() view returns(address)
func (_Tavern *TavernSession) JewelToken() (common.Address, error) {
	return _Tavern.Contract.JewelToken(&_Tavern.CallOpts)
}

// JewelToken is a free data retrieval call binding the contract method 0x237cc032.
//
// Solidity: function jewelToken() view returns(address)
func (_Tavern *TavernCallerSession) JewelToken() (common.Address, error) {
	return _Tavern.Contract.JewelToken(&_Tavern.CallOpts)
}

// NewSummonCooldown is a free data retrieval call binding the contract method 0x2c0c0177.
//
// Solidity: function newSummonCooldown() view returns(uint256)
func (_Tavern *TavernCaller) NewSummonCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "newSummonCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewSummonCooldown is a free data retrieval call binding the contract method 0x2c0c0177.
//
// Solidity: function newSummonCooldown() view returns(uint256)
func (_Tavern *TavernSession) NewSummonCooldown() (*big.Int, error) {
	return _Tavern.Contract.NewSummonCooldown(&_Tavern.CallOpts)
}

// NewSummonCooldown is a free data retrieval call binding the contract method 0x2c0c0177.
//
// Solidity: function newSummonCooldown() view returns(uint256)
func (_Tavern *TavernCallerSession) NewSummonCooldown() (*big.Int, error) {
	return _Tavern.Contract.NewSummonCooldown(&_Tavern.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tavern *TavernCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tavern *TavernSession) Owner() (common.Address, error) {
	return _Tavern.Contract.Owner(&_Tavern.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tavern *TavernCallerSession) Owner() (common.Address, error) {
	return _Tavern.Contract.Owner(&_Tavern.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_Tavern *TavernCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "ownerCut")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_Tavern *TavernSession) OwnerCut() (*big.Int, error) {
	return _Tavern.Contract.OwnerCut(&_Tavern.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_Tavern *TavernCallerSession) OwnerCut() (*big.Int, error) {
	return _Tavern.Contract.OwnerCut(&_Tavern.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tavern *TavernCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tavern *TavernSession) Paused() (bool, error) {
	return _Tavern.Contract.Paused(&_Tavern.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tavern *TavernCallerSession) Paused() (bool, error) {
	return _Tavern.Contract.Paused(&_Tavern.CallOpts)
}

// StatScience is a free data retrieval call binding the contract method 0x466901c3.
//
// Solidity: function statScience() view returns(address)
func (_Tavern *TavernCaller) StatScience(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "statScience")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StatScience is a free data retrieval call binding the contract method 0x466901c3.
//
// Solidity: function statScience() view returns(address)
func (_Tavern *TavernSession) StatScience() (common.Address, error) {
	return _Tavern.Contract.StatScience(&_Tavern.CallOpts)
}

// StatScience is a free data retrieval call binding the contract method 0x466901c3.
//
// Solidity: function statScience() view returns(address)
func (_Tavern *TavernCallerSession) StatScience() (common.Address, error) {
	return _Tavern.Contract.StatScience(&_Tavern.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Tavern *TavernCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Tavern *TavernSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Tavern.Contract.SupportsInterface(&_Tavern.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Tavern *TavernCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Tavern.Contract.SupportsInterface(&_Tavern.CallOpts, interfaceId)
}

// UserAuctions is a free data retrieval call binding the contract method 0xb4fbe80a.
//
// Solidity: function userAuctions(address , uint256 ) view returns(uint256)
func (_Tavern *TavernCaller) UserAuctions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "userAuctions", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserAuctions is a free data retrieval call binding the contract method 0xb4fbe80a.
//
// Solidity: function userAuctions(address , uint256 ) view returns(uint256)
func (_Tavern *TavernSession) UserAuctions(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Tavern.Contract.UserAuctions(&_Tavern.CallOpts, arg0, arg1)
}

// UserAuctions is a free data retrieval call binding the contract method 0xb4fbe80a.
//
// Solidity: function userAuctions(address , uint256 ) view returns(uint256)
func (_Tavern *TavernCallerSession) UserAuctions(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Tavern.Contract.UserAuctions(&_Tavern.CallOpts, arg0, arg1)
}

// UserCrystals is a free data retrieval call binding the contract method 0x90c757eb.
//
// Solidity: function userCrystals(address , uint256 ) view returns(uint256)
func (_Tavern *TavernCaller) UserCrystals(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "userCrystals", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserCrystals is a free data retrieval call binding the contract method 0x90c757eb.
//
// Solidity: function userCrystals(address , uint256 ) view returns(uint256)
func (_Tavern *TavernSession) UserCrystals(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Tavern.Contract.UserCrystals(&_Tavern.CallOpts, arg0, arg1)
}

// UserCrystals is a free data retrieval call binding the contract method 0x90c757eb.
//
// Solidity: function userCrystals(address , uint256 ) view returns(uint256)
func (_Tavern *TavernCallerSession) UserCrystals(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Tavern.Contract.UserCrystals(&_Tavern.CallOpts, arg0, arg1)
}

// Vrf is a free data retrieval call binding the contract method 0x4b757f98.
//
// Solidity: function vrf(uint256 blockNumber) view returns(bytes32 result)
func (_Tavern *TavernCaller) Vrf(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Tavern.contract.Call(opts, &out, "vrf", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Vrf is a free data retrieval call binding the contract method 0x4b757f98.
//
// Solidity: function vrf(uint256 blockNumber) view returns(bytes32 result)
func (_Tavern *TavernSession) Vrf(blockNumber *big.Int) ([32]byte, error) {
	return _Tavern.Contract.Vrf(&_Tavern.CallOpts, blockNumber)
}

// Vrf is a free data retrieval call binding the contract method 0x4b757f98.
//
// Solidity: function vrf(uint256 blockNumber) view returns(bytes32 result)
func (_Tavern *TavernCallerSession) Vrf(blockNumber *big.Int) ([32]byte, error) {
	return _Tavern.Contract.Vrf(&_Tavern.CallOpts, blockNumber)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 _tokenId, uint256 _bidAmount) returns()
func (_Tavern *TavernTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "bid", _tokenId, _bidAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 _tokenId, uint256 _bidAmount) returns()
func (_Tavern *TavernSession) Bid(_tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.Bid(&_Tavern.TransactOpts, _tokenId, _bidAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 _tokenId, uint256 _bidAmount) returns()
func (_Tavern *TavernTransactorSession) Bid(_tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.Bid(&_Tavern.TransactOpts, _tokenId, _bidAmount)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_Tavern *TavernTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_Tavern *TavernSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.CancelAuction(&_Tavern.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_Tavern *TavernTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.CancelAuction(&_Tavern.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_Tavern *TavernTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_Tavern *TavernSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.CancelAuctionWhenPaused(&_Tavern.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_Tavern *TavernTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.CancelAuctionWhenPaused(&_Tavern.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x4ee42914.
//
// Solidity: function createAuction(uint256 _tokenId, uint128 _startingPrice, uint128 _endingPrice, uint64 _duration, address _winner) returns()
func (_Tavern *TavernTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration uint64, _winner common.Address) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration, _winner)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x4ee42914.
//
// Solidity: function createAuction(uint256 _tokenId, uint128 _startingPrice, uint128 _endingPrice, uint64 _duration, address _winner) returns()
func (_Tavern *TavernSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration uint64, _winner common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.CreateAuction(&_Tavern.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _winner)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x4ee42914.
//
// Solidity: function createAuction(uint256 _tokenId, uint128 _startingPrice, uint128 _endingPrice, uint64 _duration, address _winner) returns()
func (_Tavern *TavernTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration uint64, _winner common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.CreateAuction(&_Tavern.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _winner)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Tavern *TavernTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Tavern *TavernSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.GrantRole(&_Tavern.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Tavern *TavernTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.GrantRole(&_Tavern.TransactOpts, role, account)
}

// IsOnAuction is a paid mutator transaction binding the contract method 0x37e246ad.
//
// Solidity: function isOnAuction(uint256 _tokenId) returns(bool)
func (_Tavern *TavernTransactor) IsOnAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "isOnAuction", _tokenId)
}

// IsOnAuction is a paid mutator transaction binding the contract method 0x37e246ad.
//
// Solidity: function isOnAuction(uint256 _tokenId) returns(bool)
func (_Tavern *TavernSession) IsOnAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.IsOnAuction(&_Tavern.TransactOpts, _tokenId)
}

// IsOnAuction is a paid mutator transaction binding the contract method 0x37e246ad.
//
// Solidity: function isOnAuction(uint256 _tokenId) returns(bool)
func (_Tavern *TavernTransactorSession) IsOnAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.IsOnAuction(&_Tavern.TransactOpts, _tokenId)
}

// Open is a paid mutator transaction binding the contract method 0x690e7c09.
//
// Solidity: function open(uint256 _crystalId) returns(uint256)
func (_Tavern *TavernTransactor) Open(opts *bind.TransactOpts, _crystalId *big.Int) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "open", _crystalId)
}

// Open is a paid mutator transaction binding the contract method 0x690e7c09.
//
// Solidity: function open(uint256 _crystalId) returns(uint256)
func (_Tavern *TavernSession) Open(_crystalId *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.Open(&_Tavern.TransactOpts, _crystalId)
}

// Open is a paid mutator transaction binding the contract method 0x690e7c09.
//
// Solidity: function open(uint256 _crystalId) returns(uint256)
func (_Tavern *TavernTransactorSession) Open(_crystalId *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.Open(&_Tavern.TransactOpts, _crystalId)
}

// RechargeCrystal is a paid mutator transaction binding the contract method 0x5d444f91.
//
// Solidity: function rechargeCrystal(uint256 _crystalId) returns()
func (_Tavern *TavernTransactor) RechargeCrystal(opts *bind.TransactOpts, _crystalId *big.Int) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "rechargeCrystal", _crystalId)
}

// RechargeCrystal is a paid mutator transaction binding the contract method 0x5d444f91.
//
// Solidity: function rechargeCrystal(uint256 _crystalId) returns()
func (_Tavern *TavernSession) RechargeCrystal(_crystalId *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.RechargeCrystal(&_Tavern.TransactOpts, _crystalId)
}

// RechargeCrystal is a paid mutator transaction binding the contract method 0x5d444f91.
//
// Solidity: function rechargeCrystal(uint256 _crystalId) returns()
func (_Tavern *TavernTransactorSession) RechargeCrystal(_crystalId *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.RechargeCrystal(&_Tavern.TransactOpts, _crystalId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tavern *TavernTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tavern *TavernSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tavern.Contract.RenounceOwnership(&_Tavern.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tavern *TavernTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tavern.Contract.RenounceOwnership(&_Tavern.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Tavern *TavernTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Tavern *TavernSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.RenounceRole(&_Tavern.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Tavern *TavernTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.RenounceRole(&_Tavern.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Tavern *TavernTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Tavern *TavernSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.RevokeRole(&_Tavern.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Tavern *TavernTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.RevokeRole(&_Tavern.TransactOpts, role, account)
}

// SetFees is a paid mutator transaction binding the contract method 0x16f81524.
//
// Solidity: function setFees(address[] _feeAddresses, uint256[] _feePercents) returns()
func (_Tavern *TavernTransactor) SetFees(opts *bind.TransactOpts, _feeAddresses []common.Address, _feePercents []*big.Int) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "setFees", _feeAddresses, _feePercents)
}

// SetFees is a paid mutator transaction binding the contract method 0x16f81524.
//
// Solidity: function setFees(address[] _feeAddresses, uint256[] _feePercents) returns()
func (_Tavern *TavernSession) SetFees(_feeAddresses []common.Address, _feePercents []*big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.SetFees(&_Tavern.TransactOpts, _feeAddresses, _feePercents)
}

// SetFees is a paid mutator transaction binding the contract method 0x16f81524.
//
// Solidity: function setFees(address[] _feeAddresses, uint256[] _feePercents) returns()
func (_Tavern *TavernTransactorSession) SetFees(_feeAddresses []common.Address, _feePercents []*big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.SetFees(&_Tavern.TransactOpts, _feeAddresses, _feePercents)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _geneScienceAddress) returns()
func (_Tavern *TavernTransactor) SetGeneScienceAddress(opts *bind.TransactOpts, _geneScienceAddress common.Address) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "setGeneScienceAddress", _geneScienceAddress)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _geneScienceAddress) returns()
func (_Tavern *TavernSession) SetGeneScienceAddress(_geneScienceAddress common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.SetGeneScienceAddress(&_Tavern.TransactOpts, _geneScienceAddress)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _geneScienceAddress) returns()
func (_Tavern *TavernTransactorSession) SetGeneScienceAddress(_geneScienceAddress common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.SetGeneScienceAddress(&_Tavern.TransactOpts, _geneScienceAddress)
}

// SetStatScienceAddress is a paid mutator transaction binding the contract method 0x5454368b.
//
// Solidity: function setStatScienceAddress(address _statScienceAddress) returns()
func (_Tavern *TavernTransactor) SetStatScienceAddress(opts *bind.TransactOpts, _statScienceAddress common.Address) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "setStatScienceAddress", _statScienceAddress)
}

// SetStatScienceAddress is a paid mutator transaction binding the contract method 0x5454368b.
//
// Solidity: function setStatScienceAddress(address _statScienceAddress) returns()
func (_Tavern *TavernSession) SetStatScienceAddress(_statScienceAddress common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.SetStatScienceAddress(&_Tavern.TransactOpts, _statScienceAddress)
}

// SetStatScienceAddress is a paid mutator transaction binding the contract method 0x5454368b.
//
// Solidity: function setStatScienceAddress(address _statScienceAddress) returns()
func (_Tavern *TavernTransactorSession) SetStatScienceAddress(_statScienceAddress common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.SetStatScienceAddress(&_Tavern.TransactOpts, _statScienceAddress)
}

// SetSummonCooldowns is a paid mutator transaction binding the contract method 0x4e970324.
//
// Solidity: function setSummonCooldowns(uint256 _newSummonCooldown, uint256 _baseCooldown, uint256 _cooldownPerGen) returns()
func (_Tavern *TavernTransactor) SetSummonCooldowns(opts *bind.TransactOpts, _newSummonCooldown *big.Int, _baseCooldown *big.Int, _cooldownPerGen *big.Int) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "setSummonCooldowns", _newSummonCooldown, _baseCooldown, _cooldownPerGen)
}

// SetSummonCooldowns is a paid mutator transaction binding the contract method 0x4e970324.
//
// Solidity: function setSummonCooldowns(uint256 _newSummonCooldown, uint256 _baseCooldown, uint256 _cooldownPerGen) returns()
func (_Tavern *TavernSession) SetSummonCooldowns(_newSummonCooldown *big.Int, _baseCooldown *big.Int, _cooldownPerGen *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.SetSummonCooldowns(&_Tavern.TransactOpts, _newSummonCooldown, _baseCooldown, _cooldownPerGen)
}

// SetSummonCooldowns is a paid mutator transaction binding the contract method 0x4e970324.
//
// Solidity: function setSummonCooldowns(uint256 _newSummonCooldown, uint256 _baseCooldown, uint256 _cooldownPerGen) returns()
func (_Tavern *TavernTransactorSession) SetSummonCooldowns(_newSummonCooldown *big.Int, _baseCooldown *big.Int, _cooldownPerGen *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.SetSummonCooldowns(&_Tavern.TransactOpts, _newSummonCooldown, _baseCooldown, _cooldownPerGen)
}

// SetSummonFees is a paid mutator transaction binding the contract method 0x03466cfd.
//
// Solidity: function setSummonFees(uint256 _baseSummonFee, uint256 _increasePerSummon, uint256 _increasePerGen) returns()
func (_Tavern *TavernTransactor) SetSummonFees(opts *bind.TransactOpts, _baseSummonFee *big.Int, _increasePerSummon *big.Int, _increasePerGen *big.Int) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "setSummonFees", _baseSummonFee, _increasePerSummon, _increasePerGen)
}

// SetSummonFees is a paid mutator transaction binding the contract method 0x03466cfd.
//
// Solidity: function setSummonFees(uint256 _baseSummonFee, uint256 _increasePerSummon, uint256 _increasePerGen) returns()
func (_Tavern *TavernSession) SetSummonFees(_baseSummonFee *big.Int, _increasePerSummon *big.Int, _increasePerGen *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.SetSummonFees(&_Tavern.TransactOpts, _baseSummonFee, _increasePerSummon, _increasePerGen)
}

// SetSummonFees is a paid mutator transaction binding the contract method 0x03466cfd.
//
// Solidity: function setSummonFees(uint256 _baseSummonFee, uint256 _increasePerSummon, uint256 _increasePerGen) returns()
func (_Tavern *TavernTransactorSession) SetSummonFees(_baseSummonFee *big.Int, _increasePerSummon *big.Int, _increasePerGen *big.Int) (*types.Transaction, error) {
	return _Tavern.Contract.SetSummonFees(&_Tavern.TransactOpts, _baseSummonFee, _increasePerSummon, _increasePerGen)
}

// SummonCrystal is a paid mutator transaction binding the contract method 0x4ea8a311.
//
// Solidity: function summonCrystal(uint256 _summonerId, uint256 _assistantId, uint16 _summonerTears, uint16 _assistantTears, address _bonusItem) returns()
func (_Tavern *TavernTransactor) SummonCrystal(opts *bind.TransactOpts, _summonerId *big.Int, _assistantId *big.Int, _summonerTears uint16, _assistantTears uint16, _bonusItem common.Address) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "summonCrystal", _summonerId, _assistantId, _summonerTears, _assistantTears, _bonusItem)
}

// SummonCrystal is a paid mutator transaction binding the contract method 0x4ea8a311.
//
// Solidity: function summonCrystal(uint256 _summonerId, uint256 _assistantId, uint16 _summonerTears, uint16 _assistantTears, address _bonusItem) returns()
func (_Tavern *TavernSession) SummonCrystal(_summonerId *big.Int, _assistantId *big.Int, _summonerTears uint16, _assistantTears uint16, _bonusItem common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.SummonCrystal(&_Tavern.TransactOpts, _summonerId, _assistantId, _summonerTears, _assistantTears, _bonusItem)
}

// SummonCrystal is a paid mutator transaction binding the contract method 0x4ea8a311.
//
// Solidity: function summonCrystal(uint256 _summonerId, uint256 _assistantId, uint16 _summonerTears, uint16 _assistantTears, address _bonusItem) returns()
func (_Tavern *TavernTransactorSession) SummonCrystal(_summonerId *big.Int, _assistantId *big.Int, _summonerTears uint16, _assistantTears uint16, _bonusItem common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.SummonCrystal(&_Tavern.TransactOpts, _summonerId, _assistantId, _summonerTears, _assistantTears, _bonusItem)
}

// ToggleEnabled is a paid mutator transaction binding the contract method 0x3044b56c.
//
// Solidity: function toggleEnabled() returns()
func (_Tavern *TavernTransactor) ToggleEnabled(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "toggleEnabled")
}

// ToggleEnabled is a paid mutator transaction binding the contract method 0x3044b56c.
//
// Solidity: function toggleEnabled() returns()
func (_Tavern *TavernSession) ToggleEnabled() (*types.Transaction, error) {
	return _Tavern.Contract.ToggleEnabled(&_Tavern.TransactOpts)
}

// ToggleEnabled is a paid mutator transaction binding the contract method 0x3044b56c.
//
// Solidity: function toggleEnabled() returns()
func (_Tavern *TavernTransactorSession) ToggleEnabled() (*types.Transaction, error) {
	return _Tavern.Contract.ToggleEnabled(&_Tavern.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tavern *TavernTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tavern.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tavern *TavernSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.TransferOwnership(&_Tavern.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tavern *TavernTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tavern.Contract.TransferOwnership(&_Tavern.TransactOpts, newOwner)
}

// TavernAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the Tavern contract.
type TavernAuctionCancelledIterator struct {
	Event *TavernAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *TavernAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TavernAuctionCancelled)
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
		it.Event = new(TavernAuctionCancelled)
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
func (it *TavernAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TavernAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TavernAuctionCancelled represents a AuctionCancelled event raised by the Tavern contract.
type TavernAuctionCancelled struct {
	AuctionId *big.Int
	TokenId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0xdb9cc99dc874f9afbae71151f737e51547d3d412b52922793437d86607050c3c.
//
// Solidity: event AuctionCancelled(uint256 auctionId, uint256 indexed tokenId)
func (_Tavern *TavernFilterer) FilterAuctionCancelled(opts *bind.FilterOpts, tokenId []*big.Int) (*TavernAuctionCancelledIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tavern.contract.FilterLogs(opts, "AuctionCancelled", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TavernAuctionCancelledIterator{contract: _Tavern.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0xdb9cc99dc874f9afbae71151f737e51547d3d412b52922793437d86607050c3c.
//
// Solidity: event AuctionCancelled(uint256 auctionId, uint256 indexed tokenId)
func (_Tavern *TavernFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *TavernAuctionCancelled, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tavern.contract.WatchLogs(opts, "AuctionCancelled", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TavernAuctionCancelled)
				if err := _Tavern.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ParseAuctionCancelled is a log parse operation binding the contract event 0xdb9cc99dc874f9afbae71151f737e51547d3d412b52922793437d86607050c3c.
//
// Solidity: event AuctionCancelled(uint256 auctionId, uint256 indexed tokenId)
func (_Tavern *TavernFilterer) ParseAuctionCancelled(log types.Log) (*TavernAuctionCancelled, error) {
	event := new(TavernAuctionCancelled)
	if err := _Tavern.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TavernAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the Tavern contract.
type TavernAuctionCreatedIterator struct {
	Event *TavernAuctionCreated // Event containing the contract specifics and raw log

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
func (it *TavernAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TavernAuctionCreated)
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
		it.Event = new(TavernAuctionCreated)
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
func (it *TavernAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TavernAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TavernAuctionCreated represents a AuctionCreated event raised by the Tavern contract.
type TavernAuctionCreated struct {
	AuctionId     *big.Int
	Owner         common.Address
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Winner        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0x9a33d4a1b0a13cd8ff614a080df31b4b20c845e5cde181e3ae6f818f62b6ddde.
//
// Solidity: event AuctionCreated(uint256 auctionId, address indexed owner, uint256 indexed tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration, address winner)
func (_Tavern *TavernFilterer) FilterAuctionCreated(opts *bind.FilterOpts, owner []common.Address, tokenId []*big.Int) (*TavernAuctionCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tavern.contract.FilterLogs(opts, "AuctionCreated", ownerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TavernAuctionCreatedIterator{contract: _Tavern.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0x9a33d4a1b0a13cd8ff614a080df31b4b20c845e5cde181e3ae6f818f62b6ddde.
//
// Solidity: event AuctionCreated(uint256 auctionId, address indexed owner, uint256 indexed tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration, address winner)
func (_Tavern *TavernFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *TavernAuctionCreated, owner []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tavern.contract.WatchLogs(opts, "AuctionCreated", ownerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TavernAuctionCreated)
				if err := _Tavern.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0x9a33d4a1b0a13cd8ff614a080df31b4b20c845e5cde181e3ae6f818f62b6ddde.
//
// Solidity: event AuctionCreated(uint256 auctionId, address indexed owner, uint256 indexed tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration, address winner)
func (_Tavern *TavernFilterer) ParseAuctionCreated(log types.Log) (*TavernAuctionCreated, error) {
	event := new(TavernAuctionCreated)
	if err := _Tavern.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TavernAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the Tavern contract.
type TavernAuctionSuccessfulIterator struct {
	Event *TavernAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *TavernAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TavernAuctionSuccessful)
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
		it.Event = new(TavernAuctionSuccessful)
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
func (it *TavernAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TavernAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TavernAuctionSuccessful represents a AuctionSuccessful event raised by the Tavern contract.
type TavernAuctionSuccessful struct {
	AuctionId  *big.Int
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0xe40da2ed231723b222a7ba7da994c3afc3f83a51da76262083e38841e2da0982.
//
// Solidity: event AuctionSuccessful(uint256 auctionId, uint256 indexed tokenId, uint256 totalPrice, address winner)
func (_Tavern *TavernFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts, tokenId []*big.Int) (*TavernAuctionSuccessfulIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tavern.contract.FilterLogs(opts, "AuctionSuccessful", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TavernAuctionSuccessfulIterator{contract: _Tavern.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0xe40da2ed231723b222a7ba7da994c3afc3f83a51da76262083e38841e2da0982.
//
// Solidity: event AuctionSuccessful(uint256 auctionId, uint256 indexed tokenId, uint256 totalPrice, address winner)
func (_Tavern *TavernFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *TavernAuctionSuccessful, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tavern.contract.WatchLogs(opts, "AuctionSuccessful", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TavernAuctionSuccessful)
				if err := _Tavern.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// ParseAuctionSuccessful is a log parse operation binding the contract event 0xe40da2ed231723b222a7ba7da994c3afc3f83a51da76262083e38841e2da0982.
//
// Solidity: event AuctionSuccessful(uint256 auctionId, uint256 indexed tokenId, uint256 totalPrice, address winner)
func (_Tavern *TavernFilterer) ParseAuctionSuccessful(log types.Log) (*TavernAuctionSuccessful, error) {
	event := new(TavernAuctionSuccessful)
	if err := _Tavern.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TavernCrystalOpenIterator is returned from FilterCrystalOpen and is used to iterate over the raw logs and unpacked data for CrystalOpen events raised by the Tavern contract.
type TavernCrystalOpenIterator struct {
	Event *TavernCrystalOpen // Event containing the contract specifics and raw log

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
func (it *TavernCrystalOpenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TavernCrystalOpen)
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
		it.Event = new(TavernCrystalOpen)
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
func (it *TavernCrystalOpenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TavernCrystalOpenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TavernCrystalOpen represents a CrystalOpen event raised by the Tavern contract.
type TavernCrystalOpen struct {
	Owner     common.Address
	CrystalId *big.Int
	HeroId    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCrystalOpen is a free log retrieval operation binding the contract event 0xa6537abb32df743f25343a89920580e228e72ebe25cbbb40b6d83ce4aab0a425.
//
// Solidity: event CrystalOpen(address indexed owner, uint256 crystalId, uint256 heroId)
func (_Tavern *TavernFilterer) FilterCrystalOpen(opts *bind.FilterOpts, owner []common.Address) (*TavernCrystalOpenIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Tavern.contract.FilterLogs(opts, "CrystalOpen", ownerRule)
	if err != nil {
		return nil, err
	}
	return &TavernCrystalOpenIterator{contract: _Tavern.contract, event: "CrystalOpen", logs: logs, sub: sub}, nil
}

// WatchCrystalOpen is a free log subscription operation binding the contract event 0xa6537abb32df743f25343a89920580e228e72ebe25cbbb40b6d83ce4aab0a425.
//
// Solidity: event CrystalOpen(address indexed owner, uint256 crystalId, uint256 heroId)
func (_Tavern *TavernFilterer) WatchCrystalOpen(opts *bind.WatchOpts, sink chan<- *TavernCrystalOpen, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Tavern.contract.WatchLogs(opts, "CrystalOpen", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TavernCrystalOpen)
				if err := _Tavern.contract.UnpackLog(event, "CrystalOpen", log); err != nil {
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

// ParseCrystalOpen is a log parse operation binding the contract event 0xa6537abb32df743f25343a89920580e228e72ebe25cbbb40b6d83ce4aab0a425.
//
// Solidity: event CrystalOpen(address indexed owner, uint256 crystalId, uint256 heroId)
func (_Tavern *TavernFilterer) ParseCrystalOpen(log types.Log) (*TavernCrystalOpen, error) {
	event := new(TavernCrystalOpen)
	if err := _Tavern.contract.UnpackLog(event, "CrystalOpen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TavernCrystalSummonedIterator is returned from FilterCrystalSummoned and is used to iterate over the raw logs and unpacked data for CrystalSummoned events raised by the Tavern contract.
type TavernCrystalSummonedIterator struct {
	Event *TavernCrystalSummoned // Event containing the contract specifics and raw log

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
func (it *TavernCrystalSummonedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TavernCrystalSummoned)
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
		it.Event = new(TavernCrystalSummoned)
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
func (it *TavernCrystalSummonedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TavernCrystalSummonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TavernCrystalSummoned represents a CrystalSummoned event raised by the Tavern contract.
type TavernCrystalSummoned struct {
	CrystalId      *big.Int
	Owner          common.Address
	SummonerId     *big.Int
	AssistantId    *big.Int
	Generation     uint16
	CreatedBlock   *big.Int
	SummonerTears  uint8
	AssistantTears uint8
	BonusItem      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCrystalSummoned is a free log retrieval operation binding the contract event 0x4508aba30a57c0fc7f1d5da83dea7dd0c36368a7080d3a6652fcd5a58168f460.
//
// Solidity: event CrystalSummoned(uint256 crystalId, address indexed owner, uint256 summonerId, uint256 assistantId, uint16 generation, uint256 createdBlock, uint8 summonerTears, uint8 assistantTears, address bonusItem)
func (_Tavern *TavernFilterer) FilterCrystalSummoned(opts *bind.FilterOpts, owner []common.Address) (*TavernCrystalSummonedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Tavern.contract.FilterLogs(opts, "CrystalSummoned", ownerRule)
	if err != nil {
		return nil, err
	}
	return &TavernCrystalSummonedIterator{contract: _Tavern.contract, event: "CrystalSummoned", logs: logs, sub: sub}, nil
}

// WatchCrystalSummoned is a free log subscription operation binding the contract event 0x4508aba30a57c0fc7f1d5da83dea7dd0c36368a7080d3a6652fcd5a58168f460.
//
// Solidity: event CrystalSummoned(uint256 crystalId, address indexed owner, uint256 summonerId, uint256 assistantId, uint16 generation, uint256 createdBlock, uint8 summonerTears, uint8 assistantTears, address bonusItem)
func (_Tavern *TavernFilterer) WatchCrystalSummoned(opts *bind.WatchOpts, sink chan<- *TavernCrystalSummoned, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Tavern.contract.WatchLogs(opts, "CrystalSummoned", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TavernCrystalSummoned)
				if err := _Tavern.contract.UnpackLog(event, "CrystalSummoned", log); err != nil {
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

// ParseCrystalSummoned is a log parse operation binding the contract event 0x4508aba30a57c0fc7f1d5da83dea7dd0c36368a7080d3a6652fcd5a58168f460.
//
// Solidity: event CrystalSummoned(uint256 crystalId, address indexed owner, uint256 summonerId, uint256 assistantId, uint16 generation, uint256 createdBlock, uint8 summonerTears, uint8 assistantTears, address bonusItem)
func (_Tavern *TavernFilterer) ParseCrystalSummoned(log types.Log) (*TavernCrystalSummoned, error) {
	event := new(TavernCrystalSummoned)
	if err := _Tavern.contract.UnpackLog(event, "CrystalSummoned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TavernOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tavern contract.
type TavernOwnershipTransferredIterator struct {
	Event *TavernOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TavernOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TavernOwnershipTransferred)
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
		it.Event = new(TavernOwnershipTransferred)
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
func (it *TavernOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TavernOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TavernOwnershipTransferred represents a OwnershipTransferred event raised by the Tavern contract.
type TavernOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tavern *TavernFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TavernOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tavern.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TavernOwnershipTransferredIterator{contract: _Tavern.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tavern *TavernFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TavernOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tavern.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TavernOwnershipTransferred)
				if err := _Tavern.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tavern *TavernFilterer) ParseOwnershipTransferred(log types.Log) (*TavernOwnershipTransferred, error) {
	event := new(TavernOwnershipTransferred)
	if err := _Tavern.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TavernPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Tavern contract.
type TavernPausedIterator struct {
	Event *TavernPaused // Event containing the contract specifics and raw log

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
func (it *TavernPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TavernPaused)
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
		it.Event = new(TavernPaused)
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
func (it *TavernPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TavernPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TavernPaused represents a Paused event raised by the Tavern contract.
type TavernPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Tavern *TavernFilterer) FilterPaused(opts *bind.FilterOpts) (*TavernPausedIterator, error) {

	logs, sub, err := _Tavern.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TavernPausedIterator{contract: _Tavern.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Tavern *TavernFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TavernPaused) (event.Subscription, error) {

	logs, sub, err := _Tavern.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TavernPaused)
				if err := _Tavern.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Tavern *TavernFilterer) ParsePaused(log types.Log) (*TavernPaused, error) {
	event := new(TavernPaused)
	if err := _Tavern.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TavernRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Tavern contract.
type TavernRoleAdminChangedIterator struct {
	Event *TavernRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TavernRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TavernRoleAdminChanged)
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
		it.Event = new(TavernRoleAdminChanged)
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
func (it *TavernRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TavernRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TavernRoleAdminChanged represents a RoleAdminChanged event raised by the Tavern contract.
type TavernRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Tavern *TavernFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TavernRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Tavern.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TavernRoleAdminChangedIterator{contract: _Tavern.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Tavern *TavernFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TavernRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Tavern.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TavernRoleAdminChanged)
				if err := _Tavern.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Tavern *TavernFilterer) ParseRoleAdminChanged(log types.Log) (*TavernRoleAdminChanged, error) {
	event := new(TavernRoleAdminChanged)
	if err := _Tavern.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TavernRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Tavern contract.
type TavernRoleGrantedIterator struct {
	Event *TavernRoleGranted // Event containing the contract specifics and raw log

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
func (it *TavernRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TavernRoleGranted)
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
		it.Event = new(TavernRoleGranted)
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
func (it *TavernRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TavernRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TavernRoleGranted represents a RoleGranted event raised by the Tavern contract.
type TavernRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Tavern *TavernFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TavernRoleGrantedIterator, error) {

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

	logs, sub, err := _Tavern.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TavernRoleGrantedIterator{contract: _Tavern.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Tavern *TavernFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TavernRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Tavern.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TavernRoleGranted)
				if err := _Tavern.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Tavern *TavernFilterer) ParseRoleGranted(log types.Log) (*TavernRoleGranted, error) {
	event := new(TavernRoleGranted)
	if err := _Tavern.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TavernRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Tavern contract.
type TavernRoleRevokedIterator struct {
	Event *TavernRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TavernRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TavernRoleRevoked)
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
		it.Event = new(TavernRoleRevoked)
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
func (it *TavernRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TavernRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TavernRoleRevoked represents a RoleRevoked event raised by the Tavern contract.
type TavernRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Tavern *TavernFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TavernRoleRevokedIterator, error) {

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

	logs, sub, err := _Tavern.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TavernRoleRevokedIterator{contract: _Tavern.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Tavern *TavernFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TavernRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Tavern.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TavernRoleRevoked)
				if err := _Tavern.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Tavern *TavernFilterer) ParseRoleRevoked(log types.Log) (*TavernRoleRevoked, error) {
	event := new(TavernRoleRevoked)
	if err := _Tavern.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TavernUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Tavern contract.
type TavernUnpausedIterator struct {
	Event *TavernUnpaused // Event containing the contract specifics and raw log

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
func (it *TavernUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TavernUnpaused)
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
		it.Event = new(TavernUnpaused)
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
func (it *TavernUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TavernUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TavernUnpaused represents a Unpaused event raised by the Tavern contract.
type TavernUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Tavern *TavernFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TavernUnpausedIterator, error) {

	logs, sub, err := _Tavern.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TavernUnpausedIterator{contract: _Tavern.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Tavern *TavernFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TavernUnpaused) (event.Subscription, error) {

	logs, sub, err := _Tavern.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TavernUnpaused)
				if err := _Tavern.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Tavern *TavernFilterer) ParseUnpaused(log types.Log) (*TavernUnpaused, error) {
	event := new(TavernUnpaused)
	if err := _Tavern.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
