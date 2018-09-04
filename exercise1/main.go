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
	"time"
)

func main() {
	filename := flag.String("filename", "problems.csv", "path to problems csv file")
	limit := flag.Int("limit", 30, "time limit in seconds to complete the quiz")
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
	timer := time.NewTimer(time.Duration(*limit) * time.Second)

problemLoop:
	for _, p := range problems {
		fmt.Printf("%s ?  >> ", p.Question)
		answerCh := make(chan string)
		go func() {
			scanner.Scan()
			answerCh <- scanner.Text()
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == p.Answer {
				correct++
			}
		}

	}

	// output the results at the end
	fmt.Printf("You answered %d out of %d questions right. Thanks for playing.\n\n", correct, len(problems))
}
