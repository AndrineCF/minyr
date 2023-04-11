package yr

import (
	"testing"
	"reflect"
)

//Hente fra tidligere oppgave
func withinTolerance(a, b, error float64) bool {
	// Først sjekk om tallene er nøyaktig like
	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	// Siden vi skal dele med b, må vi sjekke om den er 0
	// Hvis b er 0, returner avgjørelsen om d er mindre enn feilmarginen
	// som vi aksepterer
	if b == 0 {
		return difference < error
	}

	// Tilslutt sjekk den relative differanse mot feilmargin
	return (difference / math.Abs(b)) < error
}

// Testing count lines by provide information from task
func TestCountLines(t *testing.T) {
	type test struct {
		input string
		want int
	}

	tests :=[]test{
		{input:"../kjevik-temp-celsius-20220318-20230318.csv", want:16756},
		{input:"../kjevik-temp-fahr-20220318-20230318.csv", want:16756},
	}


	for _, tc := range tests {
		got := CountLines(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

//Testing the conversion from celsius to fahrenheit
func TestConvertCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input string
		want string
	}

	tests :=[]test {
		{input:"Kjevik;SN39040;18.03.2022 01:50;6",  want:"Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input:"Kjevik;SN39040;07.03.2023 18:20;0", want:"Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input:"Kjevik;SN39040;08.03.2023 02:20;-11", want:"Kjevik;SN39040;08.03.2023 02:20;12.2"},
		{input:"Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;",
		want:"Data er basert p   gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Andrine Celine Flatby;;;"},
	}

	for _, tc := range tests {
		got :=ConvertCelsiusToFahrenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
                        t.Errorf("expected: %v, got: %v", tc.want, got)
                }
        }
}

func TestAverage(t *testing.T {
	type test struct {
		input string
		want float64
	}

	tests :=[]test {
		{input:"../kjevik-temp-celsius-20220318-20230318.csv", want:8.56},
	}

	for _, tc := range tests {
		got := Average(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
