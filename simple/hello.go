package simple

type SayHello interface {
	Hello(name string) string
}

type SayHelloImpl struct {
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

type HelloService struct {
	SayHello SayHello
}

func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{SayHello: sayHello}
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}
