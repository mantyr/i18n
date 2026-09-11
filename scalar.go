package i18n

import (
	"fmt"
)

func scalarToString(v any) (string, error) {
	switch s := v.(type) {
	case string:
		return s, nil
	case int:
		return fmt.Sprintf("%d", s), nil
	case int64:
		return fmt.Sprintf("%d", s), nil
	case float64:
		return fmt.Sprintf("%v", s), nil
	case bool:
		return fmt.Sprintf("%v", s), nil
	default:
		return "", fmt.Errorf("expected string but actual %T", v)
	}
}
