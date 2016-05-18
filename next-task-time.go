package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
			os.Exit(1)
		}
	}()

	args := os.Args
	if len(args) < 2 {
		panic(fmt.Sprint("You need to pass one argument, a simulated time in the format HH:MM!"))
	}
	if len(args) < 3 {
		panic(fmt.Sprint("Just for development, you need to pass two arguments, the hour and the minute!"))
	}
	simulatedHour, err := strconv.Atoi(os.Args[1])
	simulatedMinute, err := strconv.Atoi(os.Args[2])

	if err != nil {
		panic(fmt.Sprint("Something improbable went wrong with your input!"))
	}

	m, loopedMinute := findGreaterOrEqualInLooping(simulatedMinute, []int{20, 30, 40})
	if loopedMinute {
		simulatedHour++
	}

	h, loopedHour := findGreaterOrEqualInLooping(simulatedHour, []int{2, 3, 4})
	var day string
	if loopedHour {
		day = "tomorrow"
	} else {
		day = "today"
	}

	fmt.Printf("%d:%2d %v -\n", h, m, day)
}

func findGreaterOrEqualInLooping(base int, acceptable []int) (int, bool) {
	looped := false
	closest := findGreaterOrEqualIn(base, acceptable)
	if closest == -1 {
		looped = true
		closest = findGreaterOrEqualIn(0, acceptable)
	}
	if closest == -1 {
		panic(fmt.Sprintf("All acceptable integers given are negative!"))
	}
	return closest, looped
}

func findGreaterOrEqualIn(base int, acceptable []int) int {
	if base < 0 {
		panic(fmt.Sprintf("%v is not a positive integer!", base))
	}
	for _, value := range acceptable {
		if value >= base {
			return value
		}
	}
	return -1
}
