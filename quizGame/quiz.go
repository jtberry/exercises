package main

// read a csv file and run with a timer
// flags package, os and CSV, channels and go routine timer package
// https://courses.calhoun.io/lessons/les_goph_01

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

// create the struct for the Question
type Question struct {
	prob   string
	answer string
}

// create the variable for the flags
var csvLoad *string


// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func readCSV(csvfile string) ([][]string, error) {
	//open the file
	f, err := os.Open(csvfile)
	// error if unable to opent the file
	if err != nil {
		log.Fatalf("could not open file: %v\n",err)
		os.Exit(1)
		return [][]string{},err
	}
	defer f.Close()

	// create the reader
	csvRead := csv.NewReader(f)

	// gather all data as records and stor in an array
	records, err := csvRead.ReadAll()
	if err != nil {
		log.Fatalf("could not parse .csv: %v\n",err)
		os.Exit(1)
		return [][]string{}, err
	}
	return records, nil
}

// func parseRecords(lines [][]string) []Question {
// 	ret := make([]question, len(records))
// 	for i,line := range lines {
// 		ret[i] := Question{
// 			prob: line[0],
// 			answer: line[1],
// 		}
// 	return ret
// }

func quizTheHuman(csvLoad string) {
	var i string
	correct := 0
	wrong := 0
	//now do something with the data once read the csv
	lines,err := readCSV(csvLoad)
	if err != nil {
		panic(err)
	}
	//loop through the lines and turn them into objects
	for _,line := range lines {
		question := Question{
			prob: line[0],
			answer: line[1],
		}
	//problems := parseRecords(lines)
		fmt.Println(question.prob)
		fmt.Print("Please provide the answer: ")
		fmt.Scanf(&i)
		if i == question.answer {
			correct++
		} else if i == "exit" {
			fmt.Println("Exiting the Game...")
			os.Exit(1)
		} else {
			wrong++
	}
	}
	fmt.Printf("You got %v right out of %v\n",correct,len(lines))
}

func init() {
	csvLoad = flag.String("f", "problems.csv", "this is the default file to load")
}

func main() {
	flag.Parse()

	fmt.Println("Hell and Welcome to the quiz game!")
	fmt.Printf("Reading from file: %v\n",*csvLoad)
	fmt.Println("Please provide the answer to the provided equation")
	
	quizTheHuman(*csvLoad)
	
}