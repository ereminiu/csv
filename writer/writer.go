package writer

import (
	"encoding/csv"
	"os"
)

const filepath = "output.csv"

func Write(records [][]string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()
	for _, r := range records {
		if err := writer.Write(r); err != nil {
			return err
		}
	}
	return nil
}
