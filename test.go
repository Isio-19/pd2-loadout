// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"html/template"
	"net/http"

	"log"
)

type FStats struct {
	MagasineSize, TotalAmmo, RateOfFire, Damage, Accuracy, Stability, Concealment, Threat int
}

type PWeapon struct {
	Id    string
	Stats FStats
}

type SWeapon struct {
	Id    string
	Stats FStats
}

type MStats struct {
	Damage, Knockdown, Concealment int
	ChargeTime, Range              float32
}

type MWeapon struct {
	Id    string
	Stats MStats
}

type Throwable struct {
	Id             string
	Damage, Amount int
}

type Equipment struct {
	Id               string
	Amount, Capacity int
}

type Armor struct {
	Id                                    string
	Armor, Concealment, Dodge, Steadiness int
	Speed, Stamina                        float32
}

type SkillTree struct {
	Mastermind MastermindSkills
	Enforcer   EnforcerSkills
	Technician TechnicianSkills
	Ghost      GhostSkills
	Fugitive   FugitiveSkills
}

type MastermindSkills struct {
}

type EnforcerSkills struct {
}

type TechnicianSkills struct {
}

type GhostSkills struct {
}

type FugitiveSkills struct {
}

type PerkDeck struct {
	Id    string
	Level int
}

type Loadout struct {
	PrimaryWeapon   PWeapon
	SecondaryWeapon SWeapon
	MeleeWeapon     MWeapon
	Throwable       Throwable
	Equipment       Equipment
	Armor           Armor
	SkillTree       SkillTree
	PerkDeck        PerkDeck
}

func redirectToIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "./index.html", http.StatusSeeOther)
}

func loadIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")

	loadout := &Loadout{
		PrimaryWeapon: PWeapon{
			Id: "amcar",
			Stats: FStats{
				MagasineSize: 0,
				TotalAmmo:    0,
				RateOfFire:   0,
				Damage:       0,
				Accuracy:     0,
				Stability:    0,
				Concealment:  0,
				Threat:       0,
			},
		},
		SecondaryWeapon: SWeapon{
			Id: "glock_17",
			Stats: FStats{
				MagasineSize: 0,
				TotalAmmo:    0,
				RateOfFire:   0,
				Damage:       0,
				Accuracy:     0,
				Stability:    0,
				Concealment:  0,
				Threat:       0,
			},
		},
		MeleeWeapon: MWeapon{
			Id: "weapon",
			Stats: MStats{
				Damage:      0,
				Knockdown:   0,
				ChargeTime:  0,
				Range:       0,
				Concealment: 0,
			},
		},
		Throwable: Throwable{
			Id:     "concussion",
			Damage: 0,
			Amount: 0,
		},
		Equipment: Equipment{
			Id:       "ammo_bag",
			Amount:   0,
			Capacity: 0,
		},
		Armor: Armor{
			Id:          "two_piece_suit",
			Armor:       0,
			Concealment: 0,
			Speed:       0,
			Dodge:       0,
			Steadiness:  0,
			Stamina:     0,
		},
		SkillTree: SkillTree{
			Mastermind: MastermindSkills{},
			Enforcer:   EnforcerSkills{},
			Technician: TechnicianSkills{},
			Ghost:      GhostSkills{},
			Fugitive:   FugitiveSkills{},
		},
		PerkDeck: PerkDeck{
			Id:    "crew_chief",
			Level: 0,
		},
	}

	t.Execute(w, loadout)
}

func main() {
	http.HandleFunc("/", redirectToIndex)
	http.HandleFunc("/index.html", loadIndex)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
