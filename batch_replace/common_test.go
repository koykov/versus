package batch_replace

var (
	breplOrigin = []byte("foo {tag0} bar {tag1} string {macro} with {cnt} tags")
	breplExpect = []byte("foo s0 bar long string string 1234567.0987654321 with 4 tags")
	brTag0      = []byte("{tag0}")
	brTag0Val   = []byte("s0")
	brTag1      = []byte("{tag1}")
	brTag1Val   = []byte("long string")
	brTag2      = []byte("{macro}")
	brTag3      = []byte("{cnt}")

	breplOriginS = "foo {tag0} bar {tag1} string {macro} with {cnt} tags"
	breplExpectS = "foo s0 bar long string string 1234567.0987654321 with 4 tags"
	brTag0S      = "{tag0}"
	brTag0ValS   = "s0"
	brTag1S      = "{tag1}"
	brTag1ValS   = "long string"
	brTag2S      = "{macro}"
	brTag3S      = "{cnt}"
)
