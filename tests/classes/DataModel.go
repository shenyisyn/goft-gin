package classes

type DataModel struct {
	Id   int
	Name string
}

func NewDataModel(id int, name string) *DataModel {
	return &DataModel{Id: id, Name: name}
}
