package task1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GradeCalculator() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Your Name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter Number of Subjects: ")
	scanner.Scan()
	amount, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	subjects := map[string]float64{}
	var i int64
	for i < amount {

		fmt.Print("Enter name of the subject ", i+1, ": ")
		scanner.Scan()
		name := scanner.Text()
		fmt.Printf("Enter score of %v: ", name)
		scanner.Scan()
		score, _ := strconv.ParseFloat(scanner.Text(), 64)
		subject := scanner.Text()
		strings.Split(" ", subject)

		subjects[name] = score

		i++

	}

	fmt.Println("----------------------------")

	sum := 0.0
	fmt.Printf("Name: %v\n", name)
	for key, val := range subjects {
		fmt.Printf("%v : %v\n", key, val)
		sum += val
	}

	average := sum / float64(amount)
	fmt.Printf("Your Average Grade is %v/100\n", average)
	fmt.Println("----------------------------")
}
