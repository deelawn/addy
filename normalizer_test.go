package addy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeAddress1(t *testing.T) {

	var tests = []struct {
		name      string
		input     string
		options   []option
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
		{
			name:      "1 match, replacement; lowercased",
			input:     "976 Cherry Lane",
			options:   []option{OptionLowerCase},
			expOutput: "976 cherry ln",
			expCount:  1,
		},
		{
			name:      "2 matches, replacements; titlecased",
			input:     "23 n columbus circle",
			options:   []option{OptionTitleCase},
			expOutput: "23 N Columbus Cir",
			expCount:  2,
		},
		{
			name:      "1 match, replacement; uppercase",
			input:     "8400 south charles",
			options:   []option{OptionUpperCase},
			expOutput: "8400 S CHARLES",
			expCount:  1,
		},
		{
			name:      "2 matches, replacements; preserve titlecase",
			input:     "101 South Broadway Avenue",
			options:   []option{OptionPreserveCase},
			expOutput: "101 S Broadway Ave",
			expCount:  2,
		},
		{
			name:      "1 match, replacement; preserve lowercase",
			input:     "1465 c street",
			options:   []option{OptionPreserveCase},
			expOutput: "1465 c st",
			expCount:  1,
		},
		{
			name:      "1 match, replacement; preserve uppercase",
			input:     "904 LIBERTY PLACE",
			options:   []option{OptionPreserveCase},
			expOutput: "904 LIBERTY PL",
			expCount:  1,
		},
		{
			name:      "3 matches, replacements; preserve multicase",
			input:     "84 east Martin luther KIng BouLevard ROAD",
			options:   []option{OptionPreserveCase},
			expOutput: "84 e Martin luther KIng Blvd RD",
			expCount:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			output, count := NormalizeAddress1(tt.input, tt.options...)
			assert.Equal(t, tt.expOutput, output)
			assert.Equal(t, tt.expCount, count)
		})
	}
}

func TestNormalizeAddress2(t *testing.T) {

	var tests = []struct {
		name      string
		input     string
		options   []option
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
		{
			name:      "2 matches, replacement; PO not cased",
			input:     "p.o. box 98",
			options:   []option{OptionTitleCase},
			expOutput: "PO Box 98",
			expCount:  2,
		},
		{
			name:      "2 matches, replacement; PO cased",
			input:     "P.O. BOX 44",
			options:   []option{OptionLowerCase, OptionExcludePOCasing},
			expOutput: "po box 44",
			expCount:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			output, count := NormalizeAddress2(tt.input, tt.options...)
			assert.Equal(t, tt.expOutput, output)
			assert.Equal(t, tt.expCount, count)
		})
	}
}
