package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestAppErrors(t *testing.T) {

	Convey("Testing message error formating", t, func() {
		Convey("From an error formated error generating a json formated from the error and but it back as error formated error should give the same object", func() {
			err := NewLocAppError("TestAppError", "message", nil, "")
			json := err.ToJSON()
			rerr := AppErrorFromJSON(strings.NewReader(json))
			err.Error()
			So(err.Message, ShouldEqual, rerr.Message)
		})

		Convey("Generating json error error message from html information should work", func() {
			rerr := AppErrorFromJSON(strings.NewReader("<html><body>This is a broken test</body></html>"))
			So("body: <html><body>This is a broken test</body></html>", ShouldEqual, rerr.DetailedError)
		})
	})
}
