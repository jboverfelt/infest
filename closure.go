package main

import (
	"fmt"
	"strings"
	"time"
)

// Closure represents and individual establishment closure
type Closure struct {
	Name        string
	Address     string
	Reason      string
	ClosureDate ClosureTime
	ReopenDate  *ClosureTime
}

const closuresTimeFmt = "01/02/2006"

// ClosureTime is the time format used by the restaurant closures page
type ClosureTime struct {
	time.Time
}

func (c ClosureTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", c.Time.Format(closuresTimeFmt))

	return []byte(stamp), nil
}

func (c *ClosureTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		c.Time = time.Time{}
		return nil
	}
	tt, err := time.Parse(closuresTimeFmt, s)

	if err != nil {
		return err
	}

	c.Time = tt
	return nil
}
