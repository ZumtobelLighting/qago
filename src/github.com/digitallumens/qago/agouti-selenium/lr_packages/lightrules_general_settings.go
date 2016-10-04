package lr_packages

import (
	"github.com/sclevine/agouti"
	"fmt"
	//"log"
	//"os"

	//"time"
	"time"
	"strings"
)


func Lightrules_general_settings(browser agouti.Page){
      	test__1_minimum_active_6387(browser)





        //navigate("Admin", browser)
        //time.sleep(2)


}



func test__1_minimum_active_6387(browser agouti.Page){
	//LR_case_id := '6387'
	result := 0
	fmt.Println("in test__1_minimum_active_6387")
	browser.FindByLink("Admin").Click()
	time.Sleep(2 * time.Second)
	browser.FindByLink("General").Click()
	time.Sleep(2 * time.Second)
	title_element, err := browser.Title()
	if err != nil{
		fmt.Println("cannot find General Settings title")
	}
	if strings.Contains(title_element, "General Settings"){
		fmt.Println("found General Settings")

	} else{
		fmt.Println("cannot find General Settings wrong page")
	}
	buttons := browser.AllByClass("button")
	buttons.At(5).Click()
	active_minimum_element := browser.FindByID("site_config_min_active_level")
	min_active_level, err := active_minimum_element.Attribute("value")
	if err != nil {
		fmt.Printf("error")
	}
        fmt.Printf("%v", min_active_level)
	if min_active_level == "20" {
		 result = 1
		fmt.Printf("%v", result)
	} else{
		result = 0
		fmt.Printf("%v", result)
		}
	fmt.Printf("result = %v", result)
	Update_testrail()




}

//if err := driver.Stop(); err != nil {
	//	//t.Fatal("Failed to close pages and stop WebDriver:", err)
	//	fmt.Println("Failed to close pages and stop WebDriver:")
	//}


