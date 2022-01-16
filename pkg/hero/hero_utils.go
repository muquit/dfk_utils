package hero

// Adapted from https://github.com/0rtis/dfk.git python code
//  muquit@muquit.com - January-15-2022 15:04:06

import (
	"fmt"
	"math/big"
)

type StatGenes struct {
    Raw *big.Int
    Class string
    SubClass string
    Profession string
    Passive1 int
    Passive2 int
    StatBoost1 string
    StatBoost2 string
    StatsUnknown1 string
    Element string
    StatsUnknown2 string
}
// The maps are no used. 
var Rarity = map[uint8]string {
	0: "common",
	1: "uncommon",
	2: "rare",
	3: "legendary",
	4: "mythic",
}

var Class = map[uint8]string {
    0: "warrior",
    1: "knight",
    2: "thief",
    3: "archer",
    4: "priest",
    5: "wizard",
    6: "monk",
    7: "pirate",
    16: "paladin",
    17: "darkKnight",
    18: "summoner",
    19: "ninja",
    24: "dragoon",
    25: "sage",
    28: "dreadKnight",
}

var VisualTraits = map[uint8]string {
    0: "gender",
    1: "headAppendage",
    2: "backAppendage",
    3: "background",
    4: "hairStyle",
    5: "hairColor",
    6: "visualUnknown1",
    7: "eyeColor",
    8: "skinColor",
    9: "appendageColor",
    10: "backAppendageColor",
    11: "visualUnknown2",
}

var StatTraits = map[uint8]string {
    0: "class",
    1: "subClass",
    2: "profession",
    3: "passive1",
    4: "passive2",
    5: "active1",
    6: "active2",
    7: "statBoost1",
    8: "statBoost2",
    9: "statsUnknown1",
    10: "element",
    11: "statsUnknown2",
}

var Professions = map[uint8]string {
    0: "mining",
    2: "gardening",
    4: "fishing",
    6: "foraging",
}

var Stats = map[uint8]string {
    0: "strength",
    2: "agility",
    4: "intelligence",
    6: "wisdom",
    8: "luck",
    10: "vitality",
    12: "endurance",
    14: "dexterity",
}

var Elements = map[uint8]string {
    0: "fire",
    2: "water",
    4: "earth",
    6: "wind",
    8: "lightning",
    10: "ice",
    12: "light",
    14: "dark",
}

// not used
func ParseRarity(rarity uint8) (string, error) {
    if _, ok := Rarity[rarity]; ok {
        return Rarity[rarity], nil
    }
    return "None", fmt.Errorf("Hero Rarity %v not found", rarity)
}

// not used
func ParseClass(class uint8) (string, error) {
    if _, ok := Class[class]; ok {
        return Class[class], nil
    }
    return "None", fmt.Errorf("Hero class %v not found", class)
}

// not used
func ParseProfession(profession uint8) (string, error) {
    if _, ok := Professions [profession]; ok {
        return Professions[profession], nil
    }
    return "None", fmt.Errorf("Hero profession %v not found", profession)
}

// not used
func ParseStat(stat uint8) (string, error) {
    if _, ok := Stats[stat]; ok {
        return Stats[stat], nil
    }
    return "None", fmt.Errorf("Hero stat %v not found", stat)
}

// not used
func ParseElement(element uint8) (string, error) {
    if _, ok := Elements[element]; ok {
        return Elements[element], nil
    }
    return "None", fmt.Errorf("Hero element %v not found", element)
}


// Parse Genenes
// TODO
func ParseStatGeneses(statGenes *big.Int) (*StatGenes, error) {
	alphabet := "123456789abcdefghijkmnopqrstuvwx"
	base := len(alphabet)
	baseBigInt := big.NewInt(32)
	fmt.Println(alphabet)
	fmt.Println(base)
	genes := new(big.Int)
	genes, ok := genes.SetString("55595053337262517174437940546058771473513094722680050621928661284030532", 10)
	if ! ok {
        return nil, fmt.Errorf("Could not convert genes to string")
	}
	for genes.Cmp(baseBigInt) >= 0 {
		mod := big.NewInt(0).Mod(genes, baseBigInt)
		fmt.Println(mod)
		x := big.NewInt(0).Sub(genes, mod)
		fmt.Println(x)
		genes = big.NewInt(0).Div(x,baseBigInt)
		fmt.Println(genes)
	}
    sg := new(StatGenes)
    sg.Raw = statGenes
    return sg, nil
}