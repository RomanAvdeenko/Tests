package main

import (
	"convar/dataStructure"
	"log"
	"os"
)

func main() {
	f, err := os.Create("system.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f1, err := os.Create("mail.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	rec := dataStructure.NewRecordWriter(f, f1)

	rec.Start()
	rec.Prompt()
}
