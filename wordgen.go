// Copyright (C) 2024  Daniel Kuehn <daniel@kuehn.foo>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package wordgen

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Generator is a struct that holds the configuration for generating words.
type Generator struct {
	Words     []string     // List of words to choose from.
	Count     int          // Number of words to generate.
	Casing    string       // Letter casing: upper, title, lower, or "" (no casing).
	Separator string       // String used to separate generated words.
	Language  language.Tag // Language tag for casing transformations.
}

// NewGenerator initializes a new Generator with default values.
func NewGenerator() *Generator {
	return &Generator{
		Words:     []string{},
		Count:     1,
		Casing:    "",
		Separator: " ",
		Language:  language.English,
	}
}

// Generate creates a string of randomly chosen words based on the Generator configuration.
func (g *Generator) Generate() (string, error) {
	if len(g.Words) == 0 {
		return "", fmt.Errorf("wordlist cannot be empty")
	}

	words := make([]string, g.Count)
	var caser func(string) string

	switch g.Casing {
	case "upper":
		caser = cases.Upper(g.Language).String
	case "title":
		caser = cases.Title(g.Language).String
	case "lower":
		caser = cases.Lower(g.Language).String
	default:
		caser = func(s string) string { return s }
	}

	maxNum := big.NewInt(int64(len(g.Words)))

	for i := 0; i < g.Count; i++ {
		randomNum, err := rand.Int(rand.Reader, maxNum)
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %v", err)
		}

		randomWord := g.Words[randomNum.Int64()]
		words[i] = caser(randomWord)
	}

	return strings.Join(words, g.Separator), nil
}
