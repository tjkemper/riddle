package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/colorstring"
)

var (
	version bool
	id      int
	random  bool
	answer  bool
)

func init() {
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&version, "v", false, "print version and exit (shorthand)")
	flag.IntVar(&id, "id", -1, "print riddle with id")
	flag.IntVar(&id, "i", -1, "print riddle with id (shorthand)")
	flag.BoolVar(&random, "random", true, "print random riddle")
	flag.BoolVar(&random, "r", true, "print random riddle (shorthand)")
	flag.BoolVar(&answer, "answer", false, "show answer")
	flag.BoolVar(&answer, "a", false, "show answer (shorthand)")

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {

	if version {
		fmt.Println("riddle version 0.1")
		return
	}

	if !random && id == -1 {
		usageAndExit("if -random=false then -id must be set", 0)
	}

	var riddle Riddle
	if id == -1 {
		riddle = GetRandomRiddle()
	} else {
		var foundRiddle bool
		riddle, foundRiddle = GetRiddleByID(id)
		if !foundRiddle {
			usageAndExit(fmt.Sprintf("id of %d does not exist", id), 0)
		}
	}

	printRiddle(riddle, answer)

}

func printRiddle(riddle Riddle, answer bool) {
	fmt.Println(colorstring.Color("\n[cyan]" + "Riddle"))
	fmt.Println(colorstring.Color("[cyan]" + "======"))
	fmt.Println(riddle.Question + "\n\n")

	if answer {
		fmt.Println(colorstring.Color("[green]" + "Answer"))
		fmt.Println(colorstring.Color("[green]" + "======"))
		fmt.Println(riddle.Answer + "\n")
	} else {
		fmt.Println(colorstring.Color("[yellow]" + "Get answer"))
		fmt.Println(colorstring.Color("[yellow]" + "=========="))
		fmt.Printf("riddle -id %d -answer\n\n", riddle.ID)
	}
}

func printError(err error) {
	fmt.Println(colorstring.Color("[red]" + err.Error()))
	os.Exit(1)
}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, colorstring.Color("[yellow]"+message))
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}
