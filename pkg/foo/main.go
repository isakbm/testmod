package foo

import (
	"fmt"

	"google.golang.org/grpc"
)

type Foo struct {
	X int
}

func (f *Foo) Print() {
	fmt.Println(f.X)

	cc, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())

}
