package main

import (

	"github.com/op/go-logging"

	"github.com/digitallumens/qago/agouti-selenium/lr_packages"

	//"fmt"
)
var log = logging.MustGetLogger("main")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

//  Chris
//const format = "%{color}[%{time:15:01:03.000} %{level}]%{color:reset}[%{module}-%{shortfile}] %{message}"


func main() {
	log.Debugf("debug")
	log.Error("error")
	//log.Fatal("fatal")
	log.Info("info")
	log.Warning("warning")
	browser := lr_packages.Lightrules_login()
	lr_packages.Lightrules_general_settings(browser)
	log.Info("\n*************************\n")
	log.Info("Completed general settings")
	number_of_tests := lr_packages.Get_num_test()
	number_of_fails := lr_packages.Get_num_fail()
	log.Infof("Number of tests run %v \n", number_of_tests)
	log.Infof("Number of failed tests %v \n", number_of_fails)
	log.Info("\n*************************\n")

	//if err := driver.Stop(); err != nil {
	//	//t.Fatal("Failed to close pages and stop WebDriver:", err)
	//	fmt.Println("Failed to close pages and stop WebDriver:")
	//}
}

