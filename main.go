package main

import (
	"fmt"
	"time"
)

type Tamagotchi struct {
	Name        string
	Hunger      int
	Happiness   int
	LastUpdated time.Time
	IsAlive     bool
}

func NewTamagotchi(name string) *Tamagotchi {
	return &Tamagotchi{
		Name:        name,
		Hunger:      50,
		Happiness:   50,
		LastUpdated: time.Now(),
		IsAlive:     true,
	}
}

func (t *Tamagotchi) Feed() {
	if t.IsAlive {
		t.Hunger -= 10
		if t.Hunger < 0 {
			t.Hunger = 0
		}
		t.updateLastUpdated()
	} else {
		fmt.Println("Your Tamagotchi is no longer alive.")
	}
}

func (t *Tamagotchi) Play() {
	if t.IsAlive {
		t.Happiness += 10
		if t.Happiness > 100 {
			t.Happiness = 100
		}
		t.updateLastUpdated()
	} else {
		fmt.Println("Your Tamagotchi is no longer alive.")
	}
}

func (t *Tamagotchi) updateLastUpdated() {
	t.LastUpdated = time.Now()
}

func (t *Tamagotchi) CheckStatus() {
	if t.IsAlive {
		fmt.Printf("Name: %s\nHunger: %d\nHappiness: %d\n", t.Name, t.Hunger, t.Happiness)
	} else {
		fmt.Println("Your Tamagotchi is no longer alive.")
	}
}

func (t *Tamagotchi) Live() {
	for t.IsAlive {
		time.Sleep(5 * time.Second) // Update status every 5 seconds
		t.Hunger += 5
		t.Happiness -= 5
		if t.Hunger >= 100 || t.Happiness <= 0 {
			t.IsAlive = false
			fmt.Printf("%s has passed away. Your Tamagotchi is no longer alive.\n", t.Name)
			break
		}
	}
}

func main() {
	tamagotchi := NewTamagotchi("Tama")
	go tamagotchi.Live() // Start the Tamagotchi lifecycle

	for {
		var choice int
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1. Feed")
		fmt.Println("2. Play")
		fmt.Println("3. Check Status")
		fmt.Println("4. Quit")

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			tamagotchi.Feed()
		case 2:
			tamagotchi.Play()
		case 3:
			tamagotchi.CheckStatus()
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
}

