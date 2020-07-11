package two_d_array


import "testing"


func TestBasicExample(t *testing.T){
	arr := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	response := hourglassSum(arr)
	if response != 19{
		t.Errorf("Excepted value for BasicExample is 19. got %d", response)
	}
}
