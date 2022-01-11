package stringer

import "github.com/koykov/fastconv"

type Language int

func (l Language) Name() string {
	if l < 0 || l >= Language(len(__lt_lst)) {
		return ""
	}
	lo, hi := __lt_lst[l].name.Decode()
	return fastconv.B2S(__lt_buf[lo:hi])
}

func (l Language) Native() string {
	if l < 0 || l >= Language(len(__lt_lst)) {
		return ""
	}
	lo, hi := __lt_lst[l].native.Decode()
	return fastconv.B2S(__lt_buf[lo:hi])
}

func (l Language) Iso6391() string {
	if l < 0 || l >= Language(len(__lt_lst)) {
		return ""
	}
	lo, hi := __lt_lst[l].iso1.Decode()
	return fastconv.B2S(__lt_buf[lo:hi])
}

func (l Language) Iso6393() string {
	if l < 0 || l >= Language(len(__lt_lst)) {
		return ""
	}
	lo, hi := __lt_lst[l].iso3.Decode()
	return fastconv.B2S(__lt_buf[lo:hi])
}
