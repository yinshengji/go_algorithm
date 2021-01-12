package palindrome

import "testing"

func TestIsPalindrome(t *testing.T)  {
	s := "A man, a plan, a canal: Panama"
	if !isPalindrome(s) {
		t.Fail()
	}

	s = "race a car"
	if isPalindrome(s) {
		t.Fail()
	}

	s = ""
	if !isPalindrome(s) {
		t.Fail()
	}
}
