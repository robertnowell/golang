package split

import (
	// "fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T){
	var tests = []struct {
		input 	string
		sep 	string
		want	int
	}{
		{"a:b:c", ":", 3},
		{"a!b!c", "!", 3},
		{"a", ":", 1},
		{"a b c", " ", 3},
		{"a:b", ":", 2},
		{"a-b-c", "-", 3},
	}
	for _, test := range tests {
		words := strings.Split(test.input, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("length of list created by string (%q) split by "+
				"delimiter (%q) equals (%d), want length = %d", 
				test.input, test.sep, got, test.want)			
		}
	}
}
