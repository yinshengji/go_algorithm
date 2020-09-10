package match

import (
	"github.com/kr/pretty"
	"testing"
)

func TestIsMatch(t *testing.T)  {
	text := "abb"
	parttern := "c*ab*"
	pretty.Println(IsMatch(text, parttern))
	t.Log(IsMatch(text, parttern))
}