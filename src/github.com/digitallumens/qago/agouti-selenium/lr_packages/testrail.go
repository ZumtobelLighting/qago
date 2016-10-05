package lr_packages

import (
	//"github.com/Etienne42/testrail"
	//"os"
	"fmt"
	"github.com/Etienne42/testrail"
	//"log"
	//"net/http"
	//"net/url"
	//"runtime"


)

//update_testrail(result, LR_case_id, LR_comment, run_id)
func Update_testrail( result bool, LR_case_id int, LR_comment string, ) {
	username := "dlqaautomation@gmail.com"
	password := "Test!ng16"
	//projectID := 2
	//suiteID := 1032

	global_num_test = global_num_test + 1
	client := testrail.NewClient("https://digitallumens.testrail.com/", username, password)
	//  7 = skipped, 1 = success, 5 = failure
	numericResult := 7
	switch result {
	case true:
		numericResult = 1

	case false:
		numericResult = 5
		global_num_fail = global_num_fail + 1
	}
	//fmt.Printf("numericResult in testrail  = %v \n", numericResult)

	//
	//myresult := testrail.SendableResult{StatusID: 1, Comment: "testing with Golang"}
	testrail_result := testrail.SendableResult{StatusID: 1, Comment: LR_comment}
	//fmt.Printf("%d\n", testrail_result.StatusID)
	//client.AddResultForCase(2838, 46759, testrail_result)

	testrailUpdateResult, err := client.AddResultForCase(global_LR_run_id, LR_case_id, testrail_result)
	if err != nil {
		fmt.Printf("testrail update failed\n")
		fmt.Printf("%v", err)
	}
	fmt.Printf("testrail update succeeded %v", testrailUpdateResult)

}