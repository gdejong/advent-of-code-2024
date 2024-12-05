package input

import (
	"reflect"
	"testing"
)

func TestToIntegers(t *testing.T) {
	var tests = []struct {
		name string
		in   []string
		want []int
	}{
		{
			name: "empty case",
			in:   []string{},
			want: []int{},
		},
		{
			name: "single case",
			in:   []string{"0"},
			want: []int{0},
		},
		{
			name: "multiple case",
			in:   []string{"0", "1", "2"},
			want: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToIntegers(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %q, wanted %q", got, tt.want)
			}
		})
	}
}
