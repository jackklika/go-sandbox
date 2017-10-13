// Testing simple streams
// This program will create a stream of random numbers,
//  then create a running average.
// It should be pretty efficient!
package main

import(
	"fmt"
	"time"
	"strconv"
	"math/rand"
)

var intchan = make(chan uint8)

func streamtest(args ...string) {
	ms, err := strconv.ParseInt(args[0], 0, 32)
	if err != nil {
		fmt.Println("Error: Invalid argument")
		return
	}

	go makestreamandfill(int32(ms))
	var cnt uint64 = 0
	var avg float64 = 0
	var sum uint64 = 0
	var recnum uint8 = 0

	for {
		select {
		case recnum = <-intchan:

			// Converting the recnum to a u64 int is probably lighter
			//  than generating random u64 numbers. But IDK!
			sum += uint64(recnum)
			cnt++
			avg = float64(sum)/float64(cnt)
			fmt.Printf("\tAVG: %3.3f\t%3.14f%% to max uint64\n", avg, float64(sum)/float64(^uint64(0)))
		}
	}

}

func makestreamandfill(ms int32) {

	var randnum uint8 //0 - 255
	for {
		randnum = uint8(rand.Intn(255))
		time.Sleep(time.Duration(ms) * time.Millisecond)
		fmt.Printf("%3d -> stream", randnum)
		intchan <- randnum;
	}

}
