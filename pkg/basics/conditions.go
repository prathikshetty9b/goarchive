package basics

import (
	"fmt"
)

func GetGrade() (grade string, err error) {
	var mark int
	fmt.Printf("Enter mark: ")
	n, err := fmt.Scanf("%d\n", &mark)
	if err != nil || n != 1 {
		// handle invalid input
		return grade, err
	}
	if mark >= 0 && mark < 40 {
		grade = "F"
	} else if mark < 50 {
		grade = "E"
	} else if mark < 60 {
		grade = "D"
	} else if mark < 75 {
		grade = "C"
	} else if mark < 90 {
		grade = "B"
	} else if mark <= 100 {
		grade = "A"
	} else {
		err = fmt.Errorf("input %v is invalid", mark)
	}
	return grade, err
}

func GetDay(day int) {
	switch day {
	case 1:
		fmt.Print("Monday")
	case 2:
		fmt.Print("Tuesday")
	case 3:
		fmt.Print("Wednesday")
	case 4:
		fmt.Print("Thursday")
	case 5:
		fmt.Print("Friday")
	case 6:
		fmt.Print("Saturday")
	case 7:
		fmt.Print("Sunday")
	default:
		fmt.Print("Invalid")
	}
}
