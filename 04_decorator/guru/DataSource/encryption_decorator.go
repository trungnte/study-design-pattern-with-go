package main

import (
	"encoding/base64"
	"fmt"
)

// Concrete decorators must call methods on the wrapped object,
// but may add something of their own to the result. Decorators
// can execute the added behavior either before or after the
// call to a wrapped object.
type EncryptionDecorator struct {
	// base          DataSource
	dataSourceDecorator DataSourceDecorator
}

// constructor
func (e *EncryptionDecorator) Decorator(ds DataSource) {
	e.dataSourceDecorator.Decorator(ds)
}

func (e *EncryptionDecorator) WriteData(dt Data) {
	// 1. Encrypt passed data.
	dataEnc64 := base64.StdEncoding.EncodeToString([]byte(dt.ToString()))
	fmt.Printf("DEBUG: EncryptionDecorator::WriteData dataEnc64: %s\n", dataEnc64)

	// 2. Pass encrypted data to the wrappee's writeData method.
	dataEncrypted := Data{
		DetailData: dataEnc64,
	}
	e.dataSourceDecorator.wrappee.WriteData(dataEncrypted)
}

func (e *EncryptionDecorator) ReadData() Data {
	// 1. Get data from the wrappee's readData method.
	dataEnc := e.dataSourceDecorator.wrappee.ReadData()

	// 2. Try to decrypt it if it's encrypted.
	dataDec, _ := base64.StdEncoding.DecodeString(dataEnc.ToString())

	fmt.Printf("DEBUG: EncryptionDecorator::ReadData decodedData: %s\n", string(dataDec))

	// 3. Return the result.
	return Data{
		DetailData: string(dataDec),
	}

}
