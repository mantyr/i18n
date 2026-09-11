package i18n_test

import (
	"fmt"

	"github.com/mantyr/i18n"
	"github.com/mantyr/i18n/extractors"
	"github.com/mantyr/i18n/sources"
)

// Example_yaml shows loading messages from a YAML file. The First extractor is
// used because the language code is the first path segment in the file
// (en: / de: at the top level). Keys become the remaining dotted path, e.g.
// "page.title".
func Example_yaml() {
	c, _ := i18n.NewCatalog()
	err := c.LoadBySource(
		extractors.First,
		sources.YAMLFile("./testdata/example.yaml"),
	)
	if err != nil {
		fmt.Println("load error:", err)
		return
	}

	title, _ := c.Execute("en", "page.title", struct{ Name string }{Name: "World"})
	subtitle, _ := c.Execute("de", "page.subtitle", nil)

	fmt.Println(title)
	fmt.Println(subtitle)
	// Output:
	// Welcome, World!
	// Sie haben neue Nachrichten
}
