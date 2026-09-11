package methods

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/mantyr/i18n"
	"github.com/mantyr/i18n/extractors"
)

const (
	First Method = "first"
	Last  Method = "last"
)

type Method string

func (m *Method) Extractor() (i18n.Extractor, error) {
	switch *m {
	case First:
		return extractors.First, nil
	case Last:
		return extractors.Last, nil
	default:
		return nil, fmt.Errorf("unexpected method (%s) for extractor", *m)
	}
}

func (m *Method) UnmarshalJSON(data []byte) error {
	*m = Method(strings.ToLower(strings.TrimSpace(strings.Trim(string(data), `"`))))
	return nil
}

func (m *Method) UnmarshalYAML(value *yaml.Node) error {
	var s string
	if err := value.Decode(&s); err != nil {
		return err
	}
	*m = Method(strings.ToLower(strings.TrimSpace(s)))
	return nil
}
