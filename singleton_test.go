package pattern

import (
	"fmt"
	"testing"
)

var Server struct {
	name string
}

func (this *Server) Listen(addr string) error{
	fmt.Println("server listen at ", addr)
}

var singleton = &Singleton{newFuc:func() interface{} {return new(Server)}}

func GetInstance() *Server {
	return singleton.GetInstance().Inst.(*Server)
}

func Test(t *testing.T) {
	GetInstance().Listen(":80")
}