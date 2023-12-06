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
		//Right Spaces -> unnecessary
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
		//unecessary
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

//Pattern11
// Input Format: N = 6
// Result:
// 1
// 01
// 101
// 0101
// 10101
// 010101
func PrintPattern11(N int) {
	//switching flag from 1-0 approach
	//Starting flag
	flag := 1
	//Outer loop for rows
	for i := 0; i < N; i++ {
		//if the row index is even 1 is printed else 0
		if i%2 == 0 {
			flag = 1
		} else {
			flag = 0
		}
		//Alternate between 1 and 0 in the inner for loop
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", flag)
			flag = 1 - flag
		}
		fmt.Println()
	}
}

//Pattern12
// Input Format: N = 6
// Result:
// 1          1
// 12        21
// 123      321
// 1234    4321
// 12345  54321
// 123456654321
func PrintPattern12(N int) {
	//Numbers spaces Numbers
	//Row = N
	//Column = 2N
	for i := 1; i <= N; i++ {
		//print left triangle
		for j := 1; j <= i; j++ {
			fmt.Printf("%d", j)
		}
		//print Spaces
		for j := 1; j <= 2*(N-i); j++ {
			fmt.Printf(" ")
		}
		//Print right triangle
		for j := i; j >= 1; j-- {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
}

//Pattern 13
// Input Format: N = 6
// Result:
// 1
// 2  3
// 4  5  6
// 7  8  9  10
// 11  12  13  14  15
// 16  17  18  19  20  21
func PrintPattern13(N int) {
	//Increment counter in each iteration
	count := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			count++
			fmt.Printf("%d ", count)
		}
		fmt.Println()
	}
}

//Pattern14
// Input Format: N = 6
// Result:
// A
// A B
// A B C
// A B C D
// A B C D E
// A B C D E F
func PrintPattern14(N int) {
	count := 0
	for i := 0; i < N; i++ {
		for j := 'A'; j <= rune(int('A')+i); j++ {
			count++
			fmt.Printf("%v ", string(j))
		}
		fmt.Println()
	}
}

//Pattern15
// Input Format: N = 6
// Result:
// A B C D E F
// A B C D E
// A B C D
// A B C
// A B
// A
func PrintPattern15(N int) {
	//Increment counter in each iteration
	count := 0
	for i := 0; i < N; i++ {
		//Start from A and print till 'A' + N - i
		for j := 'A'; j < rune(int('A')+N-i); j++ {
			count++
			fmt.Printf("%v ", string(j))
		}
		fmt.Println()
	}
}

//Pattern 16
// Input Format: N = 6
// Result:
// A
// B B
// C C C
// D D D D
// E E E E E
// F F F F F F
func PrintPattern16(N int) {
	count := 0
	for i := 0; i < N; i++ {
		char := string(rune(int('A') + i))
		for j := 0; j <= i; j++ {
			count++
			fmt.Printf("%v ", char)
		}
		fmt.Println()
	}
}

//Pattern17
// Input Format: N = 6
// Result:
//      A
//     ABA
//    ABCBA
//   ABCDCBA
//  ABCDEDCBA
// ABCDEFEDCBA
func PrintPattern17(N int) {
	//Spaces Char Spaces
	//Col -> 2N-1
	//spaces N-i
	for i := 0; i < N; i++ {
		//print left spaces
		for j := 0; j < N-i-1; j++ {
			fmt.Printf(" ")
		}

		char := 'A'
		middle := (2 * i) / 2
		//Print characters
		for j := 0; j < (2*i)+1; j++ {
			fmt.Printf("%v", string(char))
			//Increase char till middle
			if j <= middle {
				char = rune(int(char) + 1)
			} else {
				char = rune(int(char) - 1)
			}

		}
		//printing right spaces is unnecessary
		for j := 0; j < N-i-1; j++ {
			fmt.Printf(" ")
		}

		//move to next line
		fmt.Println()
	}
}

// Pattern 18
// Input Format: N = 6
// Result:
// F
// E F
// D E F
// C D E F
// B C D E F
// A B C D E F
func PrintPattern18(N int) {
	for i := 0; i < N; i++ {
		char := rune('A' + N - i - 1)
		for j := 0; j <= i; j++ {
			fmt.Printf("%v ", string(char))
			char += 1
		}
		fmt.Println()
	}
}

//Pattern 19
// Input Format: N = 6
// Result:
// ************
// *****  *****
// ****    ****
// ***      ***
// **        **
// *          *
// *          *
// **        **
// ***      ***
// ****    ****
// *****  *****
// ************
func PrintPattern19(N int) {
	//row = 2N, col = 2N
	//Combination of right angle triangle and inverted triangle
	spaces := 0
	for i := 0; i < 2*N; i++ {
		if i < N {
			//print top left inverted triangle
			for j := N; j > i; j-- {
				fmt.Printf("*")
			}
			//print spaces
			for j := 0; j < spaces; j++ {
				fmt.Printf(" ")
			}
			//print top right inverted triangle
			for j := N; j > i; j-- {
				fmt.Printf("*")
			}
			spaces += 2
		} else {
			spaces -= 2
			//print bottom left  triangle
			for j := N; j <= i; j++ {
				fmt.Printf("*")
			}
			//print spaces
			for j := 0; j < spaces; j++ {
				fmt.Printf(" ")
			}
			//print bottom right triangle
			for j := N; j <= i; j++ {
				fmt.Printf("*")
			}

		}
		fmt.Println()
	}
}

// Pattern 20
// Input Format: N = 6
// Result:
// *          *
// **        **
// ***      ***
// ****    ****
// *****  *****
// ************
// *****  *****
// ****    ****
// ***      ***
// **        **
// *          *
func PrintPattern20(N int) {
	//row = 2N-1, col = 2N
	//Combination of right angle triangle and inverted triangle
	spaces := 2*N - 2
	for i := 1; i < 2*N; i++ {
		//Stars for first half
		stars := i

		//stars for second half
		if i > N {
			stars = 2*N - i
		}
		//print left stars
		for j := 1; j <= stars; j++ {
			fmt.Printf("*")
		}
		//print Spaces
		for j := 1; j <= spaces; j++ {
			fmt.Printf(" ")
		}
		//print right stars
		for j := 1; j <= stars; j++ {
			fmt.Printf("*")
		}

		fmt.Println()
		if i >= N {
			spaces += 2
		} else {
			spaces -= 2
		}

	}
}

//Pattern 21
// Input Format: N = 6
// Result:
// ******
// *    *
// *    *
// *    *
// *    *
// ******
func PrintPattern21(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			//print stars at edges
			if i == 1 || i == N ||
				j == 1 || j == N {
				fmt.Printf("*")
			} else { //print space otherwise
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

// Pattern 22
// Input Format: N = 6
// Result:
// 6 6 6 6 6 6 6 6 6 6 6
// 6 5 5 5 5 5 5 5 5 5 6
// 6 5 4 4 4 4 4 4 4 5 6
// 6 5 4 3 3 3 3 3 4 5 6
// 6 5 4 3 2 2 2 3 4 5 6
// 6 5 4 3 2 1 2 3 4 5 6
// 6 5 4 3 2 2 2 3 4 5 6
// 6 5 4 3 3 3 3 3 4 5 6
// 6 5 4 4 4 4 4 4 4 5 6
// 6 5 5 5 5 5 5 5 5 5 6
// 6 6 6 6 6 6 6 6 6 6 6

func PrintPattern22(N int) {
	//least distance from the border
	//read about this
	for i := 0; i < 2*N-1; i++ {
		for j := 0; j < 2*N-1; j++ {
			top, left, right, bottom := i, j, (2*N-2)-j, (2*N-2)-i
			minimum := min(top, left, right, bottom)
			fmt.Printf("%d ", N-minimum)
		}
		fmt.Println()
	}
}
