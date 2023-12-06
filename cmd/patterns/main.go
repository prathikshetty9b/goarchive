package main

import (
	"fmt"

	"github.com/prathikshetty9b/goarchive/pkg/patterns"
)

func main() {
	//Pattern 1
	fmt.Println("Pattern 1")
	patterns.PrintSquare2(5)
	fmt.Println("Pattern 2")
	patterns.PrintPattern2(5)
	fmt.Println("Pattern 3")
	patterns.PrintPattern3(5)
	fmt.Println("Pattern 4")
	patterns.PrintPattern4(5)
	fmt.Println("Pattern 5 - Inverted Pyramid")
	patterns.PrintPattern5(5)
	fmt.Println("Pattern 6 - Inverted Numbered Pyramid")
	patterns.PrintPattern22(6)
}
