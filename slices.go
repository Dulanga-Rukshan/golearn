package main

import "fmt"

func main() {
	// orders:= [5]int{10,20,30,40,50}// array
	// orders01 := []int{10,20,30,40} //slice

	//three ways to create slices
	//	s1 := []int{}

	//make(type, length, capacity)
	//	s2 := make([]int, 3,5)

	//array
	// s3 := [3]int{1,2,3}
	// s4:= s3[1:2]

	//make a array length long
	new := []int{1, 2, 3, 4, 85}
	new2 := new[:4:5] // [1,2,3,4,85]
	new2 = append(new2, 99)
	fmt.Println(new2)

	//loop slice
	names := []string{"kevin", "fah", "tung tung"}
	for i, name := range names {
		fmt.Printf("My name is %s and im number %d \n", name, i)
	}

	src := []int{1, 2, 3}
	drs := make([]int, len(src))
	n := copy(drs, src)
	fmt.Println(n)

	//nil slices
	var s []int
	fmt.Println(s == nil) //true
	fmt.Println(len(s))   //0
	s = append(s, 1, 2)

	//2d slices
	dslice := [][]int{
		{11, 22, 33},
		{44, 55, 66},
		{77, 88, 99},
	}

	fmt.Println(dslice[1][2])

	//dynamic 2d grid
	rows, cols := 5, 3
	dynamicgrid := make([][]int, rows)
	for i := range dynamicgrid {
		dynamicgrid[i] = make([]int, cols)
	}
	fmt.Println(dynamicgrid)

	//summary

	[]T{}          ///empty slice
	make([]int, n) //len = n
	append(t, s)   //add s to the t

}
