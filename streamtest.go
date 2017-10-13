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
	var avg uint64 = 0
	var sum uint64 = 0
	var recnum uint8 = 0
	fmt.Println(recnum)

	for {
		select {
		case recnum := <-intchan:
			sum += uint64(recnum)
			cnt++
			avg = (sum/cnt)
			fmt.Printf("\tAVG: %d\n", avg)
		}
	}

}

func makestreamandfill(ms int32) {

	var randnum uint8 //0 - 255
	for {
		randnum = uint8(rand.Intn(255))
		time.Sleep(time.Duration(ms) * time.Millisecond)
		fmt.Printf("%d --> stream", randnum)
		intchan <- randnum;
	}

}
