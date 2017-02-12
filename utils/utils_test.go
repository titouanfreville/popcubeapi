package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtilsPackage(t *testing.T) {
	Convey("Testing Array intersections", t, func() {
		Convey("Given an array", func() {
			a := []string{
				"abc",
				"def",
				"ghi",
			}

			empty := []string{}

			Convey("Trying intersection with an empty array or an array without common parts, it should be empty", func() {

				b := []string{
					"jkl",
					"mnp",
				}

				c := []string{
					"jkl",
				}

				So(StringArrayIntersection(a, empty), ShouldBeEmpty)
				So(StringArrayIntersection(a, b), ShouldBeEmpty)
				So(StringArrayIntersection(a, c), ShouldBeEmpty)

			})

			Convey("Trying intersection with common point should return a table containing the common elements", func() {
				b := []string{
					"jkl",
					"abc",
				}

				c := []string{
					"def",
					"mno",
				}

				d := []string{
					"abc",
					"ghi",
					"ameno",
				}

				So(StringArrayIntersection(a, a), ShouldResemble, a)
				So(StringArrayIntersection(a, b), ShouldContain, "abc")
				So(StringArrayIntersection(a, c), ShouldContain, "def")
				So(StringArrayIntersection(a, d), ShouldContain, "abc")
				So(StringArrayIntersection(a, d), ShouldContain, "ghi")
			})
		})
	})
}
