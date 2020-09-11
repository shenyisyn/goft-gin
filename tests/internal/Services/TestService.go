package Services

type TestService struct {
	TestName string
	Naming   *NameService `inject:"-"`
}

func NewTestService(testName string) *TestService {
	return &TestService{TestName: testName}
}

func (this *TestService) Name() string {
	return "test"
}
