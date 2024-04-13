package dialog

import "fmt"

type Dialog struct {
	btn Button
}

func (d *Dialog) Render() {
	fmt.Println("Dialog:: Render call d.btn.render()")
	d.btn.render()
}

func (d *Dialog) OnClick() {
	fmt.Println("Dialog:: OnClick call d.btn.onClick()")
	d.btn.onClick()
}

func (d *Dialog) CreateButton(style string) {
	if style == "window" {
		d.btn = WindowsDialog{}
		fmt.Println("Dialog:: CreateButton create WindowsDialog")
	} else {
		d.btn = WebDialog{}
		fmt.Println("Dialog:: CreateButton create WebDialog")
	}
}
