package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func incFmt() {
	fmt.Println("")
}

type murphy struct{}

func (m *murphy) law() (bool, error) {
	if time.Now().Second()%2 == 0 {
		return true, errEven
	}
	return true, errOdd
}

var (
	errEven = errors.New("expected failure")
	errOdd  = errors.New("unexpected failure")
)

// START OMIT
func main() {
	m := &murphy{}
	if b, err := m.law(); err != nil {
		log.Println(b, err)
	}
	// fmt.Println(b)
}

// END OMIT
