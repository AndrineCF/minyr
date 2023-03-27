package yr

import (
	"strconv"
	"fmt"
	"strings"
	"github.com/AndrineCF/funtemps/conv"
)

func FileCelsiusToFahrenheit(cel string) (string) {
	//split the string
	fahr, _ := strconv.ParseFloat(cel, 64)
	
	fahrString := fmt.Sprintf("%.1f", conv.CelsiusToFahrenheit(fahr))
	return fahrString
}
