package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Dukedom struct {
	year         int
	population   int
	grainStores  int
	land         int
	treasury     int
	army         int
	grainPerAcre int
}

func (d *Dukedom) displayStatus() {
	fmt.Printf("\nYear: %d\n", d.year)
	fmt.Printf("Population: %d\n", d.population)
	fmt.Printf("Grain Stores: %d bushels\n", d.grainStores)
	fmt.Printf("Land: %d acres\n", d.land)
	fmt.Printf("Treasury: %d gold coins\n", d.treasury)
	fmt.Printf("Army: %d soldiers\n\n", d.army)
}

func (d *Dukedom) plantGrain() {
	for {
		fmt.Printf("How many acres will you plant? (0-%d): ", d.land)
		var acresInput string
		fmt.Scanln(&acresInput)
		acres, err := strconv.Atoi(acresInput)
		if err != nil || acres < 0 || acres > d.land {
			fmt.Println("Please enter a valid number.")
			continue
		}

		requiredGrain := acres / 2
		if requiredGrain <= d.grainStores {
			d.grainStores -= requiredGrain
			yieldGrain := acres * d.grainPerAcre
			d.grainStores += yieldGrain
			fmt.Printf("You planted %d acres, using %d bushels of grain.\n", acres, requiredGrain)
			fmt.Printf("You harvested %d bushels of grain.\n", yieldGrain)
			break
		} else {
			fmt.Println("You don't have enough grain to plant that many acres.")
		}
	}
}

func (d *Dukedom) feedPopulation() {
	for {
		fmt.Printf("How much grain will you distribute to your people? (0-%d): ", d.grainStores)
		var grainInput string
		fmt.Scanln(&grainInput)
		grain, err := strconv.Atoi(grainInput)
		if err != nil || grain < 0 || grain > d.grainStores {
			fmt.Println("Please enter a valid number.")
			continue
		}

		d.grainStores -= grain
		foodNeeded := d.population * 2
		if grain < foodNeeded {
			deaths := (foodNeeded - grain) / 2
			d.population -= deaths
			fmt.Printf("%d people starved.\n", deaths)
		} else {
			fmt.Println("Everyone has been fed.")
		}
		break
	}
}

func (d *Dukedom) recruitArmy() {
	for {
		fmt.Print("How many soldiers will you recruit? (0-50): ")
		var recruitsInput string
		fmt.Scanln(&recruitsInput)
		recruits, err := strconv.Atoi(recruitsInput)
		cost := recruits * 10
		if err != nil || recruits < 0 || recruits > 50 || cost > d.treasury {
			fmt.Println("You can't afford that many recruits or invalid input.")
			continue
		}

		d.army += recruits
		d.treasury -= cost
		fmt.Printf("You recruited %d soldiers, costing %d gold coins.\n", recruits, cost)
		break
	}
}

func (d *Dukedom) collectTaxes() {
	taxIncome := d.population * 3
	d.treasury += taxIncome
	fmt.Printf("You collected %d gold coins in taxes.\n", taxIncome)
}

func (d *Dukedom) randomEvents() {
	rand.Seed(time.Now().UnixNano())
	events := []string{"good_harvest", "bad_harvest", "enemy_attack", "plague", ""}
	event := events[rand.Intn(len(events))]

	switch event {
	case "good_harvest":
		bonus := d.land / 5
		d.grainStores += bonus
		fmt.Printf("A bumper crop! You gain %d bushels of grain.\n", bonus)
	case "bad_harvest":
		loss := d.land / 5
		d.grainStores -= loss
		fmt.Printf("A drought hit your fields. You lost %d bushels of grain.\n", loss)
	case "enemy_attack":
		soldiersLost := rand.Intn(d.army / 2)
		if soldiersLost < d.army {
			d.army -= soldiersLost
			fmt.Printf("Your dukedom was attacked! %d soldiers were lost defending the realm.\n", soldiersLost)
		} else {
			fmt.Println("The enemy overran your army! Your dukedom has fallen.")
			d.population = 0 // Game Over
		}
	case "plague":
		plagueDeaths := rand.Intn(d.population / 3)
		d.population -= plagueDeaths
		fmt.Printf("A plague swept through your dukedom, killing %d people.\n", plagueDeaths)
	}
}

func (d *Dukedom) nextYear() {
	d.year++
	populationGrowth := rand.Intn(10)
	d.population += populationGrowth
	fmt.Printf("%d people were born this year.\n", populationGrowth)
}

func (d *Dukedom) run() {
	fmt.Println("Welcome to Dukedom! Lead your people wisely.")
	for d.population > 0 && d.grainStores > 0 {
		d.displayStatus()
		d.collectTaxes()
		d.plantGrain()
		d.feedPopulation()
		d.recruitArmy()
		d.randomEvents()
		d.nextYear()
	}
	fmt.Println("Your reign has ended. The dukedom has failed.")
}

func main() {
	dukedom := Dukedom{
		year:         1,
		population:   100,
		grainStores:  1000,
		land:         500,
		treasury:     500,
		army:         20,
		grainPerAcre: 3,
	}
	dukedom.run()
}
