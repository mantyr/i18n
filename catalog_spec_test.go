package i18n_test

import (
	"testing"

	"github.com/mantyr/assertions"
	"github.com/mantyr/i18n"
	"github.com/mantyr/i18n/methods"
	"github.com/mantyr/i18n/sources"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Spec loading", t, func() {
		c, err := i18n.NewCatalog()
		So(err, ShouldBeNil)
		So(c, ShouldNotBeNil)

		testItems := []struct {
			lang i18n.Lang
			key  i18n.Key
			data any
		}{
			{
				lang: "en",
				key:  "messages.greeting",
				data: struct{ Name string }{Name: "World"},
			},
			{
				lang: "ru",
				key:  "messages.greeting",
				data: struct{ Name string }{Name: "World"},
			},
			{
				lang: "en",
				key:  "messages.farewell",
				data: struct{ Name string }{Name: "World"},
			},
			{
				lang: "ru",
				key:  "messages.farewell",
				data: struct{ Name string }{Name: "World"},
			},
		}

		Convey("Last method from YAML", func() {
			spec, err := sources.SpecYAMLFile("./testdata/messages.last.spec.yaml")
			So(err, ShouldBeNil)
			So(spec, ShouldNotBeNil)

			err = c.LoadBySpec(spec)
			So(err, ShouldBeNil)

			d := []string{}
			for _, test := range testItems {
				out, err := c.Execute(test.lang, test.key, test.data)
				So(err, ShouldBeNil)
				So(out, ShouldNotEqual, "")
				d = append(d, out)
			}
			So(d, assertions.ShouldEqualYamlCassette, "./testdata/messages.last.spec.cassette.yaml")
		})

		Convey("First method from YAML", func() {
			spec, err := sources.SpecYAMLFile("./testdata/messages.first.spec.yaml")
			So(err, ShouldBeNil)
			So(spec, ShouldNotBeNil)

			err = c.LoadBySpec(spec)
			So(err, ShouldBeNil)

			d := []string{}
			for _, test := range testItems {
				out, err := c.Execute(test.lang, test.key, test.data)
				So(err, ShouldBeNil)
				So(out, ShouldNotEqual, "")
				d = append(d, out)
			}
			So(d, assertions.ShouldEqualYamlCassette, "./testdata/messages.first.spec.cassette.yaml")
		})
		Convey("unknown method is rejected", func() {
			err := c.LoadBySpec(&sources.Spec{
				Method: "Wrong",
				Data:   map[string]any{"en": map[string]any{"k": "v"}},
			})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "unexpected method (Wrong) for extractor")
		})

		Convey("nil items is rejected", func() {
			err := c.LoadBySpec(&sources.Spec{Method: methods.First})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "empty items")
		})

		Convey("nil spec is rejected", func() {
			err := c.LoadBySpec(nil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "empty spec")
		})
	})
}
