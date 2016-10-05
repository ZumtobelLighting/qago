package lr_packages

import (
	"github.com/sclevine/agouti"
	"fmt"
	"os"
)





func Lightrules_login() agouti.Page{
	driver := agouti.Selenium()
	if err := driver.Start(); err != nil {
		//t.Fatal("Failed to start Selenium:", err)
		fmt.Println("Failed to start Selenium:")
		log.Fatalf("%v", err)
	}


	browser, err := driver.NewPage(agouti.Browser("firefox"))
	if err != nil {
		fmt.Println("Failed to open browser:")
	}

	browser.Navigate("http://10.1.30.126/login")
	email_element := browser.FirstByClass("email")
	if email_element == nil {
		log.Error("Failed to find email_element")
	}

	email_element.SendKeys("lightrules@digitallumens.com")
	//
	password_element := browser.FirstByClass("password")
	if password_element == nil {
		log.Error("Failed to find password_element")
	}
	password_element.SendKeys(os.Getenv("LRPASSWORD"))

//
	login_element := browser.FirstByClass("log-in")
	if password_element == nil {
		log.Error("Failed to find login_element")
	}
	login_element.Click()
	return(*browser)






	//if err := driver.Stop(); err != nil {
	//	//t.Fatal("Failed to close pages and stop WebDriver:", err)
	//	fmt.Println("Failed to close pages and stop WebDriver:")
	//}
}
