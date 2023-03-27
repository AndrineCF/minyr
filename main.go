package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	//"github.com/AndrineCF/minyr/yr"
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
				
				var readLines []string

				//keep track of the position when reading the file
				pos := 0

				//read throught the celsius file
				for scannerCel.Scan() {
					lines := strings.Split(scannerCel.Text(), ";")
					
					//check if at the  start or the end of file
					if pos != 0 && lines[3] != "" {
						// add convert function
					} else if lines[3] == "" {
						lines[0] = "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Andrine Celine Flatby"
					}
					line := strings.Join(lines, ";")
					
					readLines = append(readLines, line)
					
					//add to pos keep the track of which line is at
					pos++
				}

				fmt.Println(readLines)
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
