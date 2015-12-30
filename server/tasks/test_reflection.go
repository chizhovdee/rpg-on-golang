package tasks
import (
	"github.com/chuckpreslar/gofer"
	"reflect"
	"fmt"
)

type Obj struct {
	A int
	B string
}

var TestRef = gofer.Register(gofer.Task{
	Label:       "ref",
	Description: "Test Reflection",
	Action: func(arguments ...string) error {

		a := Obj{99, "HUI"}

		at := reflect.ValueOf(a)

		fmt.Println(at)

		return nil
	},
})