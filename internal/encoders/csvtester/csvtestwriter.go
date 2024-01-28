package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"Jaro", "5", "ALA, IOI"},
		{"Mala", "4", "ABD, B00"},
		{"Kay", "3", "HBJ, D3N"},
	}

	file, err := os.Create("cfilew.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.Comma = ';'

	// w.WriteAll(records)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}
	// write data to disk
	w.Flush()

	err = w.Error()
	if err != nil {
		log.Fatal(err)
	}

}