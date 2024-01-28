package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("cfile.csv")
	if err != nil {
		fmt.Println("There is error to open csv file")
		log.Fatal(err)
	}
	defer file.Close()

	// read csv file 
	r := csv.NewReader(file)
	// define the comment or comma
	r.Comment = '#'

	// read line by line
	for {
		record, err := r.Read()
		// end of document
		if err == io.EOF {
			break	
		}
		if err != nil {
			if pe, ok := err.(*csv.ParseError);ok {
				fmt.Println("bad column:", pe.Column)
				fmt.Println("bad line:", pe.Line)
				fmt.Println("Error reported", pe.Err)
				if pe.Err == csv.ErrFieldCount {
					continue
				}
			}
		
			log.Fatal(err)
		}
		fmt.Println("CSV Row", record)

		i, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println( i * 4)
	}
}