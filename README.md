# About
This is an interactive program that I use to practice simple golang stuff.

# Usage
1. Have a recent version of golang installed. I wrote this on 1.8.
2. Run `go get github.com/jackklika/go-sandbox`
3. Run `go-sandbox` to see what projects are available.

# Projects
## streamtest
This is a project to test golang channels. It creates a channel of many random uint8 which are averaged and displayed.
Run this with `go-sandbox streamtest 300` or another uint32 value. The argument to streamtest is the milisecond sleep during the loop.
Run `go-sandbox streamtest 0` for no limit!

### TODO
Currently the average is calculated with an ever-growing sum. There has to be a way to do a running average.
