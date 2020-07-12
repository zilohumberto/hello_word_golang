package bytes_example

import (
	"testing"
)

func TestGetNPlusNBytes(t *testing.T) {
	result := GetNPlusNBytes("annnna", 2)
	if string(result) != "anna"{
		t.Errorf("Excepted result anna. got %s", string(result))
	}
}
