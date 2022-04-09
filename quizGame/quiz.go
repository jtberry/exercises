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
	//"sync"
	"time"
)

// create the struct for the Question
type Question struct {
	prob   string
	answer string
}

// create the variable for the flags
var (
	csvLoad *string
	qtimer  *int
)


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

func quizTheHuman(csvLoad string, qtimer int) {
	// the answer we are expecting from the user
	var answer string
	// a counter for how many correct answers we capture
	correct := 0
	// initialize the timer and the value we provide, this creates a timer.C channel
	timer := time.NewTimer(time.Second * time.Duration(qtimer))
	//now do something with the data once read the csv
	lines,err := readCSV(csvLoad)
	if err != nil {
		panic(err)
	}
	totalQue := len(lines)
	//loop through the lines and turn them into objects
	for i,line := range lines {
		question := Question{
			prob: line[0],
			answer: line[1],
		}

		// creating a channel for the answer to provide the data
		answerCh := make(chan string)
		// giving question at the beginning of the loop so that the timer can exit
		// no matter when timer goes off
		fmt.Printf("Problem #%v: %v\n",i+1, question.prob)
		fmt.Println("Answer: ")

		// create a go func to scan for the answer and the provide that answer to the channel
		// this will allow the timer to interupt if waiting
		go func() {
			fmt.Scan(&answer)
			// return i to the answer channel
			answerCh <- answer
		}()
		// we put the timer in the loop with a select

		select {
		// once the timer in the background has been reached  we return and exit with message
		case <-timer.C:
			fmt.Println("\ntime's up!")
			totalScore(correct,totalQue)
			return
		// while the timer is counting we keep looping through the code
		// and return the answerCh to i and evaluate
		case answer := <-answerCh:
			if answer == question.answer {
				correct++
			}else if answer == "exit" {
				fmt.Println("Exiting the Game...")
				totalScore(correct,totalQue)
				return
			}
		}
	}
	totalScore(correct,totalQue)
}

// creating a function to evaluate the answers
func totalScore(correct ,totalQue int) {
	fmt.Printf("You got %v right out of %v\n",correct,totalQue)
}


func init() {
	csvLoad = flag.String("f", "problems.csv", "this is the default file to load")
	qtimer = flag.Int("t", 30, "default timer for the questions")
}

func main() {
	flag.Parse()

	fmt.Println("Hell and Welcome to the quiz game!")
	fmt.Printf("Reading from file: %v\n",*csvLoad)
	fmt.Println("Please provide the answer to the provided equation")
	
	quizTheHuman(*csvLoad,*qtimer)
	
}