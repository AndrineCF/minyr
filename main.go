package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"github.com/AndrineCF/minyr/yr"
)

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

func options() {
	fmt.Println("Det er 3 alternativ aa velge mellom:")
	fmt.Println("convert (lager en fil som koncert celsius til fahrenheit)")
	fmt.Println("average (tar gjennomsnitt temperaturen i celsius og fahrenheit")
	fmt.Println("q/exit (avslutter programmet)")
}

func main() {
	// Open the main fil
	filCel, err :=  os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	scannerCel := bufio.NewScanner(filCel)

	// trying opening fil
	fileConv, errConv := os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_RDWR|os.O_CREATE, 0755)

	//check for errors
	if err != nil {
		log.Fatal(err)
	}

	if errConv != nil {
		log.Fatal(errConv)
	}

	// closing the fil
	defer filCel.Close()
	defer fileConv.Close()


	//reading input from user
	var input string
	scannerInput := bufio.NewScanner(os.Stdin)

	options()

	// loop to the user quit the program
	for scannerInput.Scan() {
		input = scannerInput.Text()
		
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" || input == "konvert" {
			fmt.Println("Konverterer all maalingene gitt i grader Celsius til grader Fahrenheit")
			
			var readLines []string

			// Check if the fil is empty
			if yr.CountLines("kjevik-temp-fahr-20220318-20230318.csv") == 0 {
				
				//read throught the celsius file
				for scannerCel.Scan() {
										
					fixLine := yr.ConvertCelsiusToFahrenheit(scannerCel.Text())					
					readLines = append(readLines, fixLine)
				}
			} else {
								
        			scannerInput := bufio.NewScanner(os.Stdin)
				fmt.Println("Filen finnes allerede. Vil du generere fil paa nytt?")
				for scannerInput.Scan() {
					fmt.Println(scannerInput.Text())
					if scannerInput.Text() == "j"|| scannerInput.Text() == "ja"  {
						for scannerCel.Scan() {
							fixLine := yr.ConvertCelsiusToFahrenheit(scannerCel.Text())
							readLines = append(readLines, fixLine)
						}
						break
					} else if scannerInput.Text() == "n" || scannerInput.Text() == "nei" {
						break
					} else {
						fmt.Println("Ikke gyldig valg, prov paa nytt")
					}
				}
			}
			
			writeToFile(fileConv, readLines)

			// reset the variable
			readLines = nil
			
		//chose the aaverage options
		} else if  input == "average" {
			fmt.Println("Hva vil du ha gjennomsnit skrive ut som? c for Celsius eller f for Fahrenheit")
			scannerInput := bufio.NewScanner(os.Stdin)

			//goes through scanner
			for scannerInput.Scan() {
				
				if scannerInput.Text() == "c" || scannerInput.Text() == "f" {
					average := fmt.Sprintf("%.2f", yr.Average(scannerInput.Text()))
					fmt.Println("Gjennomsnitt verdien i Celsius er ", average)
					break
				} else {
					fmt.Println("Ikke gyldig valg, prove paa nytt")
				}
			}
		} else {
			fmt.Println("Ikke gylding alternativ, venligst prove igjen")
		}

		options()
	}

}
