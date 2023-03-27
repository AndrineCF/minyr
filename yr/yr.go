package yr

import (
	"strconv"
	"fmt"
)

func FileCelsiusToFahrenheit(cel String) (String) {
	fahr := strconv.ParseFloat(cel, 64)
	return  fmt.Springf("%.1f", fahr")
}
