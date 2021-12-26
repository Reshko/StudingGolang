package main

import "fmt"

/*
go support next types data
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

func main() {
	fmt.Println("types data in go")

	var (
		toBool   bool   = true                           // default false
		toString string = "It didnt make any difference" // default ""
		toInt    int    = 158                            // default 0
	)
	fmt.Printf("Type: %T Value: %v\n", toBool, toBool)
	fmt.Printf("Type: %T Value: %v\n", toString, toString)
	fmt.Printf("Type: %T Value: %v\n", toInt, toInt)

	//Short variable declarations

	k := 3
	fmt.Println(k)

	const (
		myName = "Oleg"
		age    = 25
	)

	fmt.Printf(myName, age)
}
