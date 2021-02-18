package main

import (
	"flag"
	"log"
	"time"

	"github.com/rakyll/portmidi"
)

var out = flag.Int("out", -1, "use this device to test MIDI output")

func main() {
	flag.Parse()

	err := portmidi.Initialize()
	if err != nil {
		panic(err)
	}
	defer portmidi.Terminate()

	if *out == -1 {

		for d := 0; d < portmidi.CountDevices(); d++ {
			i := portmidi.Info(portmidi.DeviceID(d))
			if i.IsInputAvailable {
				log.Println(d, "INPUT", i.Interface, i.Name)
			}
			if i.IsOutputAvailable {
				log.Println(d, "OUTPUT", i.Interface, i.Name)
			}
		}

	} else {
		log.Println("open output stream")
		midiout, err := portmidi.NewOutputStream(portmidi.DeviceID(*out), 256, 0)
		if err != nil {
			log.Println("error creating output stream", err)
			return
		}
		log.Println("write C on")
		midiout.WriteShort(0x90, 60, 100)
		log.Println("wait 2 seconds")
		time.Sleep(2 * time.Second)
		log.Println("write C off")
		midiout.WriteShort(0x80, 60, 100)
		log.Println("close output stream")
		midiout.Close()
	}
}
