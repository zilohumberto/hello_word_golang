package new_year_chaos

import "testing"


func TestMinimumBribesTestCaseOne(t *testing.T){
	count := minimumBribes([]int32{2,1,5,3,4})
	if count != 3{
		t.Errorf("The minumum of bridbes expected is: 3. got %d", count)
	}
}

func TestMinimumBribesTestCaseTwo(t *testing.T){
	count := minimumBribes([]int32{5, 1, 2, 3, 7, 8, 6, 4})
	if count != -1{
		t.Errorf("The minumum of bridbes expected is: -1. got %d", count)
	}
}