// Command-line program next-task-time parses an argument for a simulated time
// and expects a "cron-style" config as input. It returns where the given
// configured tasks will run next
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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

	version := flag.Bool("version", false, "whether to display the version information")
	flag.Parse()

	if *version {
		fmt.Println("Next Scheduled Task Time version 0.1.0")
		os.Exit(0)
	}

	args := os.Args
	if len(args) < 2 {
		panic(fmt.Sprint("You need to pass one argument, a simulated time in the format HH:MM!"))
	}

	simulatedHour, simulatedMinute := extractHourAndMinute(os.Args[1])

	var tasks []Task
	readyToStop := false
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			if readyToStop {
				break
			}
			readyToStop = true
			continue
		} else {
			readyToStop = false
		}
		data := strings.Split(input, " ")

		if len(data) < 3 {
			panic(fmt.Sprint("Badly formatted string: ", input))
		}

		minutes := extractAcceptableIntegers(data[0], 23)
		hours := extractAcceptableIntegers(data[1], 59)

		if len(data) > 3 {
			data[2] = strings.Join(data[2:], " ")
		}

		tasks = append(tasks, Task{hours: hours, minutes: minutes, action: data[2]})
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprint("reading standard input: ", err))
	}

	for _, task := range tasks {
		fmt.Println(task.Next(simulatedHour, simulatedMinute))
	}
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

// extractAcceptableIntegers expects a comma separated integers in string form,
// or "*", at which case it just fills the returned slice with incrementing
// numbers up to the given limit
func extractAcceptableIntegers(valuesString string, limit int) []int {
	var acceptableInts []int
	if valuesString == "*" {
		for i := 0; i <= limit; i++ {
			acceptableInts = append(acceptableInts, i)
		}
		return acceptableInts
	}

	valueStrings := strings.Split(valuesString, ",")
	for _, valueString := range valueStrings {
		value, err := strconv.Atoi(valueString)
		if err != nil {
			panic(fmt.Sprint("Badly formatted input: ", valuesString))
		}
		acceptableInts = append(acceptableInts, value)
	}
	sort.Ints(acceptableInts)
	return acceptableInts
}
