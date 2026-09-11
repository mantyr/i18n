// Tests for the plural subpackage (uses testdata/plurals.ttl.yaml).
package plurals_test

import (
	"testing"
	"time"

	"github.com/mantyr/i18n"
	"github.com/mantyr/i18n/extractors"
	"github.com/mantyr/i18n/humanize/ttl"
	"github.com/mantyr/i18n/plurals"
	"github.com/mantyr/i18n/sources"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPluralTTL(t *testing.T) {
	Convey("Plural over TTL", t, func() {
		c, err := i18n.NewCatalog()
		So(err, ShouldBeNil)
		err = c.LoadBySource(extractors.First, sources.YAMLFile("./testdata/plural.ttl.yaml"))
		So(err, ShouldBeNil)

		p := plurals.NewCardinal().SetCatalog(c)
		So(p.Error(), ShouldBeNil)

		Convey("russian forms", func() {
			out, err := p.Execute("ru", "ttl.days", 1)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, "1 день")

			out, err = p.Execute("ru", "ttl.days", 3)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, "3 дня")

			out, err = p.Execute("ru", "ttl.days", 5)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, "5 дней")
		})

		Convey("english forms", func() {
			out, err := p.Execute("en", "ttl.hours", 1)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, "1 hour")

			out, err = p.Execute("en", "ttl.hours", 2)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, "2 hours")
		})

		Convey("splitting a duration and rendering each part", func() {
			d, err := ttl.Split(50 * time.Hour) // 2 days 2 hours
			So(err, ShouldBeNil)
			So(d.Days, ShouldEqual, 2)
			So(d.Hours, ShouldEqual, 2)

			days, err := p.Execute("ru", "ttl.days", d.Days)
			So(err, ShouldBeNil)
			So(days, ShouldEqual, "2 дня")
		})
	})
}

func TestPluralErrors(t *testing.T) {
	Convey("Plural error handling", t, func() {
		Convey("nil rules yields error", func() {
			p := plurals.New(nil)
			So(p.Error(), ShouldNotBeNil)
			_, err := p.Execute("ru", "k", 1)
			So(err, ShouldNotBeNil)
		})
		Convey("nil catalog yields error", func() {
			p := plurals.NewCardinal().SetCatalog(nil)
			So(p.Error(), ShouldNotBeNil)
		})
		Convey("bad language yields error", func() {
			c, _ := i18n.NewCatalog()
			p := plurals.NewCardinal().SetCatalog(c)
			_, err := p.Execute("not-a-lang!!", "k", 1)
			So(err, ShouldNotBeNil)
		})
	})
}
