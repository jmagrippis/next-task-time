package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindGreaterOrEqualInLooping(t *testing.T) {
	Convey("Given an integer for the current value", t, func() {
		x := 3

		Convey("And a slice of acceptable integers", func() {
			acceptable := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}

			Convey("It returns the acceptable integer closest to the current value, and whether it needed to loop", func() {
				r, looped := findGreaterOrEqualInLooping(x, acceptable)
				So(r, ShouldEqual, 3)
				So(looped, ShouldEqual, false)

				x = 5
				r, looped = findGreaterOrEqualInLooping(x, acceptable)
				So(r, ShouldEqual, 5)
				So(looped, ShouldEqual, false)

				x = 24
				r, looped = findGreaterOrEqualInLooping(x, acceptable)
				So(r, ShouldEqual, 24)
				So(looped, ShouldEqual, false)
			})
		})

		Convey("And a slice of one acceptable integer", func() {
			acceptable := []int{6}

			Convey("It still returns the acceptable integer closest to the current value, and whether it needed to loop", func() {
				r, looped := findGreaterOrEqualInLooping(x, acceptable)
				So(r, ShouldEqual, 6)
				So(looped, ShouldEqual, false)

				x = 5
				r, looped = findGreaterOrEqualInLooping(x, acceptable)
				So(r, ShouldEqual, 6)
				So(looped, ShouldEqual, false)

				x = 24
				r, looped = findGreaterOrEqualInLooping(x, acceptable)
				So(r, ShouldEqual, 6)
				So(looped, ShouldEqual, true)
			})
		})
	})

	Convey("Given a negative integer", t, func() {
		x := -3

		Convey("And a slice of acceptable integers", func() {
			acceptable := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}

			Convey("Should panic with a helpful message for non-negative integers", func() {
				So(func() { findGreaterOrEqualInLooping(x, acceptable) }, ShouldPanicWith, "-3 is not a positive integer!")
				x = -200
				So(func() { findGreaterOrEqualInLooping(x, acceptable) }, ShouldPanicWith, "-200 is not a positive integer!")
			})
		})
	})

	Convey("Given an integer for the current value", t, func() {
		x := 3

		Convey("And a slice of only negative integers", func() {
			acceptable := []int{-7, -6, -5, -4, -3, -2, -1}

			Convey("Should panic with a helpful message for non-negative integers", func() {
				So(func() { findGreaterOrEqualInLooping(x, acceptable) }, ShouldPanicWith, "All acceptable integers given are negative!")
				acceptable = []int{-200}
				So(func() { findGreaterOrEqualInLooping(x, acceptable) }, ShouldPanicWith, "All acceptable integers given are negative!")
			})
		})
	})
}

func TestExtractHourAndMinute(t *testing.T) {
	Convey("Given a string in the format HH:MM", t, func() {
		time := "12:32"

		Convey("Returns an integer for the hour, and an integer for the minute", func() {
			hour, minute := extractHourAndMinute(time)
			So(hour, ShouldEqual, 12)
			So(minute, ShouldEqual, 32)

			time = "07:47"
			hour, minute = extractHourAndMinute(time)
			So(hour, ShouldEqual, 7)
			So(minute, ShouldEqual, 47)
		})
	})

	Convey("Given a string in the format H:MM", t, func() {
		time := "2:32"

		Convey("Returns an integer for the hour, and an integer for the minute", func() {
			hour, minute := extractHourAndMinute(time)
			So(hour, ShouldEqual, 2)
			So(minute, ShouldEqual, 32)

			time = "7:47"
			hour, minute = extractHourAndMinute(time)
			So(hour, ShouldEqual, 7)
			So(minute, ShouldEqual, 47)

			time = "7:47"
			hour, minute = extractHourAndMinute(time)
			So(hour, ShouldEqual, 7)
			So(minute, ShouldEqual, 47)
		})
	})

	Convey("Given a string in the format HH:MM, but with values out of range", t, func() {
		time := "27:32"

		Convey("Panics with a helpful message", func() {
			So(func() { extractHourAndMinute(time) }, ShouldPanicWith, "Unfortunately, there are only 24 hours in a day! You passed 27.")

			time = "17:70"
			So(func() { extractHourAndMinute(time) }, ShouldPanicWith, "There are only 60 minutes in an hour! You passed 70.")

			time = "-7:47"
			So(func() { extractHourAndMinute(time) }, ShouldPanicWith, "Unfortunately, there are only 24 hours in a day! You passed -7.")

			time = "05:-12"
			So(func() { extractHourAndMinute(time) }, ShouldPanicWith, "There are only 60 minutes in an hour! You passed -12.")
		})
	})
}

func TestExtractAcceptableIntegers(t *testing.T) {
	Convey("Given a string representation of an integer and an integer limit", t, func() {
		valuesString := "12"
		limit := 23

		Convey("Returns a slice with only that single integer", func() {
			So(extractAcceptableIntegers(valuesString, limit), ShouldResemble, []int{12})

			valuesString = "1"
			So(extractAcceptableIntegers(valuesString, limit), ShouldResemble, []int{1})
		})
	})

	Convey("Given a string of comma separated integers and an integer limit", t, func() {
		valuesString := "12,16,18"
		limit := 23

		Convey("Returns a slice with the comma separated integers", func() {
			So(extractAcceptableIntegers(valuesString, limit), ShouldResemble, []int{12, 16, 18})

			valuesString = "1,6,8,11,21,22,23,24"
			So(extractAcceptableIntegers(valuesString, limit), ShouldResemble, []int{1, 6, 8, 11, 21, 22, 23, 24})
		})
	})

	Convey("Given a string of comma separated integers in a random order and an integer limit", t, func() {
		valuesString := "18,12,16"
		limit := 23

		Convey("Returns a sorted slice with the comma separated integers", func() {
			So(extractAcceptableIntegers(valuesString, limit), ShouldResemble, []int{12, 16, 18})

			valuesString = "11,21,22,1,23,24,8,6"
			So(extractAcceptableIntegers(valuesString, limit), ShouldResemble, []int{1, 6, 8, 11, 21, 22, 23, 24})
		})
	})

	Convey("Given a string \"*\" and an integer limit", t, func() {
		valuesString := "*"
		limit := 23

		Convey("Returns a slice with the comma separated integers", func() {
			So(extractAcceptableIntegers(valuesString, limit), ShouldResemble, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23})

			limit = 11
			So(extractAcceptableIntegers(valuesString, limit), ShouldResemble, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
		})
	})
}
