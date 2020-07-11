package two_d_array



func hourglassSum(arr [][]int32) int32 {
	var _max, sum, i, j int32
	var xValues = [7]int32{-1, -1, -1, 0, 1, 1, 1}
	var yValues = [7]int32{-1, 0, 1, 0, -1, 0, 1}
	_max = -100
	for i=1; i<=4; i++{
		for j=1; j<=4; j++{
			sum = 0
			for index:=0; index<=6; index++{
				x := xValues[index] + i
				y := yValues[index] + j
				sum += arr[x][y]
			}
			if sum>_max{
				_max = sum
			}
		}
	}
	return _max
}
