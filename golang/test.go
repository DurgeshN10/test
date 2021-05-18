// package main

// import "fmt"

// func main() {
// 	// -dimensional array using index
// 	var arr1 [2][2]int
// 	arr1[0][0] = 100
// 	arr1[0][1] = 200
// 	arr1[1][0] = 300
// 	arr1[1][1] = 400

// 	// Accessing the values of the array
// 	fmt.Println("Elements of array 2")
// 	for p := 0; p < 2; p++ {
// 		for q := 0; q < 2; q++ {
// 			fmt.Println(arr1[p][q])
// 		}
// 	}
// }

// Go program to illustrate the
// concept of ellipsis in an array
// package main

// import "fmt"

// func main() {

// 	// Creating an array whose size is determined
// 	// by the number of elements present in it
// 	// Using ellipsis
// 	myarray := []string{"GFG", "gfg", "geeks", "GeeksforGeeks", "GEEK"}

// 	fmt.Println("Elements of the array: ", myarray)

// 	// Length of the array
// 	// is determine by
// 	// Using len() method
// 	fmt.Println("Length of the array is:", len(myarray))
// }
// package main

// import "fmt"

// func main() {
// 	n := 4
// 	array := []int{1, 2, 3, 4}
// 	fmt.Println(len(array) + len(array)/2)
// 	fmt.Println(n, array)
// }
