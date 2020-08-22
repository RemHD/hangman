package hangman

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"p", "i", "s", "t", "a", "c", "h", "e"}
	guess := "p"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s. Got =%v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"p", "i", "s", "t", "a", "c", "h", "e"}
	guess := "r"
	hasLetter := letterInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s doest not contain letter %s. Got =%v", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := New(3, "")
	if err == nil {
		t.Errorf("Error should be returned when using an invalid word=''")
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	validState(t, "goodGuess", g.State)

}

func validState(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("state should be '%v'. Got=%v", expectedState, actualState)
		return false
	}
	return true
}
