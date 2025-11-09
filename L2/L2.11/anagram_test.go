package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagramGroups(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  map[string][]string
	}{
		{
			name:  "Basic test",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			want: map[string][]string{
				"пятак":  []string{"пятак", "пятка", "тяпка"},
				"листок": []string{"листок", "слиток", "столик"},
			},
		},

		{
			name:  "Empty input",
			input: []string{},
			want:  map[string][]string{},
		},

		{
			name:  "No anagrams",
			input: []string{"дом", "квартира", "жилье"},
			want:  map[string][]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findAnagramGroups(test.input)
			assert.Equal(t, test.want, result)
		})
	}
}
