package i18n

import (
	"errors"
	"fmt"
	"github.com/mantyr/i18n/internal/registers"
)

type catalog struct {
	data       map[Lang]Dictionary
	extensions []Extension
}

func NewCatalog(extensions ...Extension) (Catalog, error) {
	c := &catalog{
		data:       make(map[Lang]Dictionary),
		extensions: extensions,
	}
	return c, nil
}

func (c *catalog) Set(lang Lang, key Key, text string) error {
	switch {
	case lang == "":
		return errors.New("empty lang")
	case key == "":
		return errors.New("empty key")
	}
	d, ok := c.data[lang]
	if !ok {
		var err error
		d, err = NewDictionary(lang, c.extensions...)
		if err != nil {
			return err
		}
		c.data[lang] = d
	}
	return d.Set(key, text)
}

func (c *catalog) Load(
	extractor Extractor,
	source map[string]any,
) error {
	switch {
	case extractor == nil:
		return errors.New("empty extractor")
	case source == nil:
		return errors.New("empty source")
	}
	var err error
	for key, value := range source {
		err = c.walk(extractor, []string{key}, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *catalog) LoadBySource(extractor Extractor, f Source) error {
	switch {
	case extractor == nil:
		return errors.New("empty extractor")
	case f == nil:
		return errors.New("empty source")
	}
	source, err := f()
	if err != nil {
		return fmt.Errorf("load source error: %s", err.Error())
	}
	return c.Load(extractor, source)
}

func (c *catalog) LoadBySpec(spec Spec) error {
	if spec == nil {
		return fmt.Errorf("empty spec")
	}
	extractor, err := spec.Extractor()
	if err != nil {
		return err
	}
	items, err := spec.Source()
	if err != nil {
		return err
	}
	if items == nil {
		return fmt.Errorf("empty items")
	}
	return c.Load(extractor, items)
}

func (c *catalog) walk(extractor Extractor, path []string, node any) error {
	switch n := node.(type) {
	case map[string]any:
		var err error
		for key, item := range n {
			if key == "" {
				return errors.New("empty key in source map")
			}
			// NOTE: append(path, key) may reuse path's backing array, which
			// would normally be a slice-aliasing hazard. It is safe here
			// because the path is consumed immediately: for a scalar node the
			// extractor turns it into a string key right away, before the next
			// append overwrites the shared array. The path is never retained,
			// so reuse across sibling keys causes no corruption.
			err = c.walk(extractor, append(path, key), item)
			if err != nil {
				return err
			}
		}
		return nil
	case nil:
		return errors.New("expected template but actual nil")
	default:
		text, err := scalarToString(node)
		if err != nil {
			return err
		}
		lang, key, err := extractor(path)
		if err != nil {
			return err
		}
		return c.Set(lang, key, text)
	}
}

func (c *catalog) Dictionary(lang Lang) (Dictionary, error) {
	d, ok := c.data[lang]
	if !ok {
		return nil, errors.New("dictionary not found")
	}
	return d, nil
}

func (c *catalog) Languages() []Lang {
	result := make([]Lang, 0, len(c.data))
	for lang := range c.data {
		result = append(result, lang)
	}
	return result
}

// Register wires the catalog's translation function into one or more
// templates so they can call it directly.
//
// The optional first argument is the function name (default "i18n"); the rest
// are templates, which may be *html/template.Template or
// *text/template.Template, in any number:
//
//	msgs.Register(tpl)               // {{i18n .Lang "key" .}}
//	msgs.Register("t", tpl1, tpl2)   // {{t .Lang "key" .}} in two templates
//
// The language is passed explicitly in the template, so one registration
// serves every language. Call before Parse, since Funcs must be set before
// the template is parsed.
func (c *catalog) Register(args ...any) error {
	name, templates := registers.SplitNameAndTemplates(DefaultI18nName, args)
	funcs := map[string]any{
		name: c.Execute,
	}
	return registers.RegisterFuncs(funcs, templates...)
}

func (c *catalog) Execute(lang Lang, key Key, data interface{}) (string, error) {
	d, ok := c.data[lang]
	if !ok {
		return "", errors.New("dictionary not found")
	}
	return d.Execute(key, data)
}
