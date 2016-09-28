package main

import gexpect "github.com/ThomasRooney/gexpect"
import (
	"log"
	"os"
	"fmt"
	"bufio"
)

func main() {
	log.Printf("fixing postmark")
	consolereader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the IP of your LRA : ")

	input, err := consolereader.ReadString('\n') // this will prompt the user for input

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }


	command_string := "ssh dlmaint@" + input
	fmt.Println(command_string)
	child, err := gexpect.Spawn(command_string)
	if err != nil {
		fmt.Println("panic")
		panic(err)
	}
	//child.Expect("Name")
	//child.SendLine("anonymous")
	fmt.Println("before password")
	child.Expect("Password")
	password := os.Getenv("PASSWORD")
	child.SendLine(password)

	child.Expect("Digital Lumens LightRules Appliance")
	child.SendLine("cd /var/www/lightrules")
	child.SendLine("cp .env.qa .env.production")
	fmt.Println("before cat")
	child.SendLine("cat .env.production")

	result, err := child.ExpectRegex(".*LightRules 3.0 QA keys.*")
	if result {
		fmt.Printf("%s", result)
		fmt.Println("\n.env.qa has been copied over to .env.production successfully")
		fmt.Println("I am going to reboot now")
		child.SendLine("sudo reboot")
		child.Expect("[sudo] password for dlmaint:")
		child.SendLine(password)
		child.Interact()
		fmt.Println("after interact")
		child.Close()

	}else{
		fmt.Println("something went wrong please check")
	}


}