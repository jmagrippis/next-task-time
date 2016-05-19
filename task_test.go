package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNext(t *testing.T) {
	Convey("Given a Task", t, func() {
		task := Task{hours: []int{1}, minutes: []int{30}, action: "/bin/run_me_daily"}

		Convey("And a simulated current time", func() {
			h := 16
			m := 10

			Convey("Returns the soonest time it can run, whether it is today or tomorrow, and the action it will take", func() {
				So(task.Next(h, m), ShouldEqual, "1:30 tomorrow - /bin/run_me_daily")

				h = 1
				m = 22
				So(task.Next(h, m), ShouldEqual, "1:30 today - /bin/run_me_daily")
			})

		})
	})

	Convey("Given a more complicated Task", t, func() {
		task := Task{hours: []int{1, 2, 3, 4, 6, 7}, minutes: []int{20, 25, 30, 35}, action: "/bin/run_me_with parameters please"}

		Convey("And a simulated current time", func() {
			h := 16
			m := 10

			Convey("Returns the soonest time it can run, whether it is today or tomorrow, and the action it will take", func() {
				So(task.Next(h, m), ShouldEqual, "1:20 tomorrow - /bin/run_me_with parameters please")

				h = 1
				m = 22
				So(task.Next(h, m), ShouldEqual, "1:25 today - /bin/run_me_with parameters please")
			})

		})
	})
}
