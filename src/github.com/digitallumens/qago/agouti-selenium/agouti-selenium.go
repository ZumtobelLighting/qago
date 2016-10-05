package main

import (



	"github.com/digitallumens/qago/agouti-selenium/lr_packages"

)

func main() {

	browser := lr_packages.Lightrules_login()
	lr_packages.Lightrules_general_settings(browser)



	//if err := driver.Stop(); err != nil {
	//	//t.Fatal("Failed to close pages and stop WebDriver:", err)
	//	fmt.Println("Failed to close pages and stop WebDriver:")
	//}
}

