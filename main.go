package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"github.com/andrineCF/minyr/yr"
)

func main() {
	// Open the main fil
	filCel, err :=  os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	scannerCel := bufio.NewScanner(filCel)

	// closing the fil
	defer filCel.Close()

	//reading input from user
	var input string
	scannerInput := bufio.NewScanner(os.Stdin)

	// loop to the user quit the program
	for scannerInput.Scan() {
		input = scannerInput.Text()
		
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" || input == "konvert" {
			fmt.Println("Konverterer all maalingene gitt i grader Celsius til grader Fahrenheit")
			// trying opening fil
			fileConv, errConv := os.OpenFile("kjevik-temp-celsius-20220318-20230318.csv", os.O_RDWR|os.O_CREATE, 0755)
			scannerConvFile := bufio.NewScanner(fileConv)
			
			// Check if the fil is empty
			if len(scannerConvFile.Text()) == 0 {
				for scannerCel.Scan() {
					fmt.Println(yr.FileCelsiusToFhrenheit("10"))
				}
			} else {
				fmt.Println("Filen finnes allerede. Vil du generere fil paa nytt?")
				scannerConv := bufio.NewScanner(os.Stdin)
				
				for scannerConv.Scan() {
					if scannerConv.Text() == "j"|| scannerConv.Text() == "ja"  {
						fmt.Println("La oss skrive over den filen!")
					} else if scannerConv.Text() == "n" || scannerConv.Text() == "nei" {
						break
					}
				}
			}
			
			//check if there errors
			if errConv != nil {
				log.Fatal(errConv)
			}

			//Closed the file
			fileConv.Close()
		} else {
			fmt.Println("Ikke gylding alternativ, venligst prove igjen")
		}
	}

	if err == nil {
		log.Fatal(err)
	}
}
