package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func quiz(problems [][]string, timeout int) int {
	total := 0
	timer := time.NewTimer(time.Second * time.Duration(timeout))

	for i, j := range problems {
		q, a := j[0], j[1]
		fmt.Printf("Q%d. What is %s? ", i, q)
		answerChannel := make(chan string)
		go func() {
			var userAnswer string
			fmt.Scanln(&userAnswer)
			answerChannel <- userAnswer
		}()
		select {
		case userAnswer := <-answerChannel:
			if userAnswer == a {
				total++
			}
		case <-timer.C:
			fmt.Println("TIMEOUT!")
			return total
		}
	}

	return total
}

func main() {
	var timeout int
	var problemsFile string
	flag.StringVar(&problemsFile, "path", "problems/problems.csv", "Problems CSV file")
	flag.IntVar(&timeout, "timeout", 10, "Max time for the quiz")
	flag.Parse()

	csvfile, err := os.Open(problemsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	problems, err := csv.NewReader(csvfile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	total := quiz(problems, timeout)

	fmt.Printf("You scored %d out of %d!\n", total, len(problems))
}
