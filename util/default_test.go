package util_test

import (
	"testing"

	"goscrum/server/util"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDefaultInt(t *testing.T) {
	m := map[string]string{"A": "10", "B": "20"}
	Convey("DefaultInt -> 10", t, func() {
		val := util.DefaultInt(m, "A", 0)
		So(val, ShouldEqual, 10)
	})

	Convey("DefaultInt -> 10", t, func() {
		val := util.DefaultInt(m, "Z", 0)
		So(val, ShouldEqual, 0)
	})
}
