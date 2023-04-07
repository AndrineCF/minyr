package yr

import (
	"strconv"
	"fmt"
	"log"
	"os"
	"bufio"
	"github.com/AndrineCF/funtemps/conv"
)

//The function convert celsius to fahrenheit by usnig funtemps
func FileCelsiusToFahrenheit(cel string) (string) {
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

//Count lines in a file
func CountLines(fileNavn string) (int) {
        // open choosen the 
	file, err :=  os.Open(fileNavn)
        
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
