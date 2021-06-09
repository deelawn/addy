package addy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeAddress1(t *testing.T) {

	var tests = []struct {
		name      string
		input     string
		expOutput string
		expCount  int
	}{
		{
			name:      "1 match",
			input:     "42 wallaby way",
			expOutput: "42 WALLABY WAY",
			expCount:  1,
		},
		{
			name:      "2 matches, replacements",
			input:     "401 North Coast Highway",
			expOutput: "401 N COAST HWY",
			expCount:  2,
		},
		{
			name:      "1 match, replacement; remove spaces",
			input:     "  1465  c     street   ",
			expOutput: "1465 C ST",
			expCount:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			output, count := NormalizeAddress1(tt.input)
			assert.Equal(t, tt.expOutput, output)
			assert.Equal(t, tt.expCount, count)
		})
	}
}

func TestNormalizeAddress2(t *testing.T) {

	var tests = []struct {
		name      string
		input     string
		expOutput string
		expCount  int
	}{
		{
			name:      "1 match",
			input:     "unit 10",
			expOutput: "UNIT 10",
			expCount:  1,
		},
		{
			name:      "2 matches, replacement",
			input:     "p.o. box 44",
			expOutput: "PO BOX 44",
			expCount:  2,
		},
		{
			name:      "1 match, replacement; remove spaces",
			input:     "   32nd     floor ",
			expOutput: "32ND FL",
			expCount:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			output, count := NormalizeAddress2(tt.input)
			assert.Equal(t, tt.expOutput, output)
			assert.Equal(t, tt.expCount, count)
		})
	}
}
