package main

import (
	"fmt"
	"os"
)

// Concrete components provide default implementations for the
// operations. There might be several variations of these
// classes in a program.
type FileDataSource struct {
	FileName   string
	BufferData string
}

func (f *FileDataSource) WriteData(dt Data) {
	f.BufferData = dt.DetailData
	err := os.WriteFile(f.FileName, []byte(f.BufferData), 0644)
	if err != nil {
		fmt.Println("Write file failed err:", err)
	}
}

func (f *FileDataSource) ReadData() Data {
	dat, err := os.ReadFile(f.FileName)
	if err != nil {
		fmt.Println("Read file failed err:", err)
	}
	f.BufferData = string(dat)

	return Data{
		DetailData: f.BufferData,
	}
}
