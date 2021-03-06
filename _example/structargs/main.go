package main

import (
	"fmt"
	"log"
	"time"

	scheduler "github.com/modeckrus/sheduler"

	"github.com/modeckrus/sheduler/storage"
)

type Gender int

const DateLayout = "2006-01-02"
const (
	Male Gender = iota
	Female
)

type Person struct {
	ID          string
	Name        string
	DateOfBirth time.Time
	Gender      Gender
}

func CheckIfBirthday(person Person) {
	if time.Now().Format(DateLayout) == person.DateOfBirth.Format(DateLayout) {
		fmt.Println("Happy birthday,", person.Name)
		return
	}
	fmt.Println("Still waiting for your birthday")
}

func main() {
	st := storage.NewMongoDBStorage(
		storage.MongoDBConfig{
			ConnectionUrl: "mongodb://localhost:27017/?readPreference=primary&ssl=false&connect=direct",
			Db:            "example",
		},
	)
	if err := st.Connect(); err != nil {
		log.Fatal("Could not connect to db", err)
	}

	if err := st.Initialize(); err != nil {
		log.Fatal("Could not intialize database", err)
	}

	s := scheduler.New(st)

	dob, _ := time.Parse(DateLayout, time.Now().Format(DateLayout))
	person := Person{
		ID:          "123-456",
		Name:        "John Smith 2",
		DateOfBirth: dob,
		Gender:      Male,
	}

	// Start a task with arguments
	if _, err := s.RunEvery(5*time.Second, CheckIfBirthday, person); err != nil {
		log.Fatal(err)
	}

	s.Start()
	s.Wait()
}
