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
	s      string
	weight float64
}

func choose(choices []weighted, rng *rand.Rand) string {
	if len(choices) == 0 {
		panic("choose: no choices")
	}
	var sum float64
	for _, c := range choices {
		sum += c.weight
	}
	r := rng.Float64() * sum
	for _, c := range choices {
		r -= c.weight
		if r <= 0 {
			return c.s
		}
	}
}

func chooseLetter(rng *rand.Rand) rune {
	return letters[rng.Intn(len(letters))]
}

func chooseOtherLetter(r rune, rng *rand.Rand) rune {
	for {
		r2 := letters[rng.Intn(len(letters))]
		if unicode.ToLower(r2) == unicode.ToLower(r) {
			continue
		}
		return r2
	}
}

func choosePunt(rng *rand.Rand) rune {
	return punctuatiion[rng.Intn(len(punctuatiion))]
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
