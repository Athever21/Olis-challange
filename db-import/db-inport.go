package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"olist/models"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		exit(fmt.Sprintf("Usage: %s <filename>", os.Args[0]))
	}

	file, err := os.Open(args[0])
	if err != nil {
		exit("Could not open a file: " + args[0])
	}

	r := csv.NewReader(file)

	valueString := ""
	count := 0

	line, err := r.Read()

	if err != nil || line[0] == "name" {
		exit("Couldn't parse csv file")
	}

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			exit(err.Error())
		}

		valueString += fmt.Sprintf("(\"%s\"),", line[0])
		count++
	}

	err = models.SaveAuthors(valueString[:len(valueString)-1])

	if err != nil {
		exit(err.Error())
	}

	fmt.Printf("Successfully imported %d authors\n", count)
}

func exit(name string) {
	fmt.Println(name)
	os.Exit(1)
}
