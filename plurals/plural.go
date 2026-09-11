package plurals

import (
	"errors"
	"fmt"

	"github.com/mantyr/i18n"
	pl "golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
)

// pluralOperandModulo bounds a plural operand so it fits in an int. x/text has
// no exported constant for this; the value comes from the MatchPlural doc
// comment: "If any of the operand values is too large to fit in an int, it is
// okay to pass the value modulo 10,000,000."
const pluralOperandModulo = 10000000

type plural struct {
	err     error
	catalog i18n.Catalog
	rules   *pl.Rules
}

// NewCardinal - Cardinal defines the plural rules for numbers indicating quantities.
func NewCardinal() i18n.Plural {
	return New(pl.Cardinal)
}

// NewOrdinal - Ordinal defines the plural rules for numbers indicating position
// (first, second, etc.).
func NewOrdinal() i18n.Plural {
	return New(pl.Ordinal)
}

func New(rules *pl.Rules) i18n.Plural {
	p := &plural{}
	if rules == nil {
		p.err = errors.New("empty rules")
		return p
	}
	p.rules = rules
	return p
}

func (p *plural) SetCatalog(catalog i18n.Catalog) i18n.Plural {
	if p.err != nil {
		return p
	}
	if catalog == nil {
		p.err = errors.New("empty catalog")
		return p
	}
	p.catalog = catalog
	return p
}

// Register wires the plural function into one or more templates so they can
// call it directly.
//
// The optional first argument is the function name (default "plural"); the
// rest are templates, which may be *html/template.Template or
// *text/template.Template, in any number:
//
//	pl.Register(tpl)               // {{plural .Lang "key" .N}}
//	pl.Register("pl", tpl1, tpl2)  // {{pl .Lang "key" .N}} in two templates
//
// The language is passed explicitly in the template.
func (p *plural) Register(args ...any) error {
	pluralName, templates := i18n.SplitNameAndTemplates(i18n.DefaultPluralName, args)
	funcs := map[string]any{
		pluralName: p.Execute,
	}
	return i18n.RegisterFuncs(funcs, templates...)
}

func (p *plural) Execute(lang i18n.Lang, key i18n.Key, value int) (string, error) {
	switch {
	case p.err != nil:
		return "", p.err
	case p.rules == nil:
		return "", errors.New("empty rules")
	}
	tag, err := language.Parse(string(lang))
	if err != nil {
		return "", fmt.Errorf("unexpected lang (%s)", string(lang))
	}
	if value < 0 {
		value = -value
	}
	form := name(p.rules.MatchPlural(tag, value%pluralOperandModulo, 0, 0, 0, 0))
	return p.catalog.Execute(lang, i18n.Key(string(key)+"."+form), value)
}

func (p *plural) Error() error {
	return p.err
}

func name(f pl.Form) string {
	switch f {
	case pl.Zero:
		return "zero"
	case pl.One:
		return "one"
	case pl.Two:
		return "two"
	case pl.Few:
		return "few"
	case pl.Many:
		return "many"
	default:
		return "other"
	}
}
