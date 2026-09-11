package i18n

import (
	"errors"
	"fmt"
	"strings"
	"text/template"
)

type dictionary struct {
	lang       Lang
	data       map[Key]*template.Template
	extensions []Extension
}

func NewDictionary(lang Lang, extensions ...Extension) (Dictionary, error) {
	if lang == "" {
		return nil, errors.New("empty lang")
	}
	d := &dictionary{
		lang:       lang,
		data:       make(map[Key]*template.Template),
		extensions: extensions,
	}
	return d, nil
}

func (d *dictionary) Lang() Lang {
	return d.lang
}

func (d *dictionary) Set(key Key, text string) error {
	if key == "" {
		return errors.New("empty key")
	}
	t := template.New(string(key) + "@" + string(d.lang)).Option("missingkey=error")
	for _, f := range d.extensions {
		t.Funcs(f(d.lang))
	}
	_, err := t.Parse(text)
	if err != nil {
		return err
	}
	d.data[key] = t
	return nil
}

func (d *dictionary) Keys() []Key {
	result := make([]Key, 0, len(d.data))
	for key := range d.data {
		result = append(result, key)
	}
	return result
}

func (d *dictionary) Execute(key Key, data interface{}) (string, error) {
	t, ok := d.data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	var b strings.Builder
	err := t.Execute(&b, data)
	if err != nil {
		return "", fmt.Errorf("render error: %s", err.Error())
	}
	return b.String(), nil
}
