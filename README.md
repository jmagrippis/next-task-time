# Next Scheduled Task Time

A modest command line program to parse a "cron-like" tab and see when the scheduled tasks are supposed to run next.

Intended to flex those [golang] muscles!

## Installation
If you have Go installed already, installing this is as simple as:

`go get github.com/jmagrippis/next-task-time`

If you don't have Go installed... Well, [it's never too late](https://golang.org/dl/)!

## Usage
Run it with:

`$GOPATH/bin/next-task-time <HH:MM>`

Where `<HH:MM>` is the desired simulated time. Like `16:10`, or `04:45` or even `1:23`.

Standard input will then be expecting a few "cron-style" lines, like:
```
30 1 /bin/run_me_daily
45 * /bin/run_me_hourly
* * /bin/run_me_every_minute
* 19 /bin/run_me_sixty_times
```

Press <kbd>Ctrl</kbd>+<kbd>D</kbd>, or just <kbd>enter</kbd> a few times to leave two empty lines, to stop capturing input and get your results printed out!

You may also pipe a text file as input. Something like:

```
$GOPATH/bin/next-task-time 16:10 < sample-schedule.txt
```

## Testing
Run the test suite in "vanilla" mode with `go test` in the cloned project's folder.

Run it in "cool cat" mode by installing [GoConvey] and running `$GOPATH/bin/goconvey`!

## Wishlist
As mentioned previously, this is a pretty modest program, so most of the functionality has already been written.
That said, there are a few things that might be worth looking at:

- Coloured output
- Actually helpful help text
- Extract testable method for getting the tasks
- Separate library for some of the utility functions

[golang]: https://golang.org/ "Makes it easy to build simple, reliable, and efficient software."
[GoConvey]: http://goconvey.co/ "Write behavioral tests in your editor. Get live results in your browser."
