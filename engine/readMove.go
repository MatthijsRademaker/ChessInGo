package main

import (
	"bufio"
	"chess/domain/pieces"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadMove reads the original and destination positions from standard input
func ReadMove() (pieces.Position, pieces.Position, error) {
	fmt.Println("Enter your move (e.g., 'A1 B2'):")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return pieces.Position{}, pieces.Position{}, err
	}
	// Convert the input string to uppercase and remove leading/trailing whitespaces
	input = strings.TrimSpace(strings.ToUpper(input))

	// Split the input into two parts: original position and destination position
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		return pieces.Position{}, pieces.Position{}, fmt.Errorf("invalid input: expected two positions separated by a space")
	}

	// Parse original position
	originalPos, err := parsePosition(parts[0])
	if err != nil {
		return pieces.Position{}, pieces.Position{}, err
	}

	// Parse destination position
	destPos, err := parsePosition(parts[1])
	if err != nil {
		return pieces.Position{}, pieces.Position{}, err
	}

	return originalPos, destPos, nil
}

// parsePosition parses a position string (e.g., "A1") into a Position struct
func parsePosition(posStr string) (pieces.Position, error) {
	// Validate input length
	if len(posStr) != 2 {
		return pieces.Position{}, fmt.Errorf("invalid position format: %s", posStr)
	}

	// Parse the input characters into X and Y coordinates
	x := int(posStr[0] - 'A')
	y, err := strconv.Atoi(string(posStr[1]))
	if err != nil {
		return pieces.Position{}, fmt.Errorf("invalid position format: %s", posStr)
	}
	// Convert y from 1-based to 0-based indexing
	y--

	// Validate coordinates
	if x < 0 || x > 7 || y < 0 || y > 7 {
		return pieces.Position{}, fmt.Errorf("invalid position: coordinates out of range")
	}

	// adjust x and y to match the board's 0-based indexing
	return pieces.Position{x, y}, nil
}
