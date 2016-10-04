package lr_packages

import (
	//"github.com/Etienne42/testrail"
	//"os"
	"fmt"
	"github.com/Etienne42/testrail"
	//"log"
	//"net/http"
	//"net/url"
)

func Update_testrail() {

	username := "dlqaautomation@gmail.com"
	password := "Test!ng16"
	//projectID := 2
	//suiteID := 1032

	client := testrail.NewClient("https://digitallumens.testrail.com/", username, password)

	//
	myresult := testrail.SendableResult{StatusID: 1, Comment: "testing with Golang"}
	fmt.Printf("%d\n", myresult.StatusID)
	client.AddResultForCase(2838, 46759, myresult)
}
