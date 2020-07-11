package new_year_chaos

import "fmt"

// Complete the minimumBribes function below.
func minimumBribes(q []int32) int32 {
	var bribes int32
	var n int32
	var i int32
	n = int32(len(q))
	for i = n - 1; i >= 0; i-- {
		if q[i] != (i + 1){
			if ((i - 1) >= 0) && q[i - 1] == (i + 1) {
				bribes++
				q[i], q[i-1] = q[i-1], q[i]
			}else if ((i - 2) >= 0) && q[i - 2] == (i + 1) {
				bribes += 2
				q[i - 2] = q[i - 1]
				q[i - 1] = q[i]
				q[i] = i + 1
			}else{
				fmt.Println("Too chaotic")
				return -1
			}
		}
	}
	return bribes
}
