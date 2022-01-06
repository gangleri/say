package say_test

import (
	"fmt"
	"testing"

	"github.com/gangleri/say"
)

func ExampleHello() {
	greeting := say.Hello("Bod")
	fmt.Println(greeting)
}


func TestHello(t *testing.T) {
	n := "Bob"
	expected := "Hello Bob."
	actual := say.Hello(n)

	if expected != actual {
		t.Logf("Hello: expected [%s] got [%s]", expected, actual)
		t.Fail()
	}
}
