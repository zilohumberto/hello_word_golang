package hash_tables


import (
	"strings"
	"testing"
)

func TestCheckMagazineTestCaseOne(t *testing.T){
	answer := checkMagazine(
		[]string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"},
		[]string{"ive", "got", "some", "coconuts"},
	)
	if answer != "No"{
		t.Errorf("Harold's magazine is missing the word some. got %s", answer)
	}
}

func TestCheckMagazineTestCaseTwo(t *testing.T){
	answer := checkMagazine(
		strings.Split("two times three is not four", " "),
		strings.Split("two times two is four", " "),
	)
	if answer != "No"{
		t.Errorf("'two' only occurs once in the magazine. got %s", answer)
	}
}

func TestCheckMagazineTestCaseThree(t *testing.T){
	answer := checkMagazine(
		strings.Split("give me one grand today night", " "),
		strings.Split("give one grand today", " "),
	)
	if answer != "Yes"{
		t.Errorf("all words is fine. got %s", answer)
	}
}

func TestCheckMagazineTestCaseFour(t *testing.T){
	answer := checkMagazine(
		strings.Split("Attack at dawn", " "),
		strings.Split("attack at dawn", " "),
	)
	if answer != "No"{
		t.Errorf(" The magazine has all the right words, but there's a case mismatch. got %s", answer)
	}
}


