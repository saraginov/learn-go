package main

import "fmt"

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)     // 3
	n := m + len(data)  // 6
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// append() returns a slice containing values passed to it, it doesn't look
	// append() works on multiple slices
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl3: ", l)

	t := []string{"hello", "how", "are", "you", "doing"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// sTest := make([]byte, 5)
	sTest := []int{1, 2, 3, 4, 5}
	fmt.Println("sTest:", sTest)
	sTest = sTest[2:4]
	fmt.Println("sTest:", sTest)
	fmt.Println("cap(sTest):", cap(sTest))
	sTest = sTest[:cap(sTest)]
	// sTest = sTest[:4]
	fmt.Println("sTest:", sTest)

	test := []byte{'a', 'b', 'c'}
	fmt.Println("test:", test)
	test = AppendByte(test, 'e', 'f', 'g')
	fmt.Println("test:", test)
}
