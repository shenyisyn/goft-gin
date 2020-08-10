package Services

import "fmt"

type NameService struct {
	MyName string
}

func NewNameService(myName string) *NameService {
	return &NameService{MyName: myName}
}
func (this *NameService) ShowName() {
	fmt.Println(this.MyName)
}
