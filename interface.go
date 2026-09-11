package i18n

import (
	"text/template"
)

// DefaultI18nName is the template function name used by Register when no
// explicit name is given.
const DefaultI18nName = "i18n"

// DefaultPluralName is the template function name used by Register when no
// explicit name is given.
const DefaultPluralName = "plural"

type Lang string
type Key string

type Dictionary interface {
	Lang() Lang
	Set(key Key, text string) error

	// Keys returns all keys in the dictionary.
	// The order is not specified. Sort the result if a stable order is needed.
	Keys() []Key
	Execute(key Key, data interface{}) (string, error)
}

type Catalog interface {
	Set(lang Lang, key Key, text string) error

	Load(
		extractor Extractor,
		source map[string]any,
	) error
	LoadBySource(
		extractor Extractor,
		f Source,
	) error
	LoadBySpec(spec Spec) error

	Dictionary(lang Lang) (Dictionary, error)

	// Languages returns all languages in the catalog.
	// The order is not specified. Sort the result if a stable order is needed.
	Languages() []Lang

	// Register wires the translation function into one or more templates.
	// Optional first arg is the function name (default "i18n"); the rest are
	// *html/template.Template or *text/template.Template.
	Register(args ...any) error

	Execute(lang Lang, key Key, data interface{}) (string, error)
}

type Plural interface {
	SetCatalog(catalog Catalog) Plural

	// Register wires the plural function into one or more templates. Optional
	// first arg is the function name (default "plural"); the rest are
	// *html/template.Template or *text/template.Template.
	Register(args ...any) error

	Execute(lang Lang, key Key, value int) (string, error)

	Error() error
}

type Spec interface {
	Extractor() (Extractor, error)
	Source() (map[string]any, error)
}

type Extractor func(path []string) (lang Lang, key Key, err error)
type Source func() (map[string]any, error)
type Extension func(lang Lang) template.FuncMap

func Func(name string, build func(lang Lang) any) Extension {
	return func(lang Lang) template.FuncMap {
		return template.FuncMap{name: build(lang)}
	}
}

func Prefix(prefix string, ext Extension) Extension {
	return func(lang Lang) template.FuncMap {
		src := ext(lang)
		out := make(template.FuncMap, len(src))
		for name, fn := range src {
			out[prefix+"_"+name] = fn
		}
		return out
	}
}
