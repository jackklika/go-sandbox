package main

import (
	"fmt"
	"os"
)

func main() {

	arglen := len(os.Args)
	if arglen == 1 {
		fmt.Println("Enter an argument, either an index or name of the program.")
		fmt.Println(`
1: streamtest -- A stream of random numbers are generated and averaged in real-time.
		`)
		return
	}

	progname := os.Args[1]
	if arglen == 2 {
		fmt.Printf("EXECUTING PROGRAM %s\n", progname);
	}

	if arglen >= 3 {
		args := os.Args[2:]
		fmt.Printf("EXECUTING PROGRAM %s WITH ARGS %s\n", progname, args);
	}

	switch progname {
	case "streamtest":
		streamtest(os.Args[2:][0])
	default:
		fmt.Printf("%s not found", progname);
	}

}
