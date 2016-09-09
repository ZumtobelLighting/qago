package main

import "fmt"
import "github.com/digitallumens/dllibgo"
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
	//"github.com/tarm/serial"

	//"reflect"


)

func main() {
	//stick_port := "/dev/" + find_usb_stick()
	//fmt.Printf(stick_port)
	//c := &serial.Config{Name: stick_port, Baud: 115200, ReadTimeout: 20}
	//fmt.Printf("%s\n", c.Name)
	g, err := dllibgo.NewEZSPGateway("/dev/tty.SLAB_USBtoUART")
	if err != nil {
		fmt.Printf("got an error")
		log.Fatal(err)
	}

	fmt.Printf("%t",g.IsJoined())
	//fmt.Printf("before join network")
	//network_err := g.JoinNetwork("B16")
	//fmt.Printf("after join network")
	//if network_err != nil {
	//	log.Fatal(network_err)
	//}
	//boolJoined := g.IsJoined()
	//fmt.Printf("%t", boolJoined)
	//addr16, result := g.NodeAddressFromSerialNumber(uint32("0401B5CA"))
	//fmt.Printf(string(result))
	//fmt.Printf(string(addr16))




}

func find_usb_stick() string{
	files, err := ioutil.ReadDir("/dev")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//fmt.Println(file.Name())
		if strings.Contains(file.Name(), "cu.SLAB_USBtoUART") {
			strName := file.Name()
			fmt.Printf( strName + "\n")
			//fmt.Println(reflect.TypeOf(file.Name))

			return(strName)




		}
	}
	return("failed")
	}

//   eof
