package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d", c.total)
}

func main() {
	var c Counter
	c.Increment()
	c.Increment()
	fmt.Println(c)
}
