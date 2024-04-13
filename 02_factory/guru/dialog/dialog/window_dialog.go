package dialog

import "fmt"

type WindowsDialog struct {
}

func (WindowsDialog) createButton() {
	fmt.Println("WindowsDialog:: createButton")
}

func (WindowsDialog) onClick() {
	fmt.Println("WindowsDialog:: onClick")
}

func (WindowsDialog) render() {
	fmt.Println("WindowsDialog:: render")
}
