package main

// Riddle represents a riddle object
type Riddle struct {
	ID       int
	Question string
	Answer   string
}

// GetRiddleByID returns a Riddle by ID
func GetRiddleByID(id int) Riddle {
	return Riddle{
		ID:       id,
		Question: "what is go?",
		Answer:   "awesome",
	}
}

// GetRandomRiddle returns a random Riddle
func GetRandomRiddle() Riddle {
	return Riddle{
		ID:       42,
		Question: "is this random?",
		Answer:   "no",
	}
}
