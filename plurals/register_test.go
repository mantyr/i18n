// Tests for Plural.Register (wiring the plural function into templates).
package plurals_test

import (
	"bytes"
	"testing"
	texttemplate "text/template"

	"github.com/mantyr/i18n"
	"github.com/mantyr/i18n/plurals"
	. "github.com/smartystreets/goconvey/convey"
)

// nolint:errcheck
func TestPluralRegister(t *testing.T) {
	Convey("Plural.Register", t, func() {
		c, _ := i18n.NewCatalog()
		c.Set("en", "items.one", "{{.}} item")
		c.Set("en", "items.other", "{{.}} items")
		pl := plurals.NewCardinal().SetCatalog(c)

		Convey("default name plural", func() {
			tpl := texttemplate.New("p")
			pl.Register(tpl)
			tpl.Parse(`{{plural .Lang "items" .N}}`)
			var b bytes.Buffer
			err := tpl.Execute(&b, map[string]any{"Lang": i18n.Lang("en"), "N": 5})
			So(err, ShouldBeNil)
			So(b.String(), ShouldEqual, "5 items")
		})

		Convey("custom name and both funcs together", func() {
			c.Set("en", "hdr", "List:")
			tpl := texttemplate.New("p")
			c.Register(tpl)        // i18n
			pl.Register("pl", tpl) // pl
			tpl.Parse(`{{i18n .Lang "hdr" .}} {{pl .Lang "items" .N}}`)
			var b bytes.Buffer
			err := tpl.Execute(&b, map[string]any{"Lang": i18n.Lang("en"), "N": 1})
			So(err, ShouldBeNil)
			So(b.String(), ShouldEqual, "List: 1 item")
		})
	})
}
