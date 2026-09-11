package extractors

import (
	"fmt"
	"strings"

	"github.com/mantyr/i18n"
)

func First(path []string) (lang i18n.Lang, key i18n.Key, err error) {
	if len(path) < 2 {
		return "", "", fmt.Errorf("expected path >= 2 but actual %d", len(path))
	}
	l := path[0]
	k := path[1:]
	return i18n.Lang(l), i18n.Key(strings.Join(k, ".")), nil
}

func Last(path []string) (lang i18n.Lang, key i18n.Key, err error) {
	if len(path) < 2 {
		return "", "", fmt.Errorf("expected path >= 2 but actual %d", len(path))
	}
	l := path[len(path)-1]
	k := path[:len(path)-1]
	return i18n.Lang(l), i18n.Key(strings.Join(k, ".")), nil
}
