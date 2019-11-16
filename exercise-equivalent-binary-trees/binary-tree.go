package main

// https://godoc.org/golang.org/x/tour/tree#Tree

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	ch <- t.Value
	// defer close(ch)
	// v, ok := <-ch

	// fmt.Println(v, ok)

	rightTree := t.Right
	leftTree := t.Left

	// fmt.Println(rightTree)
	// fmt.Println(rightTree != nil)

	// fmt.Println(leftTree)
	// fmt.Println(leftTree != nil)

	if t == nil {
		return
	}

	if rightTree != nil {
		Walk(rightTree, ch)
	}

	if leftTree != nil {
		Walk(leftTree, ch)
	}

	// if !ok {
	//
	// }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	return false
}

func main() {
	c := make(chan int)
	t := tree.New(1)

	go func() {
		//
		defer close(c)
		Walk(t, c)
	}()

	for i := range c {
		fmt.Println(i)
	}

}

// https://github.com/golang/tour/blob/master/tree/tree.go#L20
// constructs a randomly-structured (but always sorted) binary tree holding the values
// Does not return a random binary tree ... always returns identical tree structure

// where k = 1, someTree is always :
// ((((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))) 10)

// where k = 2
// ((((2 (4)) 6 (8)) 10 ((12) 14 ((16) 18))) 20)

// where k =3
// ((((3 (6)) 9 (12)) 15 ((18) 21 ((24) 27))) 30)
// ... and so on

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }
