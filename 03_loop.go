package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum_2 := 1
	for sum_2 < 10 {
		sum_2 += sum_2
	}
	fmt.Println(sum_2)

	// if

	//for i := 1; i <= 10; i++ {
	//	if i%2 == 0 {
	//		fmt.Print(i, "even")
	//	} else {
	//		fmt.Print(i, "odd")
	//	}
	//}

	//switch case

	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("U system is OS X.")
	case "linux":
		fmt.Println("U system is Linux.")
	case "windows":
		fmt.Println("U system is windows")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Monday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switch without a condition
	// is default "true"

	switch {
	case time.Now().Hour() < 12:
		fmt.Println("Good morning!")
	case time.Now().Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
