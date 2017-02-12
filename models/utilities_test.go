// Testing base tools for DB models.
package models

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtilities(t *testing.T) {
	Convey("Testing ID generation", t, func() {
		numberOfGeneration := 1000
		assertion := "Checking validity of " + strconv.Itoa(numberOfGeneration) + " random ids"

		Convey(assertion, func() {
			for i := 0; i < numberOfGeneration; i++ {
				id := NewID()
				So(len(id), ShouldBeLessThan, 27)
			}
		})
	})

	Convey("Testing random string generation", t, func() {
		numberOfGeneration := 1000
		assertion := "Checking correct generation of " + strconv.Itoa(numberOfGeneration) + " random string"

		Convey(assertion, func() {
			for i := 0; i < numberOfGeneration; i++ {
				r := NewRandomString(32)
				So(len(r), ShouldEqual, 32)
			}
		})
	})

	// Convey("Testing Map from/to Json conversions", t, func() {

	// 	Convey("Convert a map to json then back to map should provide same map", func() {
	// 		m := make(map[string]string)
	// 		m["id"] = "test_id"
	// 		json := MapToJSON(m)
	// 		correct := MapFromJSON(strings.NewReader(json))
	// 		So(correct["id"], ShouldEqual, "test_id")
	// 	})

	// 	Convey("Using an empty json to generate map should provide empty map", func() {
	// 		invalid := MapFromJSON(strings.NewReader(""))
	// 		So(len(invalid), ShouldEqual, 0)
	// 	})
	// })

	Convey("Testing email validation", t, func() {
		correctMail := "test.test+xala@something.co"
		invalidMail := "@test.test+xala@something.co"

		Convey("Validating a correctly formated email should be accepted", func() {
			So(IsValidEmail(correctMail), ShouldBeTrue)
		})

		Convey("Validating a non correctly formated email should correctly be refused", func() {
			So(IsValidEmail(invalidMail), ShouldBeFalse)
			So(IsValidEmail("Corey+test@hulen.com"), ShouldBeFalse)
		})
	})

	// Convey("Testing Lower case string checker", t, func() {

	// 	Convey("Providing a lower case test to lowercase checker should return true", func() {
	// 		So(u.IsLower("corey+test@hulen.com"), ShouldBeTrue)
	// 	})

	// 	Convey("Providing a non lower case test to lowercase checker should return false", func() {
	// 		So(u.IsLower("Corey+test@hulen.com"), ShouldBeFalse)
	// 	})
	// })

	Convey("Testing Etag creation", t, func() {
		Convey("Providing two parameters to function should return a string composed of version number.par1.par2", func() {
			Etag := Etag("hello", 24)
			result := CurrentVersion + ".hello.24"
			So(Etag, ShouldEqual, result)
		})
	})

	// Convey("Testing string into array function", t, func() {
	// 	Convey("Given an array", func() {
	// 		array := []string{"test", "dragon", "stupid"}
	// 		Convey("Searching for existing string should return true", func() {
	// 			So(StringInArray("test", array), ShouldBeTrue)
	// 			So(StringInArray("dragon", array), ShouldBeTrue)
	// 			So(StringInArray("stupid", array), ShouldBeTrue)
	// 		})

	// 		Convey("Searching non existing strings should return false", func() {
	// 			So(StringInArray("test", []string{}), ShouldBeFalse)
	// 			So(StringInArray("libellule", array), ShouldBeFalse)
	// 			So(StringInArray("man", array), ShouldBeFalse)
	// 		})
	// 	})
	// })

	Convey("Testing Hastags parsing ", t, func() {

		var hashtags = map[string]string{
			"#test":           "#test",
			"test":            "",
			"#test123":        "#test123",
			"#123test123":     "",
			"#test-test":      "#test-test",
			"#test?":          "#test",
			"hi #there":       "#there",
			"#bug #idea":      "#bug #idea",
			"#bug or #gif!":   "#bug #gif",
			"#hüllo":          "#hüllo",
			"#?test":          "",
			"#-test":          "",
			"#yo_yo":          "#yo_yo",
			"(#brakets)":      "#brakets",
			")#stekarb(":      "#stekarb",
			"<#less_than<":    "#less_than",
			">#greater_than>": "#greater_than",
			"-#minus-":        "#minus",
			"_#under_":        "#under",
			"+#plus+":         "#plus",
			"=#equals=":       "#equals",
			"%#pct%":          "#pct",
			"&#and&":          "#and",
			"^#hat^":          "#hat",
			"##brown#":        "#brown",
			"*#star*":         "#star",
			"|#pipe|":         "#pipe",
			":#colon:":        "#colon",
			";#semi;":         "#semi",
			"#Mötley;":        "#Mötley",
			".#period.":       "#period",
			"¿#upside¿":       "#upside",
			"\"#quote\"":      "#quote",
			"/#slash/":        "#slash",
			"\\#backslash\\":  "#backslash",
			"#a":              "",
			"#1":              "",
			"foo#bar":         "",
		}

		Convey("A text containing or not containing # should be correctly parse", func() {
			for input, output := range hashtags {
				o, _ := ParseHashtags(input)
				So(o, ShouldEqual, output)
			}
		})
	})
}
