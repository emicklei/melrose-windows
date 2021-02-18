package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"gitlab.com/gomidi/rtmididrv/imported/rtmidi"
	// when using portmidi, replace the line above with
	// driver gitlab.com/gomidi/portmididrv
)

var port = flag.Int("p", -1, "port number for out")

func main() {
	flag.Parse()

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
		fmt.Println(i, name)
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
			fmt.Println(i, name)
		}
	}
	if *port == -1 {
		return
	}

	out, err := rtmidi.NewMIDIOutDefault()
	if err != nil {
		log.Fatalln("can't open default MIDI out: ", err)
	}
	defer out.Close()

	err = out.OpenPort(*port, "")
	if err != nil {
		log.Fatalln("can't open default MIDI out: ", err)
	}

	out.SendMessage([]byte{0x90, 60, 60})
	time.Sleep(1 * time.Second)
	out.SendMessage([]byte{0x80, 60, 60})
}
