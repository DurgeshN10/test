// func vals() (int, int) {
// 	return 4, 6
// }

// func main() {
// 	a, b := vals()
// 	fmt.Println(a)
// 	fmt.Println(b)

// 	_, c := vals()
// 	fmt.Println(c)
// }

// func vals() (int, int) {
// 	return 3, 7
// }
// func main() {
// 	a, b := vals()
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	_, c := vals()
// 	fmt.Println(c)
// }
package main

import "fmt"

func main() {

	var i, j, rows int

	fmt.Print("Rows to Print the Star Pyramid = ")
	fmt.Scanln(&rows)

	fmt.Println("\nStar Pyramid Pattern")
	for i = 1; i <= rows; i++ {
		for j = 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k != (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// n =6, k =2

// two loops
// outer loop for no of sets i = 0;i<n
//a[j]=a[(j+k)%n]
// inner loop for rotate elements of set to k position
// how to find number of set
// how to to one by one shift elements	in set to k possition
// how to move on next set

// func rotate(n int, k int) {

// }
