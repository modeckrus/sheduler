package main

import (
	"log"
	"time"

	scheduler "github.com/modeckrus/sheduler"

	"github.com/modeckrus/sheduler/storage"
)

type Message struct {
	Str string
}

func TaskWithoutArgs() {
	log.Println("TaskWithoutArgs is executed")
}

func TaskWithArgs(message string) {
	log.Println("TaskWithArgs is executed. message:", message)
}

func main() {
	storage := storage.NewMongoDBStorage(
		storage.MongoDBConfig{
			ConnectionUrl: "mongodb://localhost:27017/?readPreference=primary&ssl=false&connect=direct",
			Db:            "example",
		},
	)

	if err := storage.Connect(); err != nil {
		log.Fatal("Could not connect to db", err)
	}

	if err := storage.Initialize(); err != nil {
		log.Fatal("Could not intialize database", err)
	}

	s := scheduler.New(storage)
	t := time.Now().Add(time.Second * 2)
	s.RunAt(t, TaskWithoutArgs)
	s.RunAt(t, TaskWithArgs, "Hello from at")
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Wait()
}
