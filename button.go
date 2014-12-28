package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// Button states
const (
	LID_CLOSED     byte = 21
	BUTTON_PRESSED byte = 22
	LID_OPEN       byte = 23
)


// We read button state in a go routine and send it back via a channel
var c = make(chan byte)

var buf = make([]byte, 8)

// Gets the button state, called in a go routine.
func getState(button *os.File) {

	for i := range buf {
		buf[i] = 0
	}
	buf[0] = 0x08
	buf[7] = 0x02

	// reset button state
	_, err := button.Write(buf)
	if err != nil {
		panic(err)
	}

	// read button state
	_, err = button.Read(buf)
	if err != nil {
		panic(err)
	}
	c <- buf[0]
}

// I have another Go server running in the garage that can open the door.
// Not using it currently.
func openIt() {
	res, _ := http.Get("http://192.168.1.78:2929/test")
	output, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("garage: %s\n", output)
}

func suspend(){
        // Using the button to suspend my laptop 
	cmd := exec.Command("/usr/local/bin/suspend.sh")
	cmd.Start()
}

func main() {

	prior := LID_CLOSED

	button, err := os.OpenFile("/dev/big_red_button", os.O_RDWR, 666)
	if err != nil {
		panic(err)
	}

	for { // ever
		go getState(button)

		// Wait for button state from channel
		current := <-c
		if prior != current {
			switch {
			case current == LID_OPEN && prior == LID_CLOSED:
				fmt.Println("Lid opened, ready to fire.")
			case current == LID_CLOSED && prior != BUTTON_PRESSED:
				fmt.Println("Lid closed")
			case current == BUTTON_PRESSED && prior != BUTTON_PRESSED:
				fmt.Println("Fire!")
				go suspend()
				//go openIt()
			case current == LID_OPEN && prior == BUTTON_PRESSED:
				// they Fired.
				fmt.Println("Torpedo away!")
			default:
				fmt.Printf("Something else %d\n", current)
			}
			prior = current
		}
		time.Sleep(100 * time.Millisecond)
	}

}
