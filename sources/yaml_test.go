package sources_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/mantyr/assertions"
	"github.com/mantyr/i18n"
	"github.com/mantyr/i18n/extractors"
	"github.com/mantyr/i18n/sources"
	. "github.com/smartystreets/goconvey/convey"
)

type testItem struct {
	Lang        i18n.Lang
	Key         i18n.Key
	Result      string
	ResultError string
}

func TestYAML(t *testing.T) {
	Convey("Source YAML", t, func() {
		Convey("Last extractor", func() {
			c, err := i18n.NewCatalog()
			So(err, ShouldBeNil)
			So(c, ShouldNotBeNil)
			err = c.LoadBySource(extractors.Last, sources.YAMLFile("./testdata/last.messages.yaml"))
			So(err, ShouldBeNil)

			Convey("Languages", func() {
				languages := c.Languages()
				slices.Sort(languages)
				So(languages, assertions.ShouldEqualYamlCassette, "./testdata/last.messages.languages.cassette.yaml")
				Convey("Dataset", func() {
					dataset := []testItem{}
					for _, lang := range languages {
						dictionary, err := c.Dictionary(lang)
						So(err, ShouldBeNil)
						So(dictionary, ShouldNotBeNil)
						So(dictionary.Lang(), ShouldEqual, lang)
						keys := dictionary.Keys()
						slices.Sort(keys)
						for _, key := range keys {
							result, err := dictionary.Execute(key, struct{ Name string }{Name: "test"})
							dataset = append(dataset, testItem{
								Lang:        lang,
								Key:         key,
								Result:      result,
								ResultError: fmt.Sprintf("%v", err),
							})
						}
					}
					So(dataset, assertions.ShouldEqualYamlCassette, "./testdata/last.messages.dataset.cassette.yaml")
				})
			})
		})
		Convey("First extractor", func() {
			c, err := i18n.NewCatalog()
			So(err, ShouldBeNil)
			So(c, ShouldNotBeNil)
			err = c.LoadBySource(extractors.First, sources.YAMLFile("./testdata/first.messages.yaml"))
			So(err, ShouldBeNil)

			Convey("Languages", func() {
				languages := c.Languages()
				slices.Sort(languages)
				So(languages, assertions.ShouldEqualYamlCassette, "./testdata/first.messages.languages.cassette.yaml")
				Convey("Dataset", func() {
					dataset := []testItem{}
					for _, lang := range languages {
						dictionary, err := c.Dictionary(lang)
						So(err, ShouldBeNil)
						So(dictionary, ShouldNotBeNil)
						So(dictionary.Lang(), ShouldEqual, lang)
						keys := dictionary.Keys()
						slices.Sort(keys)
						for _, key := range keys {
							result, err := dictionary.Execute(key, struct{ Name string }{Name: "test"})
							dataset = append(dataset, testItem{
								Lang:        lang,
								Key:         key,
								Result:      result,
								ResultError: fmt.Sprintf("%v", err),
							})
						}
					}
					So(dataset, assertions.ShouldEqualYamlCassette, "./testdata/first.messages.dataset.cassette.yaml")
				})
			})
		})
	})
}
