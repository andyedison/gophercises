package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filename := flag.String("filename", "problems.csv", "path to problems csv file")
	flag.Parse()

	// read a csv file from command line, default to problems.csv -filename flag
	content, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewReader(content))

	input, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// parse csv into struct of question and answer string
	type Problem struct {
		Question string
		Answer   string
	}

	var problems []Problem
	problems = make([]Problem, len(input))
	for idx, line := range input {
		if len(line) != 2 {
			log.Fatal("invalid format of problems csv file")
		}
		var prob Problem
		prob.Question = line[0]
		prob.Answer = line[1]
		problems[idx] = prob
	}

	// read out each question, then read in answer. keep track of how many are correct and wrong
	scanner := bufio.NewScanner(os.Stdin) // define our scanner to read stdin
	var correct int
	for _, p := range problems {
		fmt.Printf("%s ?  >> ", p.Question)
		scanner.Scan()
		a := scanner.Text()
		if a == p.Answer {
			correct++
		}
	}

	// output the results at the end
	fmt.Println("")
	fmt.Printf("You answered %d out of %d questions right. Thanks for playing.\n\n", correct, len(problems))
}
