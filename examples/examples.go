package examples

import (
	"github.com/WelscherMarkus/Go-Sample-Data"
	"log"
)

// Generate a new slice of structs and fill them with sample data
func fillStructs() {
	type User struct {
		IrrelevantField string
		FirstName       string `sample:"first_name"`
		LastName        string `sample:"last_name"`
		DateOfBirth     string `sample:"date_of_birth"`
		Date            string `sample:"date"`
	}

	structs, err := sample.GenerateNewStructs(User{}, 10)
	if err != nil {
		log.Fatalln(err)
	}

	users, ok := structs.([]User)
	if !ok {
		log.Fatal("expected generated structs to be of type []User")
	}

	for _, user := range users {
		log.Println(user)
	}
}

// Existing structs and fill specific fields with sample data
func fillSpecificFields() {
	type User struct {
		IrrelevantField string
		FirstName       string `sample:"first_name"`
		LastName        string `sample:"last_name"`
		DateOfBirth     string `sample:"date_of_birth"`
		Date            string `sample:"date"`
	}

	var users = []User{
		{
			IrrelevantField: "irrelevant",
			FirstName:       "first_name",
			LastName:        "last_name",
			DateOfBirth:     "28.02.2024",
			Date:            "01.01.1970",
		},
		{
			IrrelevantField: "irrelevant",
			FirstName:       "first_name",
			LastName:        "last_name",
			DateOfBirth:     "28.02.2024",
			Date:            "01.01.1970",
		},
	}

	_, err := sample.FillStructs(&users)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		log.Println(user)
	}
}
