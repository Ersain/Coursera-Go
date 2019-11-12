package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//	"bytes"
	//	"fmt"
	//	"io/ioutil"
)

type Persons []*Person

type Person struct {
	fname string
	lname string
}

func (p Persons) print() {
	for _, pers := range p {
		fmt.Printf("Name: %s, Lastname: %s\n",pers.fname, pers.lname)
	}
}

func createPerson(str string) *Person {
	var ppers = new(Person)

	str = strings.TrimSpace(str)
	fields := strings.Fields(str)

	if len(fields) > 1 {
		ppers = &Person{fname: fields[0], lname: fields[1]}
	} else {
		return nil
	}

	return ppers

}

func main() {
	var  fileName string
        fmt.Println("Please enter file name:")
        fmt.Scanf("%s\n", &fileName)
	fileName = strings.TrimSpace(fileName)
        
	file, err := os.Open(fileName)
	if err != nil {
		panic("File error")
	}
	defer file.Close()

	lst := make([]*Person, 0, 100)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		if str == ""{
			continue
		}
		lst = append(lst, createPerson(str))
	}

	persons := Persons(lst)

	persons.print()
}
