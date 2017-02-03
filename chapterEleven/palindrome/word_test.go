package word

import (
	// "fmt"
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func TestPalindrome(t *testing.T){
	var tests = []struct {
		input string
		want bool
	}{
		{"kayak", true},
		{"A man, a plan, a canal: Panama", true},
		{"Et se resservir, ivresse reste", true},
		{"MA RR a M", true},
		{":::sas?::sas!", true},
		{"été", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"feat", false},
		{"ab", false},
		{"MARRAMS", false},
		{"lemon", false},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

//randomPalindrome returns a palindrome whose length and contents are 
//derived from the pseudorandom number generator rng
//palindrome returned will contain unicode characters between 0x0020 and 0x007A
func randomPalindrome(rng *rand.Rand) string{
	n := rng.Intn(25) // random length up to 24
	// fmt.Println(n)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		value := rng.Intn(0x007B)
		r := rune(value)
		if r >= 0x0020 {
			runes[i] = r
			runes[n-1-i] = r
		} else {
			runes[i] = rune(' ')
			runes[n-1-i] = rune(' ')
		}
	}
	return string(runes)
}

func countLetters(s string) int {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	return len(letters)
}

func TestRandomPalindromes(t *testing.T) {
	//initialize a pseudorandom number generator
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	// fmt.Printf("rng=%d\n", rng)
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		// fmt.Printf("%d: %q\n", i, p)
		if !IsPalindrome(p) && countLetters(p) > 1 {
			t.Errorf("IsPalindrome(%q) == false", p)
		}
	}
}

// func randomNonPalindrome(rng1 *rand.Rand, rng2 *rand.Rand) string{
// 	n := rng1.Intn(25) // random length up to 24
// 	runes := make([]rune, n)
// 	for i := 0; i < (n+1)/2; i++ {
// 		value1 := rng1.Intn(0x007B)
// 		value2 := rng2.Intn(0x007B)
// 		// r := rune(value)
// 		for unicode.ToLower(rune(value1)) == unicode.ToLower(rune(value2)) {
// 			fmt.Println("trying again")
// 			rng1 = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
// 			value1 = rng1.Intn(0x007B) 
// 			if !(unicode.ToLower(rune(value1)) == unicode.ToLower(rune(value2))){
// 				fmt.Printf("%q, %q\n", unicode.ToLower(rune(value1)), unicode.ToLower(rune(value2)))
// 			}
// 		}
// 		runes[i] = rune(value1)
// 		runes[n-1-i] = rune(value2)
// 	}
// 	return string(runes)
// }

// func TestRandomNonPalindromes(t *testing.T) {
// 	//initialize a pseudorandom number generator
// 	seed := time.Now().UTC().UnixNano()
// 	t.Logf("Random seed: %d", seed)
// 	rng1 := rand.New(rand.NewSource(seed))
// 	rng2 := rand.New(rand.NewSource(seed))
// 	// fmt.Printf("rng=%d\n", rng)
// 	for i := 0; i < 500; i++ {
// 		p := randomNonPalindrome(rng1, rng2)
// 		// fmt.Printf("%d: %q\n", i, p)
// 		if IsPalindrome(p) {
// 			t.Errorf("IsPalindrome(%q) == true", p)
// 		}
// 	}
// }

// 	if !IsPalindrome("detartrated") {
// 		t.Error("IsPalindrome('detartrated') == false")
// 	}
// 	if !IsPalindrome("kayak") {
// 		t.Error("IsPalindrome('kayak') == false")
// 	}
// }

// func TestCanalPalindrome(t *testing.T) {
// 	input := "A man, a plan, a canal: Panama"
// 	if !IsPalindrome(input){
// 		t.Errorf("IsPalindrome(%q) == false", input)
// 	}
// }

// func TestFrenchPalindrome(t *testing.T) {
// 	if !IsPalindrome("été") {
// 		t.Error("IsPalindrome('été') == false")
// 	}
// }

// func TestNonPalindrome(t *testing.T) {
// 	if IsPalindrome("palindrome") {
// 		t.Error("IsPalindrome('palindrome') == true")
// 	}
// }
