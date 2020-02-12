package util_test

import (
	"net/http"
	"testing"

	"goscrum/server/util"

	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRedirect(t *testing.T) {
	Convey("Response -> Redirect", t, func() {
		res, err := util.Redirect("test.png")
		So(res.StatusCode, ShouldEqual, http.StatusTemporaryRedirect)
		So(err, ShouldBeNil)
	})

	Convey("Response -> Success", t, func() {
		res, err := util.Success("test.png")
		So(res.StatusCode, ShouldEqual, http.StatusOK)
		So(err, ShouldBeNil)
	})

	Convey("Response -> ClientError", t, func() {
		res, err := util.ClientError(http.StatusBadGateway)
		So(res.StatusCode, ShouldEqual, http.StatusBadGateway)
		So(err, ShouldBeNil)
	})

	Convey("Response -> ServerError", t, func() {
		res, err := util.ServerError(errors.New("error"))
		So(res.StatusCode, ShouldEqual, http.StatusInternalServerError)
		So(err, ShouldBeNil)
	})
}
