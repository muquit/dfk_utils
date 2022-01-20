package hero

// Adapted from https://github.com/0rtis/dfk.git python code
//  muquit@muquit.com - January-15-2022 15:04:06

import (
	"fmt"
	"math/big"
	"strings"
)
const (
    alphabet = "123456789abcdefghijkmnopqrstuvwx"
    alphabetLen = 32
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

var Class = map[int]string {
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
    1: "subclass",
    2: "profession",
    3: "passive1",
    4: "passive2",
    5: "active1",
    6: "active2",
    7: "statboost1",
    8: "statboost2",
    9: "statsunknown1",
    10: "element",
    11: "statsunknown2",
}

var Professions = map[int]string {
    0: "mining",
    2: "gardening",
    4: "fishing",
    6: "foraging",
}

var Stats = map[int]string {
    0: "strength",
    2: "agility",
    4: "intelligence",
    6: "wisdom",
    8: "luck",
    10: "vitality",
    12: "endurance",
    14: "dexterity",
}

var Elements = map[int]string {
    0: "fire",
    2: "water",
    4: "earth",
    6: "wind",
    8: "lightning",
    10: "ice",
    12: "light",
    14: "dark",
}

func RightJust(s string, n int, fill string) string {
	if len(s) < n {
		return strings.Repeat(fill, n) + s
	}
	return s
}

// not used
func ParseRarity(rarity uint8) (string, error) {
    if _, ok := Rarity[rarity]; ok {
        return Rarity[rarity], nil
    }
    return "None", fmt.Errorf("Hero Rarity %v not found", rarity)
}

// not used
func ParseClass(class int) (string, error) {
    if _, ok := Class[class]; ok {
        return Class[class], nil
    }
    return "None", fmt.Errorf("Hero class %v not found", class)
}

// not used
func ParseProfession(profession int) (string, error) {
    if _, ok := Professions [profession]; ok {
        return Professions[profession], nil
    }
    return "None", fmt.Errorf("Hero profession %v not found", profession)
}

// not used
func ParseStat(stat int) (string, error) {
    if _, ok := Stats[stat]; ok {
        return Stats[stat], nil
    }
    return "None", fmt.Errorf("Hero stat %v not found", stat)
}

// not used
func ParseElement(element int) (string, error) {
    if _, ok := Elements[element]; ok {
        return Elements[element], nil
    }
    return "None", fmt.Errorf("Hero element %v not found", element)
}

// map genens to 48 characters string of alphabets and return
// the string
func AlphabetaizeGenes(genes *big.Int)(string, error) {
    alenBigInt := big.NewInt(alphabetLen)
    abuf := "" // returns
    for genes.Cmp(alenBigInt) >= 0 {
        // mod = genes % 32
        mod := big.NewInt(0).Mod(genes, alenBigInt)
        if big.NewInt(0).Cmp(mod) >= alphabetLen {
            return abuf, fmt.Errorf("Invalid Genes")
        }

        // get the character from alphabet[mod]
        a := string(alphabet[mod.Int64()])

        // keep appending to abuf until the entire gene is processed
        abuf = a + abuf
        
        // subract mod from genes: genes - mod
        sg := big.NewInt(0).Sub(genes, mod)
        // genes = (genes - mod ) / 32
        genes = big.NewInt(0).Div(sg, alenBigInt)
    }
    abuf = string(alphabet[genes.Int64()]) + abuf
    // leftpad with 1 upto 48 if needed
    abuf = RightJust(abuf, 48, "1")
    return abuf, nil
}

// Parse Genenes
// TODO
func DecodeStatGenes(statGenes *big.Int) (*StatGenes, error) {
    abuf, err := AlphabetaizeGenes(statGenes)
    if err != nil {
        return nil, fmt.Errorf("Could not decode stat genes")
    }
    ar := []rune(abuf)
    f := ""
    res := ""
    m := make(map[string]int)
    for i, r := range(ar) {
        trait_idx := i /4
        stat_trait := StatTraits[uint8(trait_idx)]
        kai := string(abuf[i])
        value := strings.Index(alphabet,kai)
        fmt.Println(trait_idx, stat_trait,kai, value)
        m[stat_trait] = value
        res = res + string(r)
        if i > 0 && (i + 1) % 4 == 0 {
            f = f + " " + res
            res = ""
        }
    }
    fmt.Println(abuf)
    fmt.Println(m)
    for key,val := range(m) {
        fmt.Println("Key:", key, "=>", "Value:", val)
    }
    s := new(StatGenes)
    s.Raw = statGenes
    s.Class, _ = ParseClass(m["class"])
    s.SubClass, _ = ParseClass(m["subclass"])
    s.Profession, _ = ParseProfession(m["profession"])
    s.StatBoost1, _ = ParseStat(m["statboost1"])
    s.StatBoost2, _ = ParseStat(m["statboost2"])
    s.Element, _ = ParseElement(m["element"])
    s.StatsUnknown1, _ = ParseStat(m["statsunknown1"])
    s.StatsUnknown2, _ = ParseStat(m["statsunknown2"])
    return s, nil
}