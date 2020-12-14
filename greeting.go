package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	// TODO Put these in a file and read in from init()
	format = []string {
		"Hi, %v.",
		"Hello %v!",
		"Oh. It's you. %v."}
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("name is empty")
	}

	message := fmt.Sprintf(getRandomFormat(), name)
	return message, nil
}

func setSeed() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomFormat() string {
	return format[rand.Intn(len(format))]
}

func init() {
	setSeed()
}
