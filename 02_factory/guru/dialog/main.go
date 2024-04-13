package main

import (
	"factory_method/dialog"
	"fmt"
)

func main() {
	fmt.Println("Factory method")
	dialogInstance := dialog.Dialog{}
	dialogInstance.CreateButton("macos")
	dialogInstance.Render()
	dialogInstance.OnClick()
}
