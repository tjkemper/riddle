package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	JSONFile string = os.Getenv("GOPATH") + "/src/github.com/tjkemper/riddle/riddles.json"
)

// Riddle to keep us from panicking
type Riddle struct {
	ID       int
	Question string
	Answer   string
}

// GetRiddleByID returns a Riddle by ID
func GetRiddleByID(id int) (Riddle, bool) {
	m := parseJson(JSONFile)
	riddle, ok := m[strconv.Itoa(id)]
	return riddle, ok
}

// GetRandomRiddle returns a random Riddle
func GetRandomRiddle() Riddle {
	m := parseJson(JSONFile)
	i := getRandomInt(len(m))
	count := 0
	for _, riddle := range m {
		if count == i {
			return riddle
		}
		count++
	}
	return Riddle{}
}

func parseJson(filename string) map[string]Riddle {

	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	//m := make(map[string]Riddle)
	var m map[string]Riddle
	json.Unmarshal(file, &m)
	return m
}

func getRandomInt(n int) int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(n)
}
