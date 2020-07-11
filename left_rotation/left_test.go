package left_rotation

import (
	"reflect"
	"testing"
)

func TestRotationOne(t *testing.T){
	expected := []int32{5,1,2,3,4}
	arr := rotLeft([]int32{1,2,3,4,5}, 4)
	if !reflect.DeepEqual(expected, arr){
		t.Errorf("Excepted arr %v and got %d", expected, arr)
	}
}

func TestRotationTwo(t *testing.T){
	expected := []int32{77, 97, 58, 1, 86, 58, 26, 10, 86, 51, 41, 73, 89, 7, 10, 1, 59, 58, 84, 77}
	arr := rotLeft([]int32{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51}, 10)
	if !reflect.DeepEqual(expected, arr){
		t.Errorf("Excepted arr %v and got %d", expected, arr)
	}
}

func TestRotationThree(t *testing.T){
	expected := []int32{3,4, 5, 1, 2}
	arr := rotLeft([]int32{1,2,3,4,5}, 2)
	if !reflect.DeepEqual(expected, arr){
		t.Errorf("Excepted arr %v and got %d", expected, arr)
	}
}

func TestRotationFour(t *testing.T){
	expected := []int32{87, 97, 33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60}
	arr := rotLeft([]int32{33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60, 87, 97}, 13)
	if !reflect.DeepEqual(expected, arr){
		t.Errorf("Excepted arr %v and got %d", expected, arr)
	}
}