package main

import "fmt"

// IsLeapYear will run through a series of numbers and test if it return true or false
func IsLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 {
		fmt.Println("It is a leap year!")
		return true
	} else if year%400 == 0 {
		fmt.Println("It is a leap year!")
		return true
	} else {
		fmt.Println("Try another year")
		return false
	}
}

func main() {
	IsLeapYear(2000)
	IsLeapYear(2015)
	IsLeapYear(1970)
	IsLeapYear(1996)
	IsLeapYear(2100)
	IsLeapYear(1800)
	IsLeapYear(2016)
	IsLeapYear(1900)
	IsLeapYear(1786)
}
