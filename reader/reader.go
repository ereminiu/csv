package reader

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

const filepath = "input.csv"

func FetchRecords() [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	res := make([][]string, 0)
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		res = append(res, rec)
	}
	return res
}
