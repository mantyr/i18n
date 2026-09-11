package sources

import (
	"errors"
	"fmt"
	"os"

	"github.com/mantyr/i18n"
	"gopkg.in/yaml.v3"
)

func YAMLFile(address string) i18n.Source {
	return func() (map[string]any, error) {
		if address == "" {
			return nil, errors.New("empty address")
		}
		data, err := os.ReadFile(address)
		if err != nil {
			return nil, fmt.Errorf("read file error: %s", err.Error())
		}
		var result map[string]any
		err = yaml.Unmarshal(data, &result)
		if err != nil {
			return nil, fmt.Errorf("parse yaml error: %s", err.Error())
		}
		return result, nil
	}
}
