package gometr

import (
	"fmt"
	// "sync"
	"time"
)

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetID() string
	Health() bool
}

type Checker struct {
	items []Checkable
}

func (c *Checker) Add(item Checkable) {
	c.items = append(c.items, item)
}

func (c Checker) String() string {
	var str string
	for _, val := range c.items {
		str += val.GetID()
		str += "\n "
	}

	return str
}

func (c *Checker) Run() {
	ticker := time.NewTicker(5 * time.Second)
	for tick := range ticker.C {
		for _, val := range c.items {
			go check(tick, val)

		}
	}

}

func check(tick time.Time, client Checkable) {
	if !client.Health() {
		fmt.Printf("%s не работает, время: %v\n", client.GetID(), tick)
	}
}

func (c *Checker) Check() {
	for _, val := range c.items {
		if !val.Health() {
			fmt.Printf("%s не работает\n", val.GetID())
		}
	}
}
