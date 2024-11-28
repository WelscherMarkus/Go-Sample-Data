package sample

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"reflect"
	"time"
)

var firstNames []string
var lastNames []string

const tagName = "sample"

func FillStructs(input interface{}) (interface{}, error) {
	value := reflect.ValueOf(input)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Slice {
		return nil, fmt.Errorf("input must be a slice")
	}

	for j := 0; j < value.Len(); j++ {
		element := value.Index(j)

		for i := 0; i < element.NumField(); i++ {
			tag := element.Type().Field(i).Tag.Get(tagName)
			if tag == "" {
				continue
			}

			data, err := generateData(tag)
			if err != nil {
				return nil, err
			}
			if data != "" {
				element.Field(i).SetString(data)
			}
		}
	}

	return nil, nil
}

func GenerateNewStructs(input interface{}, numOfElements int) (interface{}, error) {
	elemType := reflect.TypeOf(input)

	returnData := reflect.MakeSlice(reflect.SliceOf(elemType), numOfElements, numOfElements)

	for j := 0; j < numOfElements; j++ {
		newInstance := reflect.New(elemType).Elem()

		for i := 0; i < elemType.NumField(); i++ {
			tag := elemType.Field(i).Tag.Get(tagName)

			if tag == "" {
				continue
			}

			data, err := generateData(tag)
			if err != nil {
				return nil, err
			}
			newInstance.Field(i).SetString(data)
		}
		returnData.Index(j).Set(newInstance)
	}
	return returnData.Interface(), nil
}

func generateData(category string) (string, error) {
	switch category {
	case "first_name":
		if len(firstNames) == 0 {
			loadFirstNames()
		}

		randNumber := rand.IntN(len(firstNames))
		return firstNames[randNumber], nil

	case "last_name":
		if len(lastNames) == 0 {
			loadLastNames()
		}

		randNumber := rand.IntN(len(lastNames))
		return lastNames[randNumber], nil

	case "date_of_birth":
		dateOfBirth := generateDateOfBirth()
		return dateOfBirth.Format("2006-01-02"), nil

	case "date":
		startOfYear := time.Date(1970, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
		endOfYear := time.Date(2021, time.Month(12), 31, 23, 59, 59, 0, time.UTC)
		amountOfDaysBetween := endOfYear.Sub(startOfYear).Hours() / 24
		randomDays := rand.IntN(int(amountOfDaysBetween))
		randomDate := startOfYear.AddDate(0, 0, randomDays)
		return randomDate.Format("01-02"), nil

	default:
		log.Println("unknown category: ", category)
		return "", fmt.Errorf("unknown category: %s", category)
	}
}

func generateDateOfBirth() time.Time {
	// Generate a random date of birth between 10 and 70 years ago
	now := time.Now()
	start := now.AddDate(-70, 0, 0)
	end := now.AddDate(-10, 0, 0)

	amountOfDaysBetween := end.Sub(start).Hours() / 24
	randomDays := rand.IntN(int(amountOfDaysBetween))

	randomDate := start.AddDate(0, 0, randomDays)
	return randomDate
}

func loadFirstNames() {
	file, err := os.Open("pkg/firstNames.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scammer := bufio.NewScanner(file)
	for scammer.Scan() {
		firstNames = append(firstNames, scammer.Text())
	}
}

func loadLastNames() {
	file, err := os.Open("pkg/lastNames.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scammer := bufio.NewScanner(file)
	for scammer.Scan() {
		lastNames = append(lastNames, scammer.Text())
	}
}
