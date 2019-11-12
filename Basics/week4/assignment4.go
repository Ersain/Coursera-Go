package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locoMethod string
	sound      string
}

func (c *Cow) Eat() {
	if c == nil {
		return
	}
	fmt.Println(c.food)
}
func (c *Cow) Move() {
	if c == nil {
		return
	}
	fmt.Println(c.locoMethod)
}
func (c *Cow) Speak() {
	if c == nil {
		return
	}
	fmt.Println(c.sound)
}

type Bird struct {
	food       string
	locoMethod string
	sound      string
}

func (b *Bird) Eat() {
	if b == nil {
		return
	}
	fmt.Println(b.food)
}
func (b *Bird) Move() {
	if b == nil {
		return
	}
	fmt.Println(b.locoMethod)
}
func (b *Bird) Speak() {
	if b == nil {
		return
	}
	fmt.Println(b.sound)
}

type Snake struct {
	food       string
	locoMethod string
	sound      string
}

func (s *Snake) Eat() {
	if s == nil {
		return
	}
	fmt.Println(s.food)
}
func (s *Snake) Move() {
	if s == nil {
		return
	}
	fmt.Println(s.locoMethod)
}
func (s *Snake) Speak() {
	if s == nil {
		return
	}
	fmt.Println(s.sound)
}

func readInput(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return text
}

var cow *Cow = &Cow{food: "grass", locoMethod: "walk", sound: "moo"}
var bird *Bird = &Bird{food: "worms", locoMethod: "fly", sound: "peep"}
var snake *Snake = &Snake{food: "mice", locoMethod: "slither", sound: "hsss"}


func createAnimal(aType string) Animal {
	var animal Animal
	

	switch aType {
	case "cow":
		return cow
	case "snake":
		return snake
	case "bird":
		return bird
	default:
		fmt.Println("Wrong animal type")
	}
	return animal
}

func initData() map[string]Animal {
	return map[string]Animal{
		"cow1":   createAnimal("cow"),
		"snake1": createAnimal("snake"),
		"bird1":  createAnimal("bird"),
	}
}

func callAction(action string , animal Animal) {
     if action == "eat" {
	animal.Eat()
     } else if action == "move" {
	animal.Move()
     }else if action == "speak" {
        animal.Speak()
     }
}

func showInfo(action string, animal Animal ) {

	switch animal.(type) {
	case *Snake:
	     callAction(action, animal)		
	case *Cow:
	     callAction(action, animal)
	case *Bird:
	     callAction(action, animal)
	default:
		fmt.Println("Wrong type")

	}
}

func dispatchRequest(input string) {
	fields := strings.Fields(input)
	if len(fields) < 3 {
		fmt.Println("Not enough information to request")
		return
	}

	action := fields[0]
	aName := fields[1]

	if action == "newanimal" {
		aType := fields[2]
		data[aName] = createAnimal(aType)
		fmt.Println("Created It!")
	} else if action == "query" {
	 	infoType := fields[2]
		v, found := data[aName];
		if !found {
			fmt.Println("animal not found")
			return
		}
		showInfo(infoType, v)
	}

}

var data map[string]Animal

func main() {
	reader := bufio.NewReader(os.Stdin)
	data = initData()
	var str string

	for {
		fmt.Print(">")
		str = readInput(reader)
		dispatchRequest(str)
	}


}
