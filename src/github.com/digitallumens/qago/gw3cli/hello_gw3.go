package main

import (
	"fmt"
	//"strconv"
	//"bufio"
	//"strings"
	//"io/ioutil"
)
import (
	//"github.com/tarm/serial"
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
	//"os"
	//"os/exec"
	"github.com/digitallumens/gateway3/gw3config/gw3usb"
	"os"

	"io/ioutil"
	"strings"
	//"github.com/digitallumens/gateway3/gw3pb"
)

var client = gw3usb.Client{}

func main() {
	fmt.Printf("hello_gw3.go\n")
	// discover the gateway and allow command to execute
	if err := mydiscoverUSBModem(&client); err != nil {
		fmt.Printf("cannot find serial port")
	}

	fmt.Fprintln(os.Stderr, "found USB device:", client.Dev)

	//result, err := client.GetConfig()
	//if err != nil {
	//	 log.Fatal(err)
	//	}
	fmt.Printf(client.Dev)
	fmt.Printf("\n")
	//result, err := client.ExecuteCommand("get-config")
	//if err != nil {
	//	 log.Fatal(err)
	//	}
	//fmt.Printf("%v", result)



	result, err := client.GetConfig()
	if err != nil {
		 log.Fatal(err)
		}
	fmt.Printf("%v", result)
	fmt.Printf("\n")
	fmt.Printf("MacAddress = %v", result.GetFactory().GetMacAddress())
	fmt.Printf("\n")
	fmt.Printf("SerialNumber = %v", result.GetFactory().GetSerialNumber())
	fmt.Printf("\n")
	fmt.Printf("EmberPort = %v", result.GetFactory().GetEmberPort())
	fmt.Printf("\n")
	fmt.Printf("MagicKey = %v", result.GetFactory().GetMagicKey())
	fmt.Printf("\n")
	fmt.Printf("Mode = %v", result.GetSetup().GetMode())
	fmt.Printf("\n")
	fmt.Printf("Setup.String = %v", result.GetSetup().String())
	fmt.Printf("\n")
	fmt.Printf("IPv4Address = %v", result.GetEthernet().GetIPv4Address())
	fmt.Printf("\n")
	fmt.Printf("IPv4Netmask = %v", result.GetEthernet().GetIPv4Netmask())
	fmt.Printf("\n")
	fmt.Printf("IPv4Method = %v", result.GetEthernet().GetIPv4Method())
	fmt.Printf("\n")
	fmt.Printf("State = %v", result.GetEthernet().GetState())
	fmt.Printf("\n")















}


func mydiscoverUSBModem(client *gw3usb.Client) error {
	files, _ := ioutil.ReadDir("/dev")
	for _, f := range files {
		if strings.Contains(f.Name(), "tty.usbmodem") {
			client.Dev = "/dev/" + f.Name()
			return nil
		}
	}
	fmt.Printf("error cannot find serial port")
	return(nil)
}




