package main

type Tree struct{
	Left *Tree
	Value int
	Right *Tree
}

type BinaryTree struct{
	root *Tree
}
func inOrder(t *Tree, c chan int){
	if t.Left == nil{
		return
	}
	inOrder(t.Left, c)
	c <- t.Value
	inOrder(t.Right, c)
}

func walk(tree *Tree, c chan int){
	inOrder(tree, c)
	close(c)
}

func check(tree, treeTwo *Tree) bool{
	c := make(chan int)
	c2 := make(chan int)
	go walk(tree, c)
	go walk(treeTwo, c2)

	for {
		value, ok := <-c
		value2, ok2 := <-c2
		if ok == false && ok2 == false{
			return true
		}
		if ok != ok2{
			return false
		}
		if value != value2{
			return false
		}
	}
}

func (b *BinaryTree) insert(value int) *BinaryTree{
	if b.root ==nil{
		b.root = &Tree{Value: value}
	}else{
		b.root.insert(value)
	}
	return b
}

func (t *Tree) insert(value int) {
	if t == nil{
		return
	}
	if  value > t.Value {
		if t.Right == nil{
			t.Right = &Tree{Value: value}
		}else{
			t.Right.insert(value)
		}
	}else{
		if t.Left == nil{
			t.Left = &Tree{Value: value}
		}else{
			t.Left.insert(value)
		}
	}
}
