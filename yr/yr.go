package yr

import (
	"strconv"
	"fmt"
	"log"
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
