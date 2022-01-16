# Introduction

Some utilities in [golang](https://go.dev) for interacting with smart contracts of [defikingdoms](https://defikingdoms.com/). Inspired by https://github.com/0rtis/dfk python code

**It's a work in progress...**

# Compiling

Before compiling, install go from https://golang.org/. The binaries can be compiled for following Platforms:

- MacOS
- Linux
- Windows
- Raspberry PI
- Raspberry PI Jessie

Type:

```
    make
```

If you are on Windows and don't have `make`, then just to compile native binaries, type:

```
    go build -o ./bin ./...
```

- Look at `Makefile` for exact `go build` commands used for compiling.
- Look at `Makefile` on how code is generated from contract ABI

All required modules will be downloaded automatically.

# Utilities

Currently the following utilities are developed

- _hero_info_ - Display Hero information in JSON to stdout. Compiled binaries are available for MacOS, Linux, Windows, Raspberry PI in ./bin directory

## Usage

```
Usage of bin/hero_info:
  -hid int
    	Hero Id (default 1)
```

**TODO**: Stats and Visual Genes decoding

**Example**

```
$ ./bin/hero_info -hid 1
{
  "OwnerAddress": "0x2e7669f61ea77f02445a015fbdcfe2de47083e02",
  "HeroInfo": {
    "Id": 1,
    "SummoningInfo": {
      "SummonedTime": 1633044270,
      "NextSummonTime": 1633044270,
      "SummonerId": 0,
      "AssistantId": 0,
      "Summons": 0,
      "MaxSummons": 11
    },
    "Info": {
      "StatGenes": 55595053337262517174437940546058771473513094722680050621928661284030532,
      "VisualGenes": 170877259497388213840353281232231169976585066888929467746175634464967719,
      "Rarity": 4,
      "Shiny": true,
      "Generation": 0,
      "FirstName": 184,
      "LastName": 85,
      "ShinyStyle": 4,
      "Class": 2,
      "SubClass": 5
    },
    "State": {
      "StaminaFullAt": 1641856758,
      "HpFullAt": 0,
      "MpFullAt": 0,
      "Level": 2,
      "Xp": 1371,
      "CurrentQuest": "0x6ff019415ee105acf2ac52483a33f5b43eadb8d0",
      "Sp": 0,
      "Status": 0
    },
    "Stats": {
      "Strength": 12,
      "Intelligence": 9,
      "Wisdom": 9,
      "Luck": 14,
      "Agility": 13,
      "Vitality": 8,
      "Endurance": 7,
      "Dexterity": 9,
      "Hp": 177,
      "Mp": 51,
      "Stamina": 26
    },
    "PrimaryStatGrowth": {
      "Strength": 5500,
      "Intelligence": 2500,
      "Wisdom": 3500,
      "Luck": 6700,
      "Agility": 7000,
      "Vitality": 5000,
      "Endurance": 4500,
      "Dexterity": 5500,
      "HpSm": 2500,
      "HpRg": 5000,
      "HpLg": 2500,
      "MpSm": 3000,
      "MpRg": 4000,
      "MpLg": 3000
    },
    "SecondaryStatGrowth": {
      "Strength": 750,
      "Intelligence": 2000,
      "Wisdom": 2000,
      "Luck": 1400,
      "Agility": 1000,
      "Vitality": 1250,
      "Endurance": 1250,
      "Dexterity": 750,
      "HpSm": 875,
      "HpRg": 1000,
      "HpLg": 625,
      "MpSm": 375,
      "MpRg": 875,
      "MpLg": 1250
    },
    "Professions": {
      "Mining": 1,
      "Gardening": 5,
      "Foraging": 15,
      "Fishing": 16
    }
  }
}

```
