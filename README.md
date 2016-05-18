# Next Scheduled Task Time

A modest command line program to parse a "cron-like" tab and see when the scheduled tasks are supposed to run next.

Intended to flex those [golang] muscles!

## Installation
If you have Go installed already, installing this is as simple as:

`go install github.com/jmagrippis/next-task-time`

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

Press <kbd>Ctrl</kbd>+<kbd>D</kbd> to stop capturing input and get your results printed out!

[golang]: https://golang.org/ "Makes it easy to build simple, reliable, and efficient software."
[GoConvey]: http://goconvey.co/ "Write behavioral tests in your editor. Get live results in your browser."

## Testing
Run the test suite in "vanilla" mode with `go test` in the cloned project's folder.

Run it in "cool cat" mode by installing [GoConvey] and running `$GOPATH/bin/goconvey`!

## Missing features

Still in development, barely any code here! Main missing features:

- Actually parse input in the HH:MM/H:MM format.
- Actually capture input for the "cron-style" lines.
- Support `--help` flag
- Support `--version` flag
