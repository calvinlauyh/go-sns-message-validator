package snserrors

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("It should return an SNS Error with type and message", t, func() {
		expected := SNSError{"TestError", "This is a test error"}
		actual := New("TestError", "This is a test error")

		So(*actual, ShouldResemble, expected)
	})

	Convey("It should be treated as an error", t, func() {
		fn := func() {
			var errHolder error
			errHolder = New("TestError", "This is a test error")

			// Bypass the Golang variable not used error
			_ = errHolder
		}

		So(fn, ShouldNotPanic)
	})
}

func TestTypeMethod(t *testing.T) {
	Convey("It should return the type in string", t, func() {
		err := New("TestError", "This is a test error")
		expected := "TestError"
		actual := err.Type()

		So(actual, ShouldEqual, expected)
	})
}

func TestIsMethod(t *testing.T) {
	Convey("It should return true when the error is the given type", t, func() {
		err := New("TestError", "This is a test error")

		So(err.Is("TestError"), ShouldBeTrue)
	})

	Convey("It should return true when the error is not the given type", t, func() {
		err := New("TestError", "This is a test error")

		So(err.Is("CustomError"), ShouldBeFalse)
	})
}

func TestError(t *testing.T) {
	Convey("It should return the message part of the error", t, func() {
		err := New("TestError", "This is a test error")
		expected := "This is a test error"
		actual := err.Error()

		So(actual, ShouldEqual, expected)
	})
}

func ExampleNew() {
	err := New("TestError", "This is a test error")
	if err != nil {
		fmt.Print(err)
	}
	// Output: This is a test error
}
