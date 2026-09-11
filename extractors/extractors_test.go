// Tests for extractors.First / extractors.Last.
package extractors_test

import (
	"testing"

	"github.com/mantyr/i18n"
	"github.com/mantyr/i18n/extractors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFirst(t *testing.T) {
	Convey("First extractor", t, func() {
		Convey("language is the first segment", func() {
			lang, key, err := extractors.First([]string{"en", "messages", "greeting"})
			So(err, ShouldBeNil)
			So(lang, ShouldEqual, i18n.Lang("en"))
			So(key, ShouldEqual, i18n.Key("messages.greeting"))
		})
		Convey("path shorter than 2 is rejected", func() {
			_, _, err := extractors.First([]string{"en"})
			So(err, ShouldNotBeNil)
		})
		Convey("two-element path yields single-segment key", func() {
			lang, key, err := extractors.First([]string{"ru", "hello"})
			So(err, ShouldBeNil)
			So(lang, ShouldEqual, i18n.Lang("ru"))
			So(key, ShouldEqual, i18n.Key("hello"))
		})
	})
}

func TestLast(t *testing.T) {
	Convey("Last extractor", t, func() {
		Convey("language is the last segment", func() {
			lang, key, err := extractors.Last([]string{"messages", "greeting", "en"})
			So(err, ShouldBeNil)
			So(lang, ShouldEqual, i18n.Lang("en"))
			So(key, ShouldEqual, i18n.Key("messages.greeting"))
		})
		Convey("path shorter than 2 is rejected", func() {
			_, _, err := extractors.Last([]string{"en"})
			So(err, ShouldNotBeNil)
		})
	})
}
