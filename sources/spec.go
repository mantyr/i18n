package sources

import (
	"errors"
	"fmt"
	"os"

	"github.com/mantyr/i18n"
	"github.com/mantyr/i18n/methods"
	"gopkg.in/yaml.v3"
)

type Spec struct {
	Method methods.Method `json:"method" yaml:"method"`
	Data   map[string]any `json:"items"  yaml:"items"`
}

func (s *Spec) Extractor() (i18n.Extractor, error) {
	return s.Method.Extractor()
}

func (s *Spec) Source() (map[string]any, error) {
	return s.Data, nil
}

// SpecYAMLFile reads a YAML file in the Spec format (method + items) and
// returns the parsed *Spec, ready for Catalog.LoadBySpec:
//
//	spec, err := sources.SpecYAMLFile("./messages.spec.yaml")
//	err = c.LoadBySpec(spec)
//
// The file looks like:
//
//	method: last
//	items:
//	  messages:
//	    greeting:
//	      en: "Hello, {{.Name}}!"
func SpecYAMLFile(address string) (i18n.Spec, error) {
	if address == "" {
		return nil, errors.New("empty address")
	}
	data, err := os.ReadFile(address)
	if err != nil {
		return nil, fmt.Errorf("read file error: %s", err.Error())
	}
	s := &Spec{
		Data: make(map[string]any),
	}
	err = yaml.Unmarshal(data, s)
	if err != nil {
		return nil, fmt.Errorf("parse yaml error: %s", err.Error())
	}
	return s, nil
}
