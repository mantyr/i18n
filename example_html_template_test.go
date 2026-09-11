package i18n_test

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/mantyr/i18n"
	"github.com/mantyr/i18n/plurals"
)

// Example_htmlTemplate shows the simplest way to wire i18n into a real
// html/template page: let the catalog and the plural engine register their
// own functions into the template with Register, then pass the language
// explicitly in the template. No manual FuncMap is needed.
// nolint:errcheck
func Example_htmlTemplate() {
	msgs, _ := i18n.NewCatalog()
	msgs.Set("en", "title", "Welcome, {{.Name}}!")
	msgs.Set("en", "items.one", "{{.}} item")
	msgs.Set("en", "items.other", "{{.}} items")

	pl := plurals.NewCardinal().SetCatalog(msgs)

	page := template.New("page")
	msgs.Register(page) // adds {{i18n .Lang "key" data}}
	pl.Register(page)   // adds {{plural .Lang "key" n}}

	template.Must(page.Parse(
		`<h1>{{i18n .Lang "title" .}}</h1>` +
			`<p>{{plural .Lang "items" .Count}}</p>`,
	))

	var buf bytes.Buffer
	_ = page.Execute(&buf, map[string]any{
		"Lang":  i18n.Lang("en"),
		"Name":  "World",
		"Count": 3,
	})

	fmt.Println(buf.String())
	// Output: <h1>Welcome, World!</h1><p>3 items</p>
}
