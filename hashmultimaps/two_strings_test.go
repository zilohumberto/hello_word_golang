package hashmultimaps

import "testing"

func TestTwoStringsCaseOne(t *testing.T){
	response := twoStrings("hello", "world" )
	if response != "YES"{
		t.Errorf("Excepted YES, got %s", response)
	}
}

func TestTwoStringsCaseTwo(t *testing.T){
	response := twoStrings("hi", "world" )
	if response != "NO"{
		t.Errorf("Excepted NO, got %s", response)
	}
}
