package i18n

import (
	"fmt"

	htmltemplate "html/template"
	texttemplate "text/template"
)

// RegisterFuncs adds funcs to each template in args, working with both
// html/template and text/template.
//
// args is a variadic mix: an optional leading string is ignored here (name
// handling is done by callers), and the remaining elements must be
// *html/template.Template or *text/template.Template. Unknown types are
// skipped. This is the shared engine behind the Register methods.
func RegisterFuncs(funcs map[string]any, args ...any) error {
	for _, arg := range args {
		switch t := arg.(type) {
		case *htmltemplate.Template:
			t.Funcs(htmltemplate.FuncMap(funcs))
		case *texttemplate.Template:
			t.Funcs(texttemplate.FuncMap(funcs))
		default:
			return fmt.Errorf("Register: unsupported argument type %T, want *html/template.Template or *text/template.Template", arg)
		}
	}
	return nil
}

// SplitNameAndTemplates pulls an optional leading function name out of args.
// If the first element is a string it becomes the name and is dropped from
// the template list; otherwise defaultName is used. Exported so subpackages
// (e.g. plurals) can build the same Register signature.
func SplitNameAndTemplates(defaultName string, args []any) (string, []any) {
	name := defaultName
	if len(args) > 0 {
		if s, ok := args[0].(string); ok {
			name = s
			args = args[1:]
		}
	}
	return name, args
}
