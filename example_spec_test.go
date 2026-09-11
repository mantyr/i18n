package i18n_test

import (
	"fmt"

	"github.com/mantyr/i18n"
	"github.com/mantyr/i18n/sources"
)

// Example_spec shows loading a catalog from a single Spec document that
// carries both the method and the items. This is the format meant to be
// embedded into a Kubernetes CRD later:
//
//	method: Last
//	items:
//	  page:
//	    title:
//	      en: "Welcome, {{.Name}}!"
//	      de: "Hallo, {{.Name}}!"
func Example_spec() {
	spec, err := sources.SpecYAMLFile("./testdata/example.spec.yaml")
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	c, _ := i18n.NewCatalog()
	if err := c.LoadBySpec(spec); err != nil {
		fmt.Println("load error:", err)
		return
	}

	en, _ := c.Execute("en", "page.title", struct{ Name string }{Name: "World"})
	de, _ := c.Execute("de", "page.subtitle", nil)

	fmt.Println(en)
	fmt.Println(de)
	// Output:
	// Welcome, World!
	// Sie haben neue Nachrichten
}
