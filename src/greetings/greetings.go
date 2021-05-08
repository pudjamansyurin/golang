package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty string")
	}
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}

// Hellos return greetings for multiple person.
func Hellos(names []string) (map[string]string, error) {
	msgs := make(map[string]string)
	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		msgs[name] = msg
	}
	return msgs, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hello %v",
		"Ohayou %v",
		"Salam %v",
	}

	return formats[rand.Intn(len(formats))]
}
