package main

import (
	"context"
	"fmt"
	"log"
	"machine"

	keyboard "github.com/sago35/tinygo-keyboard"
	"github.com/sago35/tinygo-keyboard/keycodes/jp"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	d := keyboard.New()

	colPins := []machine.Pin{
		machine.BCM5,
		machine.BCM6,
	}

	rowPins := []machine.Pin{
		machine.PA04, // BCM13,
		machine.BCM19,
		machine.BCM26,
	}

	mk := d.AddMatrixKeyboard(colPins, rowPins, [][][]keyboard.Keycode{
		{
			{jp.KeyT, jp.KeyI},
			{jp.KeyN, jp.KeyY},
			{jp.KeyG, jp.KeyO},
		},
	}, keyboard.InvertDiode(true))

	mk.SetCallback(func(layer, row, col int, state keyboard.State) {
		fmt.Printf("mk: %d %d %d %d\n", layer, row, col, state)
	})

	gpioPins := []machine.Pin{
		machine.WIO_KEY_A,
		machine.WIO_KEY_B,
		machine.WIO_KEY_C,
		machine.WIO_5S_UP,
		machine.WIO_5S_LEFT,
		machine.WIO_5S_RIGHT,
		machine.WIO_5S_DOWN,
		machine.WIO_5S_PRESS,
	}

	for c := range gpioPins {
		gpioPins[c].Configure(machine.PinConfig{Mode: machine.PinInput})
	}

	// KeyMediaXXX will be supported starting with tinygo-0.28.
	gk := d.AddGpioKeyboard(gpioPins, [][][]keyboard.Keycode{
		{
			{jp.KeyA, jp.KeyB, jp.KeyC, jp.KeyMediaVolumeInc, jp.KeyMediaPrevTrack, jp.KeyMediaNextTrack, jp.KeyMediaVolumeDec, jp.KeyMediaPlayPause},
		},
	})

	gk.SetCallback(func(layer, row, col int, state keyboard.State) {
		fmt.Printf("gk: %d %d %d %d\n", layer, row, col, state)
	})

	d.Debug = true
	return d.Loop(context.Background())
}
