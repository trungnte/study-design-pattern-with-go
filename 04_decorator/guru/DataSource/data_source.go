package main

// The component interface defines operations that can be
// altered by decorators.
type DataSource interface {
	WriteData(data Data)
	ReadData() Data
}
