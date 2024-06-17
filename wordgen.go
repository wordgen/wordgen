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

// Generator is a struct that holds configuration for generating words.
type Generator struct {
	Words     []string     // List of words to choose from.
	Count     int          // Number of words to generate.
	Casing    string       // Letter casing: upper, title, lower, or "" (no casing).
	Language  language.Tag // Language tag for casing transformations.
	Separator string       // String used to separate generated words.
}

// NewGenerator initializes a new Generator with default values.
func NewGenerator() (g Generator) {
	g.Words = []string{}
	g.Count = 1
	g.Casing = ""
	g.Separator = " "
	g.Language = language.English

	return g
}

// Generate creates a string of randomly chosen words based on the Generator configuration.
func (g Generator) Generate() (string, error) {
	if len(g.Words) == 0 {
		return "", fmt.Errorf("ERROR: wordlist cannot be empty")
	}

	var b strings.Builder

	for i := 0; i < g.Count; i++ {
		randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(g.Words))))

		if err != nil {
			return "", fmt.Errorf("ERROR: failed to generate random number: %v", err)
		}

		randomWord := g.Words[randomNum.Int64()]

		switch g.Casing {
		case "upper":
			b.WriteString(cases.Upper(g.Language).String(randomWord))
		case "title":
			b.WriteString(cases.Title(g.Language).String(randomWord))
		case "lower":
			b.WriteString(cases.Lower(g.Language).String(randomWord))
		default:
			b.WriteString(randomWord)
		}

		if i < g.Count-1 {
			b.WriteString(g.Separator)
		}
	}

	return b.String(), nil
}
