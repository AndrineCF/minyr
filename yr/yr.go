package yr

import (
	"strconv"
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
	"github.com/AndrineCF/funtemps/conv"
)

//The function convert celsius to fahrenheit by usnig funtemps
func CelsiusToFahrenheit(cel string) (string) {
	//split the string
	fahr, err := strconv.ParseFloat(cel, 64)

	//check for errors
	if err != nil {
		log.Fatal(err)
	}
	
	//convert back to string
	fahrString := fmt.Sprintf("%.1f", conv.CelsiusToFahrenheit(fahr))
	return fahrString
}

//The function convert a string to fahrenheit 
func ConvertCelsiusToFahrenheit(text string) (string) { 
	lines := strings.Split(text, ";")
	
	// last line will have empty string[] since ;;;
	if lines[3] == "" {
		lines[0] = "Data er basert p   gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Andrine Celine Flatby"
	} else if lines[1] != "Stasjon" {
		lines[3] = CelsiusToFahrenheit(lines[3])
	}

	line := strings.Join(lines, ";")

	return line
}

//Count lines in a file
func CountLines(fileName string) (int) {
        // open choosen the 
	file, err :=  os.Open(fileName)
        
        //check for errors
	if err != nil {
		log.Fatal(err)
	}

	//closing file when finished
	defer file.Close()

	counter := 0
	scannerFile := bufio.NewScanner(file)
	for scannerFile.Scan() {
		counter++
	}

	return counter
}

func Average(tempType string, fileName string) (float64){
	//Open file that is constant
	file, err := os.Open(fileName)

	//check for errors
	if err != nil {
		log.Fatal(err)
	}

	//closing the file when finished
	defer file.Close()

	// variable to calcuted the data
	var counter float64
	var total float64

	//setup the scanner
	scanner := bufio.NewScanner(file)

	//goes through
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), ";")

		//check if it is first or the last line
		if (lines[3] == "Lufttemperatur" || lines [3] == "") {
			continue
		}
		
		num, err := strconv.ParseFloat(lines[3], 64)
		
		if err != nil {
			log.Fatal(err)
		}
		
		//check if needs to be convert to fahrenheit
		if tempType == "f" {
			num = conv.CelsiusToFahrenheit(num)
		}

		counter++
		total += num
	}
	return (total/counter)
}
