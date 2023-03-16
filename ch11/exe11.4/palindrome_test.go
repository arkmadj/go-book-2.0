package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

var grammar = map[string][]weighted{
	"NON": []weighted{
		{"a c b", 1},
		{"a b", 1},
		{"a NON a", 30},
		{"a NON b", 30},
		{"a PAL b", 30},
	},
	"PAL": []weighted{
		{"eps ", 1},
		{"a ", 1},
		{"a a", 1},
		{"a PAL a", 40},
	},
}

var letters []rune
var punctuatiion []rune
var punctProb = 0.1

type weighted struct {
	s        string
	weighted float64
}

func init() {
	for r := rune(0x21); r < 0x7e; r++ {
		switch {
		case unicode.IsLetter(r):
			letters = append(letters, r)
		case unicode.IsPunct(r):
			punctuatiion = append(punctuatiion, r)
		}
	}
}
