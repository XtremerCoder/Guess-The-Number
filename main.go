package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	// Generate a random number between 100 and 999 (inclusive)
	randomNumber := rand.Intn(999)

	// Convert the random number to a string and get the first digit
	randomNumberStr := strconv.Itoa(randomNumber)
	firstDigit := string(randomNumberStr[0])

	// Prompt the user for input
	var userInput string
	fmt.Print("Enter a single digit: ")
	fmt.Scanln(&userInput)

	// Compare the user's input with the first digit of the random number
	if userInput == firstDigit {
		var userInput2 string
		fmt.Print("Enter a single digit: ")
		fmt.Scanln(&userInput2)
		// Ensure the number has at least two digits
		if len(randomNumberStr) == 1 {
			fmt.Println("Generated number is too small.")
			return
		}

		// Extract the first and second digits
		secondDigit := string(randomNumberStr[1])

		if userInput2 == secondDigit {
			fmt.Println("Match! Your input matches the second digit of the random number.")

			if randomNumber > 99 {
				var userInput3 string
				fmt.Print("Enter a single digit: ")
				fmt.Scanln(&userInput3)

				ThirdDigit := string(randomNumberStr[2])

				if userInput3 == ThirdDigit {
					fmt.Println("Match! Your input matches the second digit of the random number.")
				} else {
					fmt.Println("No match. Your input does not match the first or second digit of the random number.")
				}
			} else {
                fmt.Println("Number: ", randomNumber)
				fmt.Println("Wrong. Exiting...")
			}

		} else {
			fmt.Println("No match. Your input does not match the first or second digit of the random number.")
		}
	} else {
fmt.Println("Number: ", randomNumber)
		fmt.Println("Wrong. Exiting...")
	}

}
