package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

func main() {
	//reading input from user
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	// loop to the user quit the program
	for scanner.Scan() {
		input = scanner.Text()
		
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" || input == "konvert" {
			fmt.Println("Konverterer all maalingene gitt i grader Celsius til grader Fahrenheit")
			// trying opening fil
			fileConv, errConv := os.OpenFile("kjevik-temp-celsius-20220318-20230318", os.O_RDWR|os.O_CREATE, 0755)
			scannerConvFile := bufio.NewScanner(fileConv)
			
			// Check if the fil is empty
			if len(scannerConvFile.Text()) == 0) {

				for scannerConvFile.Scan() {
					fmt.Println("for loop")
					line := scanner.Text()
					fmt.Println(line)
				}
			} else {
				fmt.Println("Filen finnes allerede. Vil du generere fil paa nytt?")
				_ := bufio.NewScanner(os.Stdin)
				
				//for scannerConv.Scan() {
				//	if scannerConv.Text() == "j"|| scannerconv.Text() == "ja"  {
				//		fmt.Println("La oss skrive over den filen!")
				//	} else if scannerConv.Text() == "n" || scannerConv.Text() == "nei" {
				//		break
				//	}
				//}
			}
			
			//check if there errors
			if errConv != nil {
				log.Fatal(errConv)
			}

		} else {
			fmt.Println("Ikke gylding alternativ, venligst prove igjen")
		}
	}
}
