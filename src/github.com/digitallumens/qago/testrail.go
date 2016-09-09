package main

import (
	//"github.com/Etienne42/testrail"
	//"os"
	"fmt"
	"log"
	"github.com/Etienne42/testrail"
	//"net/http"
	//"net/url"
)


  func main(){

    //username := os.Getenv("TESTRAIL_USERNAME")
    //password := os.Getenv("TESTRAIL_TOKEN")
    username := "dlqaautomation@gmail.com"
    password := "Test!ng16"
    projectID := 2
    suiteID := 1032


    client := testrail.NewClient("https://digitallumens.testrail.com/", username, password)


    cases, err := client.GetCases(projectID, suiteID)
    if err != nil {
	fmt.Printf("got an error\n")
	log.Fatal(err)
	}

    for _, c := range cases{
      fmt.Println(c.ID)
    }
    //
    myresult := testrail.SendableResult{StatusID:1,Comment:"testing with Golang"}
    fmt.Printf("%d\n", myresult.StatusID)
    client.AddResultForCase(2838,46759, myresult)
  }