package main

import (
	"encoding/base64"
	"fmt"
)

type CompressionDecorator struct {
	dataSourceDecorator DataSourceDecorator
}

// constructor
func (c *CompressionDecorator) Decorator(ds DataSource) {
	c.dataSourceDecorator.Decorator(ds)
}

func (c *CompressionDecorator) WriteData(dt Data) {
	// 1. Compress passed data.
	dataEnc64 := base64.StdEncoding.EncodeToString([]byte(dt.ToString()))
	fmt.Printf("DEBUG: CompressionDecorator::WriteData compressData: %s\n", dataEnc64)

	// 2. Pass compressed data to the wrappee's writeData method.
	dataEncrypted := Data{
		DetailData: dataEnc64,
	}
	c.dataSourceDecorator.wrappee.WriteData(dataEncrypted)
}

func (c *CompressionDecorator) ReadData() Data {
	// 1. Get data from the wrappee's readData method.
	dataEnc := c.dataSourceDecorator.wrappee.ReadData()

	// 2. Try to decrypt it if it's encrypted.
	dataDec, _ := base64.StdEncoding.DecodeString(dataEnc.ToString())

	fmt.Printf("DEBUG: EncryptionDecorator::ReadData uncompressData: %s\n", string(dataDec))

	// 3. Return the result.
	return Data{
		DetailData: string(dataDec),
	}

}
