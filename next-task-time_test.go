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
