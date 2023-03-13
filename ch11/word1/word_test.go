package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated) = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak) = false`)
	}
}
