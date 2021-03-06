# Introduction

Some utilities in [golang](https://go.dev) for interacting with smart contracts of [defikingdoms](https://defikingdoms.com/). Inspired by https://github.com/0rtis/dfk python code

**It's a work in progress...**

# Compiling

Before compiling, install go from https://go.dev/. The binaries can be compiled for following Platforms:

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

Compiled binaries are available for MacOS, Linux, Windows, Raspberry PI in ./bin directory

- _hero_info_ - Print Hero information from blockchain in JSON to stdout.
- _profile_info_ - Print profile information from blockchain to stdout in JSON format.
- _land_info_ - Print land information from blockchain to stdout in JSON format.

## Usage

```
Usage of bin/hero_info:
  -hid int
    	Hero Id (default 1)
```

```
Usage of bin/profile_info:
  -addr string
    	Profile Hex Address
```

```
Usage of bin/land_info:
  -lid int
    	Land Id. If not specified all lands will be displayed
```

**TODO**: Stats and Visual Genes decoding

**Example**

**Print Hero information**

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

**Print Profile inforamtion**

```
$ ./bin/profile_info -addr 0x2E7669F61eA77F02445A015FBdcFe2DE47083E02
{
   "id": 38,
   "owner_address": "0x2e7669f61ea77f02445a015fbdcfe2de47083e02",
   "name": "FriskyFox",
   "created_on_epoch": 1629680114,
   "created_on_time_local": "2021-08-22 20:55:14 -0400 EDT",
   "pic_id": 6,
   "hero_id": 0
}
```

**Print Land inforamtion**

```
$ ./bin/land_info
```

```
[
  {
    "LandId": 1,
    "Name": "Wevelmont",
    "Owner": "0x8fb7ae993198a0369efe959ec99b6674cf672731",
    "Region": 0,
    "Level": 1,
    "Steward": 0,
    "Score": 0
  },
  {
    "LandId": 2,
    "Name": "Vohenland",
    "Owner": "0x77d991987ca85214f9686131c58c1abe4c93e547",
    "Region": 0,
    "Level": 1,
    "Steward": 0,
    "Score": 0
  },
  {
    "LandId": 3,
    "Name": "Gothune",
    "Owner": "0x6880add29b5e84c543ff745856a0f3368e58f165",
    "Region": 0,
    "Level": 1,
    "Steward": 0,
    "Score": 0
  },
  ....
```

Show just the Owners' names:

```
$ ./bin/land_info | jq '.[].Name'
```

To print info about a specific land:

```
$ ./bin/land_info -lid 10
{
   "LandId": 10,
   "Name": "Ashem",
   "Owner": "0x8ca47f7bd4c284bfbea3ca0f052eaf98f2990261",
   "Region": 0,
   "Level": 1,
   "Steward": 0,
   "Score": 0
}
```
