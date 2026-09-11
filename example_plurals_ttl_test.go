package i18n_test

import (
	"fmt"
	"time"

	"github.com/mantyr/i18n"
	"github.com/mantyr/i18n/humanize/ttl"
	"github.com/mantyr/i18n/plurals"
)

// Example_durationInWords shows the ttl + plural subpackages working together:
// split a time.Duration into parts, then render each part with the correct
// plural form for the language. Russian is a good demo because it has three
// forms (one/few/many).
// nolint:errcheck
func Example_durationInWords() {
	c, _ := i18n.NewCatalog()
	c.Set("ru", "ttl.days.one", "{{.}} день")
	c.Set("ru", "ttl.days.few", "{{.}} дня")
	c.Set("ru", "ttl.days.many", "{{.}} дней")
	c.Set("ru", "ttl.hours.one", "{{.}} час")
	c.Set("ru", "ttl.hours.few", "{{.}} часа")
	c.Set("ru", "ttl.hours.many", "{{.}} часов")

	pl := plurals.NewCardinal().SetCatalog(c)

	parts, _ := ttl.Split(50 * time.Hour) // 2 days 2 hours

	days, _ := pl.Execute("ru", "ttl.days", parts.Days)
	hours, _ := pl.Execute("ru", "ttl.hours", parts.Hours)

	fmt.Println(days)
	fmt.Println(hours)
	// Output:
	// 2 дня
	// 2 часа
}
