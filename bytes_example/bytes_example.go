package bytes_example

import (
	"bytes"
)

func GetNPlusNBytes(words string, n int) []byte{
	summary := make([]byte, 0)
	summary = []byte(words)
	bufSummary := bytes.NewBuffer(summary)
	a := bufSummary.Next(n)
	_ = bufSummary.Next(n) // Skip
	c := bufSummary.Next(n)

	return append(a, c...)
}

