package main

type Data struct {
	DetailData string
}

func (d *Data) ToString() string {
	return d.DetailData
}
