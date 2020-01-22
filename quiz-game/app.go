package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func quiz(problems [][]string, timeout int, shuffle bool) int {
	total := 0
	timer := time.NewTimer(time.Second * time.Duration(timeout))

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })

	for i, j := range problems {
		q, a := j[0], j[1]
		fmt.Printf("Q%d. What is %s? ", i, q)
		answerChannel := make(chan string)
		go func() {
			var userAnswer string
			fmt.Scanln(&userAnswer)
			answerChannel <- strings.Trim(userAnswer, " ")
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
	var shuffle bool
	flag.StringVar(&problemsFile, "path", "problems/problems.csv", "Problems CSV file")
	flag.IntVar(&timeout, "timeout", 10, "Max time for the quiz")
	flag.BoolVar(&shuffle, "shuffle", true, "Question shffling")
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

	total := quiz(problems, timeout, shuffle)

	fmt.Printf("You scored %d out of %d!\n", total, len(problems))
}
