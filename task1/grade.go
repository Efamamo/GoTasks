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
	amount, _ := strconv.Atoi(scanner.Text())

	subjects := map[string]float64{}
	for i := 0; i < amount; i++ {

		fmt.Print("Enter name of the subject ", i+1, ": ")
		scanner.Scan()
		name := scanner.Text()
		fmt.Printf("Enter score of %v: ", name)
		scanner.Scan()
		score, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println(err)
		}
		subject := scanner.Text()
		strings.Split(" ", subject)

		subjects[name] = score

	}

	fmt.Println("----------------------------")

	sum := 0.0
	fmt.Printf("Name: %v\n", name)
	for key, val := range subjects {
		fmt.Printf("%v : %v\n", key, val)
		sum += val
	}

	average := sum / float64(amount)
	if amount == 0 {
		fmt.Println("Your Dont have any Subjects")
	} else {
		fmt.Printf("Your Average Grade is %v/100\n", average)
	}
	fmt.Println("----------------------------")
}
