//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint8
	energy, maxEnergy uint
}

func main() {

	player := Player{
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
		name:      "birey",
	}

	player.applyDamage(10)
	player.consumeEnergy(100)
	player.consumeEnergy(10)
	player.applyDamage(89)
	player.addHealth(40)
	player.applyDamage(60)
	player.consumeEnergy(300)
	player.addEnergy(30)
	player.consumeEnergy(100)
	player.addHealth(30)

}

func (p *Player) addHealth(amount uint8) {
	p.health += amount

	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}

	fmt.Println(p.name+" added", amount, "health ->", p.health)
}

func (p *Player) applyDamage(amount uint8) {
	if p.health-amount > p.health {
		p.health = 0
	} else {
		p.health -= amount
	}

	fmt.Println(p.name+" took damage", amount, "health ->", p.health)
}

func (p *Player) addEnergy(amount uint) {
	p.energy += amount

	if p.energy > p.maxEnergy {
		p.energy = p.maxEnergy
	}

	fmt.Println(p.name+" added", amount, "energy ->", p.energy)
}

func (p *Player) consumeEnergy(amount uint) {
	if p.energy-amount > p.energy {
		p.energy = 0
	} else {
		p.energy -= amount
	}

	fmt.Println(p.name+" consumed energy", amount, "energy ->", p.energy)
}
