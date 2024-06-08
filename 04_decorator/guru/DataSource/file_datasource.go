package main

// Concrete components provide default implementations for the
// operations. There might be several variations of these
// classes in a program.
type FileDataSource struct {
	FileName   string
	BufferData string
}

func (f *FileDataSource) WriteData(dt Data) {
	f.BufferData = dt.DetailData
}

func (f *FileDataSource) ReadData() Data {
	return Data{
		DetailData: f.BufferData,
	}
}
