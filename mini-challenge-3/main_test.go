package main

import (
	"fmt"
	"testing"
)

func TestFindPerson(t *testing.T) {
	var people []Person

	people = append(people, Person{ID: 0, name: "Fitri", address: "Jl. Lorem", occupation: "Backend", reason: "Alasan Fitri"})
	people = append(people, Person{ID: 1, name: "Budi", address: "Jl. Ipsum", occupation: "Frontend", reason: "Alasan Budi"})
	people = append(people, Person{ID: 2, name: "Anggraini", address: "Jl. Dolor", occupation: "Fullstack", reason: "Alasan Anggraini"})
	people = append(people, Person{ID: 3, name: "Cahya", address: "Jl. Sit", occupation: "DevOps", reason: "Alasan Cahya"})
	people = append(people, Person{ID: 4, name: "Dewi", address: "Jl. Amet", occupation: "QA", reason: "Alasan Dewi"})
	people = append(people, Person{ID: 5, name: "Eko", address: "Jl. Consectetur", occupation: "Backend", reason: "Alasan Eko"})

	tests := []struct {
		input    string
		expected bool
	}{
		{"Fitri", true},
		{"Vito", false},
		{"eko", true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("FindPerson(%s)", test.input), func(t *testing.T) {
			person := findPerson(people, test.input)
			if test.expected && person == nil {
				t.Errorf("Expected person to be found, but got nil")
			}
			if !test.expected && person != nil {
				t.Errorf("Expected person not to be found, but got %v", person)
			}
		})
	}
}
