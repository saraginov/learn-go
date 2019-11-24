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
	c1 := make(chan int)
	c2 := make(chan int)

	areT1andT2Identical := false

	var sliceT1, sliceT2 []int

	go func() {
		defer close(c1)
		Walk(t1, c1)
	}()

	for i := range c1 {
		sliceT1 = append(sliceT1, i)
	}

	go func() {
		defer close(c2)
		Walk(t2, c2)
	}()

	for i := range c2 {
		sliceT2 = append(sliceT2, i)
	}

	fmt.Printf("t1 : %v\n", sliceT1)
	fmt.Printf("t2 : %v\n", sliceT2)

	defer func() {
		// loop through first slice and compare its values to those of second slice
		for index, value := range sliceT1 {
			fmt.Println(index)
			fmt.Println(value)

			hasFoundMatch := false

			for i2, v2 := range sliceT2 {
				if value == v2 {
					p1 := sliceT2[0:i2]
					p2 := sliceT2[i2:len(sliceT2)]

					// https://yourbasic.org/golang/clear-slice/
					sliceT2 = nil
					// sliceT2 = sliceT2[:0]

					for _, valueP1 := range p1 {
						sliceT2 = append(sliceT2, valueP1)
					}

					for _, valueP2 := range p2 {
						sliceT2 = append(sliceT2, valueP2)
					}

					hasFoundMatch = true
					break
				}
			}

			if !hasFoundMatch {
				areT1andT2Identical = false
				break
			}
		}

	}()

	return areT1andT2Identical // false
}

func main() {
	c := make(chan int)
	t := tree.New(1)
	t1 := tree.New(1)
	t2 := tree.New(2)

	go func() {
		defer close(c)
		Walk(t, c)
	}()

	for i := range c {
		fmt.Println(i)
	}

	areTAndT1TheSame := Same(t, t1)
	areTAndT2TheSame := Same(t, t2)

	fmt.Printf("Are t and t1 the same ? -> %v \n", areTAndT1TheSame)
	fmt.Printf("Are t and t2 the same ? -> %v \n ", areTAndT2TheSame)
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
