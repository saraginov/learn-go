package main

// https://godoc.org/golang.org/x/tour/tree#Tree

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	/*
		todo:
			walk along tree, send to ch <- value we just read

			close channel once all values have been read

			get starting value, add value to array/list
			get left tree
			get right tree

			get starting value of left tree
			get left left tree
			get left right tree

			get starting value of left left tree
			get left left left tree
			get left left right tree

			get starting value of left left left tree
			assuming left left left left tree is empty
			assuming left left left right tree is empty

			go back up a branch (level)
			get starting value for left left right tree
			assuming left left right left tree is empty
			assuming left left right right tree is empty

			go back up a branch, both left and right of this node/branch/level explored
			get starting value of left right tree
			get left right left tree
			get left right right tree
	*/

	for {

		fmt.Println(t.Value)
		fmt.Println(t.Left)
		fmt.Println(t.Right)

	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {

	newChannel := make(chan int)

	someTree := tree.New(1)
	// treeString := someTree.String()

	Walk(someTree, newChannel)

	// fmt.Println(someTree)
	// fmt.Println(treeString)
}

// https://github.com/golang/tour/blob/master/tree/tree.go#L20
// Does not return a random binary tree ... always returns the same tree
// ... check below

// where k = 1 someTree is always :
// ((((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))) 10)

// where k = 2
// ((((2 (4)) 6 (8)) 10 ((12) 14 ((16) 18))) 20)

// where k =3
// ((((3 (6)) 9 (12)) 15 ((18) 21 ((24) 27))) 30)

// ... and so on

/**

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

*/
