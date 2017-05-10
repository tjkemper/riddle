package main

import (
	"flag"
	"fmt"
	//  "os"
	//  "strings"
)

func main() {

	idPtr := flag.Int("id", 0, "Access riddle with id.")

	rPtr := flag.Bool("r", false, "")
	randomPtr := flag.Bool("random", true, "Random riddle.")

	aPtr := flag.Bool("a", false, "")
	answerPtr := flag.Bool("answer", false, "Show answer.")
	flag.Parse()

	fmt.Printf("id: %d, random: %t %t, answer: %t %t\n", *idPtr, *rPtr, *randomPtr, *aPtr, *answerPtr)

}
