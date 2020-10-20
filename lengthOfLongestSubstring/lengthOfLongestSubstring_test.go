package lengthOfLongestSubstring

import (
	"github.com/kr/pretty"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T)  {
	s := "abcbadec"
	pretty.Println(LengthOfLongestSubstring(s))
	t.Log(LengthOfLongestSubstring(s))
}