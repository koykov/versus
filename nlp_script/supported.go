package main

import (
	"unicode"

	"github.com/koykov/nlp"
)

type pair struct {
	a nlp.Script
	b *unicode.RangeTable
}

var (
	scripts = []pair{
		{nlp.ScriptLatin, unicode.Latin},
		{nlp.ScriptArabic, unicode.Arabic},
		{nlp.ScriptCyrillic, unicode.Cyrillic},
		{nlp.ScriptDevanagari, unicode.Devanagari},
		{nlp.ScriptEthiopic, unicode.Ethiopic},
		{nlp.ScriptHan, unicode.Han},
		{nlp.ScriptTagalog, unicode.Tagalog},
		{nlp.ScriptPhags_Pa, unicode.Phags_Pa},
		{nlp.ScriptTelugu, unicode.Telugu},
		{nlp.ScriptHebrew, unicode.Hebrew},
		{nlp.ScriptBopomofo, unicode.Bopomofo},
		{nlp.ScriptMyanmar, unicode.Myanmar},
		{nlp.ScriptBengali, unicode.Bengali},
		{nlp.ScriptDeseret, unicode.Deseret},
		{nlp.ScriptShavian, unicode.Shavian},
		{nlp.ScriptDuployan, unicode.Duployan},
		{nlp.ScriptGeorgian, unicode.Georgian},
		{nlp.ScriptRunic, unicode.Runic},
		{nlp.ScriptGreek, unicode.Greek},
		{nlp.ScriptGujarati, unicode.Gujarati},
		{nlp.ScriptArmenian, unicode.Armenian},
		{nlp.ScriptMahajani, unicode.Mahajani},
		{nlp.ScriptOgham, unicode.Ogham},
		{nlp.ScriptSyriac, unicode.Syriac},
		{nlp.ScriptHiragana, unicode.Hiragana},
		{nlp.ScriptKatakana, unicode.Katakana},
		{nlp.ScriptJavanese, unicode.Javanese},
		{nlp.ScriptKannada, unicode.Kannada},
		{nlp.ScriptKhmer, unicode.Khmer},
		{nlp.ScriptHangul, unicode.Hangul},
		{nlp.ScriptMalayalam, unicode.Malayalam},
		{nlp.ScriptModi, unicode.Modi},
		{nlp.ScriptMongolian, unicode.Mongolian},
		{nlp.ScriptTirhuta, unicode.Tirhuta},
		{nlp.ScriptElbasan, unicode.Elbasan},
		{nlp.ScriptGurmukhi, unicode.Gurmukhi},
		{nlp.ScriptSinhala, unicode.Sinhala},
		{nlp.ScriptOsmanya, unicode.Osmanya},
		{nlp.ScriptTamil, unicode.Tamil},
		{nlp.ScriptThai, unicode.Thai},
	}
)
