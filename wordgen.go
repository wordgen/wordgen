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

// Generator is used to generate random words from a provided wordlist with
// customizable options for the number of words, casing, and separator.
type Generator struct {
	Words     []string
	Count     int
	Casing    string
	Separator string
}

// NewGenerator initializes and returns a new Generator with default settings.
func NewGenerator() (g Generator) {
	g.Words = []string{}
	g.Count = 1
	g.Casing = ""
	g.Separator = " "

	return g
}

// Generate creates a string composed of random words from the Generator's wordlist.
// The number of words, casing, and separator are determined by the Generator's settings.
// Returns an error if the wordlist is empty or if random number generation fails.
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

		randomWord := g.Words[int(randomNum.Int64())]

		switch g.Casing {
		case "upper":
			b.WriteString(cases.Upper(language.English).String(randomWord))
		case "title":
			b.WriteString(cases.Title(language.English).String(randomWord))
		case "lower":
			b.WriteString(cases.Lower(language.English).String(randomWord))
		default:
			b.WriteString(randomWord)
		}

		if i < g.Count-1 {
			b.WriteString(g.Separator)
		}
	}

	return b.String(), nil
}
