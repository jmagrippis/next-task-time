package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, "Error:", r)
			os.Exit(1)
		}
	}()

	args := os.Args
	if len(args) < 2 {
		panic(fmt.Sprint("You need to pass one argument, a simulated time in the format HH:MM!"))
	}

	simulatedHour, simulatedMinute := extractHourAndMinute(os.Args[1])

	var tasks []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprint("reading standard input: ", err))
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

// findGreaterOrEqualInLooping calls findGreaterOrEqualIn with the given
// arguments then, if no match can be made, resets the base argument to zero
// and calls it again; so it "loops" through the slice of acceptable integers
// before giving up. Returns the matching integer and whether it needed to loop.
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

// findGreaterOrEqualIn takes an integer and iterates through the given slice
// of acceptable integers to find the one closest to it going forward.
// Returns -1 if no such number can be found.
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

// extractHourAndMinute takes a string in the format of HH:MM
// and returns integers for the hour and the minute.
func extractHourAndMinute(formatted string) (int, int) {
	time := strings.Split(formatted, ":")
	if len(time) < 2 {
		panic(fmt.Sprintf("Invalid argument \"%v\", supposed to be in the HH:MM format", formatted))
	}

	hour, err := strconv.Atoi(time[0])
	if err != nil {
		panic(fmt.Sprint("Something improbable went wrong with your input!"))
	}
	if 0 > hour || hour > 24 {
		panic(fmt.Sprintf("Unfortunately, there are only 24 hours in a day! You passed %d.", hour))
	}

	minute, err := strconv.Atoi(time[1])
	if err != nil {
		panic(fmt.Sprint("Something improbable went wrong with your input!"))
	}
	if 0 > minute || minute > 60 {
		panic(fmt.Sprintf("There are only 60 minutes in an hour! You passed %d.", minute))
	}

	return hour, minute
}
