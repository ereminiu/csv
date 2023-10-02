package main

import (
	"log"

	"github.com/ereminiu/vlr/writer"
)

func main() {
	// r := [][]string{{"Ivan", "18"}, {"Katya", "26"}}
	r := [][]string{{"Valera", "228"}, {"Tambov", "47"}}
	err := writer.Write(r)
	if err != nil {
		log.Fatal(err)
	}
}
