package lr_packages

import (
	"github.com/sclevine/agouti"
	"fmt"
	"github.com/op/go-logging"
	"time"
	"strings"
	"os"
)

var log = logging.MustGetLogger("main")

func Lightrules_general_settings(browser agouti.Page){
      	//test__1_minimum_active_6387(browser)
	test_2_DH_reporting_ok_6388(browser)





        //navigate("Admin", browser)
        //time.sleep(2)


}



func test__1_minimum_active_6387(browser agouti.Page){
	LR_case_id := 6387
	LR_comment := "agouti testing 111\n"
	result := false
	log.Info("\nin test__1_minimum_active_6387\n")
	pwd, err := os.Getwd()
	    if err != nil {
		log.Error("error in os.Getwd")
	    }
	log.Infof("%v", pwd)
	browser.FindByLink("Admin").Click()
	time.Sleep(2 * time.Second)
	browser.FindByLink("General").Click()
	time.Sleep(2 * time.Second)
	title_element, err := browser.Title()
	if err != nil{
		fmt.Println("cannot find General Settings title")
	}
	if strings.Contains(title_element, "General Settings"){
		fmt.Println("found General Settings\n")

	} else{
		fmt.Println("cannot find General Settings wrong page\n")
	}
	buttons := browser.AllByClass("button")
	buttons.At(5).Click()
	active_minimum_element := browser.FindByID("site_config_min_active_level")
	min_active_level, err := active_minimum_element.Attribute("value")
	if err != nil {
		fmt.Printf("error")
	}
        //fmt.Printf("%v", min_active_level)
	if min_active_level == "20" {
		 result = true

	}
	log.Infof("result = %v \n", result)

	Update_testrail(result, LR_case_id, LR_comment )



}


func test_2_DH_reporting_ok_6388(browser agouti.Page){
	LR_case_id := 6387
	LR_comment := "agouti testing 111\n"
	result := false
	log.Info("\nin test_2_DH_reporting_ok_6388\n")

	Upload_map(browser, "dh_none.map")
	browser.FindByLink("Admin").Click()
	time.Sleep(2 * time.Second)
	browser.FindByLink("General").Click()
	edit_general_element := browser.FindByLink("Edit General Settings")
	edit_general_element.Click()
	dh_element := browser.FirstByXPath("//*[@id='site_config_daylight_harvesting_enabled_input']/p")
	log.Info(dh_element.Text())

	// update testrail
	log.Infof("result = %v \n", result)
	Update_testrail(result, LR_case_id, LR_comment )

}

func Upload_map(browser agouti.Page, mapname string){
	log.Info("in Upload_map")
	browser.FindByLink("Admin").Click()
	time.Sleep(2 * time.Second)
	browser.FindByLink("Config").Click()
	time.Sleep(2 * time.Second)
	browser.FindByLink("Manage Map File").Click()
	time.Sleep(2 * time.Second)
	pwd, err := os.Getwd()
	    if err != nil {
		log.Error("error in os.Getwd")
	    }
	log.Infof("%v", pwd)
	filename := pwd + "/" + mapname
	log.Infof("filename = %v", filename)
	mapname_element := browser.FindByID("map_filename")
	mapname_element.SendKeys(filename)
	commit_element := browser.FindByName("commit")
	commit_element.Click()
	time.Sleep(3 * time.Second)
	commit_element_1 := browser.FindByName("commit")
	commit_element_1.Click()
	time.Sleep(30 * time.Second)

}

//if err := driver.Stop(); err != nil {
	//	//t.Fatal("Failed to close pages and stop WebDriver:", err)
	//	fmt.Println("Failed to close pages and stop WebDriver:")
	//}


