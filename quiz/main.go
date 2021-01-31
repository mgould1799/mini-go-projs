package main

import (
	"strings"
	"encoding/csv"
	"fmt"
	"os"
	"flag"
)


func main() {

	// the flag for the csv file
	csvFileName := flag.String("csv", "problems.csv",
	 "a csv file in the format of question, answer")
	// parse those flags 
	// flag will actually tell you those commands if you do -h in the cli
	flag.Parse()

	// open the file
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit("Failed to open CSV file: " + *csvFileName)
	}

	// read the file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse provided csv file")
	}

	// pass this to a function to parse the lines from the file
	// comes back as 2d slice
	problems := parseLines(lines)
	
	// keep count of total correct answers
	correctAnswers := 0
	// loop through the questions
	for i, problem := range problems{
		// print the question for the user to see
		fmt.Printf("Problem #%d: %s = \n", i + 1, problem.Question)
		// read in the answer from the user
		var answer string
		fmt.Scanf("%s\n", &answer)
		//check if the answer is correct and increment the correctAnswer var
		if answer == problem.Answer{
			correctAnswers ++
		}
	}

	// print out the results
	fmt.Printf("You scored %d out of %d\n", correctAnswers, len(problems))

}


// parse lines is taking in a 2d array to put in a slice of Problems
func parseLines(lines [][]string) []Problem {
	// create a slice of problems with a defined length
	ret := make([]Problem, len(lines))

	for key, value := range lines {
		ret[key] = Problem{
			Question: strings.TrimSpace(value[0]),
			Answer: strings.TrimSpace(value[1]),
		}
	}
	return ret
}

// Problem struct takes in an answer and question
type Problem struct {
	Question string 
	Answer string 
}


// a function that takes in a message to print to standout and exit on failure
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}