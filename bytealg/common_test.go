package bytealg

var (
	trimOrigin = []byte("..foo bar!!???")
	trimExpect = []byte("foo bar")
	trimCutStr = "?!."
	trimCut    = []byte(trimCutStr)

	trimOriginS = "..foo bar!!???"
	trimExpectS = "foo bar"
	trimCutS    = "?!."
)
