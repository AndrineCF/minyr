package yr

import (
	"strconv"
	"fmt"
	"github.com/andrinecCF/funtemps/conv"
)

func FileCelsiusToFahrenheit(cel string) (string) {
	fahr, _ := strconv.ParseFloat(cel, 64)
	
	fahrString := fmt.Sprintf("%.1f", fahr)
	return fahrString
}
