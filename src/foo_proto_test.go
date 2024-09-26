package proto_test

import (
	"testing"

	foo_proto "github.com/thesayyn/proto_example/src"
)

func TestFoo(t *testing.T) {
	msg := &foo_proto.Foo{
		Value: 1,
	}
	if msg.Value != 1 {
		t.Fail()
	}
}
