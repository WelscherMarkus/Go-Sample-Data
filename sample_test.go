package sample

import (
	"testing"
)

func TestFillStructs(t *testing.T) {
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

	_, err := FillStructs(&users)
	if err != nil {
		t.Fatal(err)
	}

	for _, user := range users {
		if user.FirstName == "first_name" {
			t.Fatal("expected first name to be filled")
		}
		if user.LastName == "last_name" {
			t.Fatal("expected last name to be filled")
		}
		if user.IrrelevantField != "irrelevant" {
			t.Fatal("expected irrelevant field to be unchanged")
		}
		if user.DateOfBirth == "28.02.2024" {
			t.Fatal("expected date of birth to be filled")
		}
		if user.Date == "01.01.1970" {
			t.Fatal("expected date to be filled")
		}
	}

	t.Log("FillStructs() successfully filled the existing struct fields")
}

func TestGenerateNewStructs(t *testing.T) {
	type User struct {
		IrrelevantField string
		FirstName       string `sample:"first_name"`
		LastName        string `sample:"last_name"`
		DateOfBirth     string `sample:"date_of_birth"`
		Date            string `sample:"date"`
	}

	structs, err := GenerateNewStructs(User{}, 10)
	if err != nil {
		t.Fatal(err)
	}

	users, ok := structs.([]User)
	if !ok {
		t.Fatal("expected generated structs to be of type []User")
	}

	for _, user := range users {
		if user.FirstName == "" {
			t.Fatal("expected first name to be filled")
		}
		if user.LastName == "" {
			t.Fatal("expected last name to be filled")
		}
		if user.IrrelevantField != "" {
			t.Fatal("expected irrelevant field to be empty")
		}
		if user.DateOfBirth == "" {
			t.Fatal("expected date of birth to be filled")
		}
		if user.Date == "" {
			t.Fatal("expected date to be filled")
		}
	}

	if len(users) != 10 {
		t.Fatalf("expected 10 users, got %d", len(users))
	}

	t.Log("GenerateNewStructs() successfully generated 10 new structs")
}
