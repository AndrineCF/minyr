package yr

import (
	"testing"
	"reflect"
)

// Testing count lines by provide information from task
func TestCountLines(t *testing.T) {
	type test struct {
		input string
		want int
	}

	tests :=[]test{
		{input:"./kjevik-temp-celsius-20220318-20230318.csv", want:16756},
		{input:"./kjevik-temp-fahr-20220318-20230318.csv", want:16756},
	}


	for _, tc := range tests {
		got := CountLines(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
