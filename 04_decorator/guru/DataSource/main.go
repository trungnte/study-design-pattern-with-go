package main

import "fmt"

type Application struct {
}

func (a Application) DumbUsageExample() DataSource {
	salaryRecords := Data{
		DetailData: "this is salary 2024",
	}

	fileDataSource := FileDataSource{
		FileName: "somefile.dat",
	}

	// fileDataSource.WriteData(salaryRecords)

	sourceCompress := CompressionDecorator{}
	sourceCompress.Decorator(&fileDataSource)

	sourceEnc := EncryptionDecorator{}
	sourceEnc.Decorator(&sourceCompress)

	sourceEnc.WriteData(salaryRecords)

	return &sourceEnc
}

type SalaryManager struct {
	source DataSource
}

func (s *SalaryManager) load() Data {
	return s.source.ReadData()
}

func (s *SalaryManager) save(dt Data) {
	s.source.WriteData(dt)
}

func main() {
	// test file_datasource
	// dataTest := Data{
	// 	DetailData: "Hello, this is message for testing",
	// }

	// ftest := FileDataSource{
	// 	FileName: "file_data.dat",
	// }

	// dataSrcRead := ftest.ReadData()
	// fmt.Printf("Data from file: %s\n", dataSrcRead.ToString())

	app := Application{}
	dataSrc := app.DumbUsageExample()

	logger := &SalaryManager{
		source: dataSrc,
	}

	loadData := logger.load()

	fmt.Println(loadData.ToString())

}
