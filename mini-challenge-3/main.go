package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	ID         int
	name       string
	address    string
	occupation string
	reason     string
}

func main() {
	var people []Person

	people = append(people, Person{ID: 0, name: "Fitri", address: "Jl. Lorem", occupation: "Backend", reason: "Alasan Fitri"})
	people = append(people, Person{ID: 1, name: "Budi", address: "Jl. Ipsum", occupation: "Frontend", reason: "Alasan Budi"})
	people = append(people, Person{ID: 2, name: "Anggraini", address: "Jl. Dolor", occupation: "Fullstack", reason: "Alasan Anggraini"})
	people = append(people, Person{ID: 3, name: "Cahya", address: "Jl. Sit", occupation: "DevOps", reason: "Alasan Cahya"})
	people = append(people, Person{ID: 4, name: "Dewi", address: "Jl. Amet", occupation: "QA", reason: "Alasan Dewi"})
	people = append(people, Person{ID: 5, name: "Eko", address: "Jl. Consectetur", occupation: "Backend", reason: "Alasan Eko"})

	fmt.Print("Enter a name or id number: ")
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return
	}

	value = strings.TrimSpace(value)

	person := findPerson(people, value)

	if person == nil {
		fmt.Println("Person not found")
	} else {
		fmt.Printf("ID: %d\nName: %s\nAddress: %s\nOccupation: %s\nReason: %s\n", person.ID, person.name, person.address, person.occupation, person.reason)
	}

}

func findPerson(people []Person, query string) *Person {
	queryAsName := strings.ToLower(query)

	queryAsId, err := strconv.Atoi(query)
	if err != nil {
		queryAsId = -1
	}

	var result *Person
	for _, person := range people {
		if (strings.ToLower(person.name) == queryAsName) || queryAsId >= 0 && person.ID == queryAsId {
			result = &person
			break
		}
	}

	return result
}
