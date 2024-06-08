// The base decorator class follows the same interface as the
// other components. The primary purpose of this class is to
// define the wrapping interface for all concrete decorators.
// The default implementation of the wrapping code might include
// a field for storing a wrapped component and the means to
// initialize it.

package main

type DataSourceDecorator struct {
	wrappee DataSource
}

// constructor
func (d *DataSourceDecorator) Decorator(ds DataSource) {
	d.wrappee = ds
}

func (d *DataSourceDecorator) WriteData(dt Data) {
	// fmt.Printf("DataSourceDecorator::WriteData -> %s\n", dt.Read())
	d.wrappee.WriteData(dt)
}

func (d *DataSourceDecorator) ReadData() Data {
	return d.wrappee.ReadData()
}
