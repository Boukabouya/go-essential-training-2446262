package main

import (
	"fmt"
	"log"
)

// Square is a square
type Square struct {
	// TODO
	x, y int
}

// NewSquare returns a new square
func NewSquare(x int, y int, length int) (*Square, error) {
	// TODO
	if x < 0 || y < 0 || length <= 0 {
		return nil, fmt.Errorf("Invalid coordinates or size")
	}
	s := Square{
		x: x,
		y: y,
	}
	return &s, nil
}

// Move moves the square
func (s *Square) Move(dx int, dy int) {
	// TODO
	s.x = dx
	s.y = dy
}

// Area returns the square are
func (s *Square) Area() int {
	// TODO
	area := s.x * s.y
	return area
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}

	s.Move(3, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
