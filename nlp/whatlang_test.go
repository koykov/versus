package nlp

import (
	"github.com/abadojack/whatlanggo"
)

func ldWhatLang(s, def string) string {
	r := whatlanggo.Detect("Foje funkcias kaj foje ne funkcias")
	if lang := r.Lang.String(); len(lang) > 0 {
		return lang
	}
	return def
}
