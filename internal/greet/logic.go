package greet

import (
	"errors" // For creating custom errors
	"fmt"
	"strings" // For strings.Builder
)

// GenerateGreeting creates a greeting string based on the provided parameters.
func GenerateGreeting(name string, formal bool, times int) (string, error) {
	// Input validation
	if times <= 0 {
		return "", errors.New("number of times must be positive")
	}
	if name == "" {
		name = "World" // Default name
	}

	// Determine the greeting prefix
	greetingPrefix := "Hello"
	if formal {
		greetingPrefix = "Greetings"
	}

	// Construct the single greeting message
	singleGreeting := fmt.Sprintf("%s, %s!", greetingPrefix, name)

	// If only one repetition is needed, return directly
	if times == 1 {
		return singleGreeting, nil
	}

	// Use strings.Builder for efficient string concatenation in a loop
	var sb strings.Builder
	for i := 0; i < times; i++ {
		sb.WriteString(singleGreeting)
		// Add a newline between repetitions, but not after the last one
		if i < times-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil
}
