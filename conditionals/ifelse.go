package conditionals

import (
	"log"
)

func CheckIfEven(num int) string {
	if num%2 == 0 {
		return "even"
	} else if num%2 == 1 {
		return "odd"
	} else {
		return "Not decided"
	}

}

func SimpleForLoop() {
	// this is traditional for loop
	for i := 0; i < 10; i++ {
		log.Println("hello", i)
	}

}

func SimpleForLoopAsWhile() {
	count := 0

	// This is acting as while loop
	for count <= 20 {
		log.Println("count", count)
		count++
	}

}

func Infiniteloop() {
	count := 0

loop:
	for {
		log.Println("Infinite loop", count)
		if count > 40 {
			break loop
		}
		count++

	}

}
