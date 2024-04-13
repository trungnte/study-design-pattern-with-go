package dialog

import "fmt"

type WebDialog struct {
}

func (WebDialog) createButton() {
	fmt.Println("WebDialog:: createButton")
}

func (WebDialog) onClick() {
	fmt.Println("WebDialog:: onClick")
}

func (WebDialog) render() {
	fmt.Println("WebDialog:: render")
}
