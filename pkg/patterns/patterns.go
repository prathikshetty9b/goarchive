package patterns

import "fmt"

// Pattern 1: Square Pattern
// Input: 5
// Output:
// * * * * *
// * * * * *
// * * * * *
// * * * * *
// * * * * *
func PrintSquare(n int) {
	// Not optimized
	//TC -> O(N*N)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func PrintSquare2(n int) {
	//1. create a line with n stars
	//2. Repeat the line n times
	var stars string
	//TC -> O(2N) -> O(N)
	for i := 0; i < n; i++ {
		stars += "* "
	}
	for i := 0; i < n; i++ {
		fmt.Println(stars)
	}
}

//Pattern 2
// Sample Input 1:
// 3
// Sample Output 1:
// *
// * *
// * * *
// Explanation Of Sample Input 1 :
// For N = 3, fill all the rows and columns in the lower triangle of 3x3 matrix with ‘*’.
// N/2 Dimensional forest
func PrintPattern2(N int) {
	// Outer Loop -> Total number of rows = N
	// Inner Loop -> Star count in each row equals row number

	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}

}

// Pattern 3
// Input Format: N = 6
// Result:
// 1
// 1 2
// 1 2 3
// 1 2 3 4
// 1 2 3 4 5
// 1 2 3 4 5 6
func PrintPattern3(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

//Pattern 4
// Input Format: N = 6
// Result:
// 1
// 2 2
// 3 3 3
// 4 4 4 4
// 5 5 5 5 5
// 6 6 6 6 6 6
func PrintPattern4(N int) {
	//Rows = N
	//Each column prints row number row times.
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
	}
}

//Pattern 5
// Input Format: N = 6
// Result:
// * * * * * *
// * * * * *
// * * * *
// * * *
// * *
// *
func PrintPattern5(N int) {
	//Rows = N
	// Columns equals to (N - Row) stars (N-row) times
	for i := N; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}

func PrintPattern6(N int) {
	//Rows N
	//Columns display all numbers first and in each iteration last number is ommited
	for i := 0; i < N; i++ {
		for j := 0; j < N-i; j++ {
			fmt.Printf("%d ", j+1)
		}
		fmt.Println()
	}
}

//Pattern 7: Star Pyramid
// Input Format: N = 6
// Result:
//      *
//     ***
//    *****
//   *******
//  *********
// ***********
func PrintPattern7(N int) {
	//Spaces + Stars + Spaces
	//n-i    + 2i + 1 + n-i
	//Rows N , Columns 2(N-1) + 1
	for i := 0; i < N; i++ {
		//Left Spaces
		for j := 0; j < N-i-1; j++ {
			fmt.Printf(" ")
		}
		//Stars
		for j := 0; j < (2*i + 1); j++ {
			fmt.Printf("*")
		}
		//Right Spaces
		for j := 0; j < N-i-1; j++ {
			fmt.Printf(" ")
		}
		fmt.Println()
	}

}

//Pattern 8 -> Inverted Star Pyramid
// Input Format: N = 6
// Result:
// ***********
//  *********
//   *******
//    *****
//     ***
//      *
func PrintPattern8(N int) {
	//Spaces, Stars, Spaces
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < 2*N-(2*i+1); j++ {
			fmt.Printf("*")
		}
		for j := 0; j < i; j++ {
			fmt.Printf(" ")
		}
		fmt.Println()
	}
}

//Pattern 9 - Diamond Pattern
// Input Format: N = 6
// Result:
//      *
//     ***
//    *****
//   *******
//  *********
// ***********
// ***********
//  *********
//   *******
//    *****
//     ***
//      *
func PrintPattern9(N int) {
	//Combine Logic of Pyramid and Inverted Pyramid
	PrintPattern7(N)
	PrintPattern8(N)
}

//Pattern 10 - Half Diamond Star pattern
// Input Format: N = 6
// Result:
//      *
//      **
//      ***
//      ****
//      *****
//      ******
//      *****
//      ****
//      ***
//      **
//      *
func PrintPattern10(N int) {
	//Total rows 2N - 1
	// Observe symmetry, combine Right angle triangle
	// and inverted right angle triangle
	// for i := 0; i < N; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Printf("* ")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 0; i < N-1; i++ {
	// 	for j := N - 1; j > i; j-- {
	// 		fmt.Printf("* ")
	// 	}
	// 	fmt.Println()
	// }
	for i := 1; i <= 2*N-1; i++ {
		//first triangle
		stars := i

		//inverted triangle
		if i > N {
			stars = 2*N - i
		}

		for j := 1; j <= stars; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
