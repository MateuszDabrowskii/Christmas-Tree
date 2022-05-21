package main

import (
	"fmt"
	"math/rand"
)

type Choinka struct {
	h        int
	gwiazdki string
	s        int
	i        int
}

func (c *Choinka) print() {
	fmt.Print(c.gwiazdki)
}

func (c *Choinka) rysujZnaki(znak string) {
	for k := 0; k < c.h-c.i+c.s; k++ {
		c.gwiazdki += znak
	}
}

func (c *Choinka) rysuj() {
	for ; c.i < c.h; c.i++ {
		c.gwiazdki += "\n"
		c.rysujZnaki(" ")
		for k := 0; k < 2*c.i+1; k++ {
			if (k+c.s+c.i)%c.h == 1 {
				c.gwiazdki += "0"
			} else if r := rand.Intn(8); r == 0 {
				c.gwiazdki += "+"
			} else if r := rand.Intn(16); r == 0 {
				c.gwiazdki += "#"
			} else {
				c.gwiazdki += "*"
			}

		}
		c.rysujZnaki(" ")
	}
	c.print()
}
func main() {
	var wysokość int
	fmt.Scanln(&wysokość)
	var choinka Choinka
	choinka.h = wysokość
	choinka.s = 5
	choinka.rysuj()
	choinka2 := Choinka{
		wysokość + 3,
		"",
		2,
		3}
	choinka2.rysuj()
	choinka3 := Choinka{
		wysokość + 5,
		"",
		0,
		5}
	choinka3.rysuj()
}
