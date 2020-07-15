package main

import (
	"testing"
)

func TestCheckCaseOne(t *testing.T){
	tree := &BinaryTree{}
	values := []int{3, 1, 1, 2, 8, 5, 13}
	for _, value := range values{
		tree.insert(value)
	}
	treeTwo := &BinaryTree{}
	valuesTwo := []int{8, 3, 1, 1, 2, 5, 13}
	for _, value := range valuesTwo{
		treeTwo.insert(value)
	}
	response := check(tree.root, treeTwo.root)
	if response != true{
		t.Errorf("Excepting true got %t", response)
	}
}

func TestCheckCaseTwo(t *testing.T){
	tree := &BinaryTree{}
	values := []int{4, 1, 1, 2, 8, 5, 13}
	for _, value := range values{
		tree.insert(value)
	}
	treeTwo := &BinaryTree{}
	valuesTwo := []int{9, 3, 1, 1, 2, 5, 13}
	for _, value := range valuesTwo{
		treeTwo.insert(value)
	}
	response := check(tree.root, treeTwo.root)
	if response != false{
		t.Errorf("Excepting false got %t", response)
	}
}
