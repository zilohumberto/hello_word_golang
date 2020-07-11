package left_rotation

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	var b []int32
	b = append(b, a...)
	b = append(b, a...)
	return b[d:len(a)+int(d)]
}
