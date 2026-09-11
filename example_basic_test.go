package i18n_test

import (
	"fmt"

	"github.com/mantyr/i18n"
)

// Example_basic shows the smallest possible use: create a catalog, add
// messages for a couple of languages and render one with data.
// nolint:errcheck
func Example_basic() {
	c, _ := i18n.NewCatalog()
	c.Set("en", "greeting", "Hello, {{.Name}}!")
	c.Set("de", "greeting", "Hallo, {{.Name}}!")

	en, _ := c.Execute("en", "greeting", struct{ Name string }{Name: "World"})
	de, _ := c.Execute("de", "greeting", struct{ Name string }{Name: "Welt"})

	fmt.Println(en)
	fmt.Println(de)
	// Output:
	// Hello, World!
	// Hallo, Welt!
}
