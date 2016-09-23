package main

import (
	"fmt"
	//"strconv"
	"bufio"
	"strings"
	"io/ioutil"
)
import (
	"github.com/tarm/serial"
	"log"
	//"time"
	//"bufio"
	//"bufio"
	//"io/ioutil"
	//"strings"
	//"strconv"
	//"github.com/tarm/serial"
	//"runtime"
	//"path/filepath"
	"os"
	//"os/exec"
)

func main() {

	arg := os.Args[1:][0]

	fmt.Println(arg)

        //serial_port := find_usb_serial()
	serial_port := "/dev/cu.usbmodem1411"
	c := &serial.Config{Name: serial_port, Baud: 19200, ReadTimeout: 20}
	fmt.Printf(c.Name + "\n")
	// log.Printf(c.Name + "\n")


	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Printf("got an error")
		log.Fatal(err)
	}

	s.Flush()

	n, err := s.Write([]byte("relay "+ arg + " 0 \n\r"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("I wrote this number of bytes \n")
	fmt.Printf("%d", n)
	fmt.Printf("\n")
	fmt.Printf("before read  \n")


	reader := bufio.NewReader(s)


	//reply, err := reader.ReadString('\r')
	reply, err := reader.ReadString('\r')
	fmt.Printf(reply)
	if err != nil {
		s.Flush()
		fmt.Printf("error flush  \n")
		fmt.Printf("%d", err)

		} else{
			fmt.Println(reply)



		}

		s.Close()
			}









func find_usb_serial() string {
	contents, _ := ioutil.ReadDir("/dev")

	// Look for what is mostly likely the Arduino device
	for _, f := range contents {
		if strings.Contains(f.Name(), "cu.usbmodem"){
			fmt.Printf("/dev/" + f.Name())
			return "/dev/" + f.Name()
		}
	}

	// Have not been able to find a USB device that 'looks'
	// like an Arduino.
	return ""
}