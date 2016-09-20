package main

import "fmt"
import (
	//"github.com/tarm/serial"
	"log"

	//"time"
	//"bufio"
	//"bufio"
	//"io/ioutil"
	//"strings"
	//"strconv"
	//"github.com/digitallumens/gateway3/gw3config/gw3usb"
//	"strconv"
//	"bufio"
//	"strings"
//	"io/ioutil"
//	"errors"
	"io/ioutil"
	//"strings"
	"strings"
	//"github.com/tarm/serial"
	//"syscall"
	//"errors"
	"github.com/tarm/serial"
	"github.com/digitallumens/dllibgo"

)

func main() {
	//serial_port := find_usb_modem()
	//fmt.Printf(serial_port)
	//c := &serial.Config{Name: serial_port, Baud: 115200, ReadTimeout: 20}
	//fmt.Printf(c.Name + "\n")
	files, err := ioutil.ReadDir("/dev")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//fmt.Println(file.Name())
		if strings.Contains(file.Name(), "cu.SLAB_USBtoUART") {
			//fmt.Printf(file.Name() + "\n")


			c := &serial.Config{Name: file.Name(), Baud: 115200, ReadTimeout: 20}
			fmt.Printf("c.Name = \n")
			fmt.Printf("%s\n", c.Name)
			g, err := dllibgo.NewEZSPGateway("/dev/cu.SLAB_USBtoUART")
			if err != nil {
				fmt.Printf("NewEZSPGateway failed: %s", err)
			}
			g.JoinNetwork("B16")

			//s, err := serial.OpenPort(file)
			if err != nil {
				fmt.Printf("got an error")
				log.Fatal(err)
			}
		}
	}




	//log.Printf(c.Name + "\n")

	//
	//s, err := serial.OpenPort(c)
	//if err != nil {
	//	fmt.Printf("got an error")
	//	log.Fatal(err)
	//}
	//
	//s.Flush()

	//func find_usb_modem(){
	//		files, _ := ioutil.ReadDir("/dev")
	//		for _, f := range files {
	//			if strings.Contains(f.Name(), "SLAB_USBtoUART") {
	//				fmt.Printf( "/dev/" + f.Name())
	//				return f
	//		}
	//	}
	//	return errors.New("USB device could not be found")
	//	}

	//func discoverUSBModem {
	//	files := ioutil.ReadDir("/dev")
	//	mylen := len(files)
	//	for mylen, f := range files {
	//		if strings.Contains(f.Name(), "SLAB_USBtoUART") {
	//			fmt.Printf(f.Name())
	//			//client.Dev = "/dev/" + f.Name()
	//			return nil
	//		}
	//	}
	//	return errors.New("USB device could not be found")
	//}
	//


	//        counter := 0
	//	for {
	//		fmt.Printf("loop counter = \n")
	//		str_counter  := strconv.Itoa(counter)
	//		fmt.Printf(str_counter + "\n")
	//		counter++
	//		//
	//		n, err := s.Write([]byte("G0002\r"))
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Printf("I wrote this number of bytes \n")
	//		fmt.Printf("%d", n)
	//		fmt.Printf("\n")
	//		fmt.Printf("before read  \n")
	//
	//		reader := bufio.NewReader(s)
	//
	//
	//		//reply, err := reader.ReadString('\r')
	//		reply, err := reader.ReadString('\r')
	//		fmt.Printf(reply)
	//		if err != nil {
	//			s.Flush()
	//			fmt.Printf("error flush  \n")
	//			fmt.Printf("%d", err)
	//
	//			} else{
	//				fmt.Println(reply)
	//				split_reply := strings.Split(reply,": ")
	//				fmt.Printf("command = " + split_reply[0] + "\n")
	//				fmt.Printf("response = " + split_reply[1] + "\n")
	//				if strings.Contains(split_reply[0], "ERROR")  {
	//					s.Flush()
	//					fmt.Printf("ERROR: flush  \n")
	//					fmt.Printf("%d", err)
	//					continue
	//
	//
	//
	//			}
	//						s.Close()
	//						break
	//				}
	//
	//		}
	//	}
	//
	//
	//
	//	//
	//
	//	//buf := make([]byte, 128)
	//	//
	//	//time.Sleep(3)
	//	//
	//	//
	//	//
	//	//n, err = s.Read(buf)
	//	//if err != nil {
	//	//	fmt.Printf("error in read\n")
	//	//	fmt.Printf("%q", buf[:16])
	//	//	log.Fatal(err)
	//	//}
	//	//fmt.Printf("after read \n")
	//	//fmt.Printf(string(n))
	//	////log.Printf("%q", buf[:n])
	//	////fmt.Printf("%q", buf[:16])
	//	//fmt.Printf("%q", buf)
	//
	//
	//func find_usb_serial() string {
	//	contents, _ := ioutil.ReadDir("/dev")
	//
	//	// Look for what is mostly likely the Arduino device
	//	for _, f := range contents {
	//		if strings.Contains(f.Name(), "cu.usbserial"){
	//			fmt.Printf("/dev/" + f.Name())
	//			return "/dev/" + f.Name()
	//		}
	//	}
	//
	//	// Have not been able to find a USB device that 'looks'
	//	// like an Arduino.
	//	return ""
}