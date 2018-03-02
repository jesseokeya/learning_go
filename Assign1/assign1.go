package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("-----------------------------------------------")
	fmt.Println(" Welcome to the CST8101 Final Mark Calculator ")
	fmt.Println("-----------------------------------------------")

	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your Lab Assignments mark out of 25: ")
	lab, labErr := scanner.ReadString('\n')
	HandleErr(labErr)

	fmt.Print("Enter your Hybrid Activities / Quizzes mark out of 20: ")
	quiz, quizErr := scanner.ReadString('\n')
	HandleErr(quizErr)

	fmt.Print("Enter your Test mark out of 25: ")
	test, testErr := scanner.ReadString('\n')
	HandleErr(testErr)

	fmt.Print("Enter your Final Assessment (Exam) mark out of 30: ")
	final, finalErr := scanner.ReadString('\n')
	HandleErr(finalErr)

	theoryGrade := ((double(test) + double(final)) / 55) * 100
	practicalGrade := ((double(lab) + double(quiz)) / 45) * 100
	finalGrade := double(lab) + double(quiz) + double(final) + double(test)

	newFinalGrade := format(finalGrade)
	newPracticeGrade := format(practicalGrade)
	newTheoryGrade := format(theoryGrade)

	if strings.ContainsAny(newFinalGrade, ".00") {
		newFinalGrade = strings.Replace(newFinalGrade, ".00", "", -1)
	}

	if strings.ContainsAny(newPracticeGrade, ".00") {
		newPracticeGrade = strings.Replace(newPracticeGrade, ".00", "", -1)
	}

	if strings.ContainsAny(newTheoryGrade, ".00") {
		newTheoryGrade = strings.Replace(newTheoryGrade, ".00", "", -1)
	}

	fmt.Println("-----------------------------------------------")
	fmt.Println("Theory grade: " + newTheoryGrade + "%")
	fmt.Println("Practical grade: " + newPracticeGrade + "%")
	fmt.Println("Final grade: " + newFinalGrade + "%")
	fmt.Println("-----------------------------------------------")

}

// HandleErr -- this comment is required by golint
func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func double(s string) float64 {
	s = strings.Replace(s, "\n", "", -1)
	_, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Invalid Input. Try Again 🙃")
		os.Exit(1)
	}
	result, err := strconv.ParseFloat(s, 64)
	HandleErr(err)
	return result
}

func format(f float64) string {
	result := strconv.FormatFloat(f, 'f', 2, 64)
	return result
}
