package main

import "fmt"

// A Task has a slice of hours it will occur, another one for minutes it will occur
// and the action it will try to take when it does.
type Task struct {
	hours   []int
	minutes []int
	action  string
}

// Next returns the next time the Task will occur and whether that will be today or tomorrow,
// according to the given parameters for the current hour and minute.
func (t Task) Next(currentHour int, currentMinute int) string {
	m, loopedMinute := findGreaterOrEqualInLooping(currentMinute, t.minutes)
	if loopedMinute {
		currentHour++
		m = findGreaterOrEqualIn(0, t.minutes)
	}

	h, loopedHour := findGreaterOrEqualInLooping(currentHour, t.hours)
	var day string
	if loopedHour {
		day = "tomorrow"
	} else {
		day = "today"
	}

	if h != currentHour {
		m = findGreaterOrEqualIn(0, t.minutes)
	}

	return fmt.Sprintf("%d:%02d %v - %v", h, m, day, t.action)
}
