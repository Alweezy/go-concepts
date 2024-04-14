package main

//func main() {
//	i := 0
//	for i <= 5 {
//		fmt.Printf("The value of i is:  %d\n", i)
//		i ++
//		if i == 3 {
//			i ++
//			continue
//		}
//	}
//	fmt.Println("We exited the loop and are ready for other work ...")
//}

//coppying arrays by reference and by value

//func main() {
//	array1 := [...]string{"Alex", "Edwin", "Bob"}
//	array2 := array1
//	// copying by reference uses memory address other than the actual value
//	array3 := &array1
//
//	fmt.Println("Array1: ", array1)
//	fmt.Println("Array2: ", array2)
//
//	array1[0] = "Mark"
//	fmt.Println("New Array1: ", array1)
//	fmt.Println("New Array2: ", array2)
//	fmt.Println("Array3: ", *array3)
//}

// Multidimensional arrays

//func main() {
//	twoDimArray := [3][3]int{
//		{1: 5, 2: 12},
//		{0: 5, 2: 10},
//		{0: 2, 1: 13},
//	}
//
//	fmt.Println("Elements in array")
//
//	for i, _ := range twoDimArray[0] {
//		for j, _ := range twoDimArray[1] {
//			fmt.Printf("%d\t", twoDimArray[i][j])
//		}
//		fmt.Println("")
//	}
//}

// slices, more flexible than arrays
// create slices using make
//func main() {
//	data := [...]int{400, 70, 608, 8, 20}
//	mainSlice := make([]int, 0)
//	fmt.Printf("Length of mainSlice %d and capacity %d\n", len(mainSlice), cap(mainSlice))
//	fmt.Println(mainSlice)
//	for _, number := range data {
//		mainSlice = append(mainSlice, number)
//	}
//
//	mainSlice[4] = 700
//
//	fmt.Println(mainSlice)
//	fmt.Printf("Array after modifying slice %v", data)
//	sort.Ints(mainSlice)
//	fmt.Printf("Mainslice after sorting %v", mainSlice)
//}
