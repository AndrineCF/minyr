package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"github.com/AndrineCF/minyr/yr"
)

func getFileSize(file *os.File) (int64) {
	fileInfo, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}
	
	return fileInfo.Size()
}

func writeToFile(file *os.File, data []string) {
	writer := bufio.NewWriter(file)
	
	for _, line := range(data) {
		_, err := writer.WriteString(line + "\n")
	
		if err != nil {
			log.Fatal(err)
		}
	}

	writer.Flush()
}

func convertCelsiusFileToFahrenheit(text string) (string) { 
	lines := strings.Split(text, ";")
	
	// last line will have empty string[] since ;;;
	if lines[3] == "" {
		lines[0] = "Data er basert p   gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Andrine Celine Flatby"
	} else if lines[1] != "Stasjon" {
		lines[3] = yr.FileCelsiusToFahrenheit(lines[3])
	}

	line := strings.Join(lines, ";")

	return line
}

func main() {
	// Open the main fil
	filCel, err :=  os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	scannerCel := bufio.NewScanner(filCel)

	// closing the fil
	defer filCel.Close()

	//reading input from user
	var input string
	scannerInput := bufio.NewScanner(os.Stdin)

	//variable
	var readLines []string

	// loop to the user quit the program
	for scannerInput.Scan() {
		input = scannerInput.Text()
		
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" || input == "konvert" {
			fmt.Println("Konverterer all maalingene gitt i grader Celsius til grader Fahrenheit")
			// trying opening fil
			fileConv, errConv := os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_RDWR|os.O_CREATE, 0755)
			
			// Check if the fil is empty
			if getFileSize(fileConv) == 0 {
				

				//read throught the celsius file
				for scannerCel.Scan() {
										
					fixLine := convertCelsiusFileToFahrenheit(scannerCel.Text())					
					readLines = append(readLines, fixLine)
					
				}
			} else {
								
        			scannerInputC := bufio.NewScanner(os.Stdin)
				fmt.Println("Filen finnes allerede. Vil du generere fil paa nytt?")
				for scannerInputC.Scan() {
					if scannerInputC.Text() == "j"|| scannerInputC.Text() == "ja"  {
						for scannerCel.Scan() {
							fixLine := convertCelsiusFileToFahrenheit(scannerCel.Text())
							readLines = append(readLines, fixLine)
						}
						break
					} else if scannerInputC.Text() == "n" || scannerInputC.Text() == "nei" {
						break
					} else {
						fmt.Println("Ikke gyldig valg, prov paa nytt")
				}
			}
			writeToFile(fileConv, readLines)
			readLines = nil
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
