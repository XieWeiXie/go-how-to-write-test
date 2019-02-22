package parttwo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd_Two(t *testing.T) {
	Convey("test add", t, func() {
		Convey("0 + 0", func() {
			So(Add(0, 0), ShouldEqual, 0)
		})
		Convey("-1 + 0", func() {
			So(Add(-1, 0), ShouldEqual, -1)
		})
	})
}

func TestFloatToString_Two(t *testing.T) {
	Convey("test float to string", t, func() {
		Convey("1.0/3.0", func() {
			result := FloatToString(1.0, 3.0)
			So(result, ShouldContainSubstring, "%")
			So(len(result), ShouldEqual, 6)
			So(result, ShouldEqual, "33.33%")
		})
	})

}
