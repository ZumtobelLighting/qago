package main

import (



	"github.com/digitallumens/qago/agouti-selenium/lr_packages"

	"fmt"
)

func main() {

	browser := lr_packages.Lightrules_login()
	lr_packages.Lightrules_general_settings(browser)
	fmt.Println("\n*************************\n")
	fmt.Println("Completed general settings")
	number_of_tests := lr_packages.Get_num_test()
	number_of_fails := lr_packages.Get_num_fail()
	fmt.Printf("Number of tests run %v \n", number_of_tests)
	fmt.Printf("Number of failed tests %v \n", number_of_fails)
	fmt.Println("\n*************************\n")

	//if err := driver.Stop(); err != nil {
	//	//t.Fatal("Failed to close pages and stop WebDriver:", err)
	//	fmt.Println("Failed to close pages and stop WebDriver:")
	//}
}

