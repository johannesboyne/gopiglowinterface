package main

import (
	"log"

	"github.com/wjessop/go-piglow"
)

type Glow struct {
	p *piglow.Piglow
}

func (g *Glow) NewGlow() {
	var err error
	g.p, err = piglow.NewPiglow()
	if err != nil {
		log.Fatal("Couldn't create a Piglow: ", err)
	}
}

func (g *Glow) TrySet() {
	var err error
	err = g.p.Apply()
	if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}
}

func (g *Glow) Green(brightness uint8) {
	g.p.SetLED(3, brightness)
	g.p.SetLED(5, brightness)
	g.p.SetLED(13, brightness)
}

func (g *Glow) Red(brightness uint8) {
	g.p.SetLED(0, brightness)
	g.p.SetLED(6, brightness)
	g.p.SetLED(17, brightness)
}

func (g *Glow) Orange(brightness uint8) {
	g.p.SetLED(2, brightness)
	g.p.SetLED(8, brightness)
	g.p.SetLED(15, brightness)
}

func (g *Glow) Blue(brightness uint8) {
	g.p.SetLED(4, 255)
	g.p.SetLED(14, 255)
	g.p.SetLED(11, 255)
}

func (g *Glow) Clear() {
	for i := 0; i < 18; i++ {
		g.p.SetLED(int8(i), 0)
	}
}

func main() {
	clear(p)

}
