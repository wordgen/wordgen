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

// WordGenerator generates a specified number of random words from the given wordlist,
// applies the specified case transformation ("upper", "title", "lower"), and joins
// them with the specified separator. If the casing is an empty string, it writes the
// word directly to the string builder.
//
// Returns a string of generated words and nil, or an empty string and an error if the
// wordlist is empty or if random number generation fails.
func WordGenerator(list []string, count int, casing, separator string) (string, error) {
	if len(list) == 0 {
		return "", fmt.Errorf("ERROR: list cannot be empty")
	}

	var b strings.Builder

	for i := 0; i < count; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))

		if err != nil {
			return "", fmt.Errorf("ERROR: failed to generate random number: %v", err)
		}

		randomWord := list[int(n.Int64())]

		switch casing {
		case "upper":
			b.WriteString(cases.Upper(language.English).String(randomWord))
		case "title":
			b.WriteString(cases.Title(language.English).String(randomWord))
		case "lower":
			b.WriteString(cases.Lower(language.English).String(randomWord))
		default:
			b.WriteString(randomWord)
		}

		if i < count-1 {
			b.WriteString(separator)
		}
	}

	return b.String(), nil
}
