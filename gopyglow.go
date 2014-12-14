package main

import (
	"github.com/wjessop/go-piglow"
	"log"
)

func main() {
	var p *piglow.Piglow
	var err error

	// Create a new Piglow
	p, err = piglow.NewPiglow()
	if err != nil {
		log.Fatal("Couldn't create a Piglow: ", err)
	}

	// green(p, 100)
	// red(p, 100)
	// orange(p, 100)
	// blue(p, 100)
  clear(p)
	err = p.Apply()
	if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}
}

func green(p *piglow.Piglow, brightness uint8) {
	p.SetLED(3, brightness)
	p.SetLED(5, brightness)
	p.SetLED(13, brightness)
}

func red(p *piglow.Piglow, brightness uint8) {
	p.SetLED(0, brightness)
	p.SetLED(6, brightness)
	p.SetLED(17, brightness)
}

func orange(p *piglow.Piglow, brightness uint8) {
	p.SetLED(2, brightness)
	p.SetLED(8, brightness)
	p.SetLED(15, brightness)
}

func blue(p *piglow.Piglow, brightness uint8) {
  p.SetLED(4, 255)
	p.SetLED(14, 255)
	p.SetLED(11, 255)
}

func clear(p *piglow.Piglow) {
  for i := 0; i < 18; i++ {
    p.SetLED(int8(i), 0)
  }
}
