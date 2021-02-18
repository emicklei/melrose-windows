package main

import (
	"fmt"
	"log"

	"gitlab.com/gomidi/rtmididrv/imported/rtmidi"
	// when using portmidi, replace the line above with
	// driver gitlab.com/gomidi/portmididrv
)

func main() {
	in, err := rtmidi.NewMIDIInDefault()
	if err != nil {
		log.Fatalln("can't open default MIDI in: ", err)
	}
	defer in.Close()
	ports, err := in.PortCount()
	if err != nil {
		log.Fatalln("can't get number of in ports: ", err.Error())
	}
	for i := 0; i < ports; i++ {
		name, err := in.PortName(i)
		if err != nil {
			name = ""
		}
		fmt.Println(name)
	}
	{
		// Outs
		out, err := rtmidi.NewMIDIOutDefault()
		if err != nil {
			log.Fatalln("can't open default MIDI out: ", err)
		}
		defer out.Close()
		ports, err := out.PortCount()
		if err != nil {
			log.Fatalln("can't get number of out ports: ", err.Error())
		}

		for i := 0; i < ports; i++ {
			name, err := out.PortName(i)
			if err != nil {
				name = ""
			}
			fmt.Println(name)
		}
	}
}
