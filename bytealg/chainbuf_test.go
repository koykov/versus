package bytealg

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/koykov/bytealg"
)

var (
	cbBytes  = []byte("bytes string")
	cbString = "foobar string"
	cbInt    = int64(7439264324321)
	cbUint   = uint64(9827390546234)
	cbFloat  = 3.1415
	cbExpect = []byte("bytes string-foobar string-7439264324321-9827390546234-3.1415")
)

func BenchmarkChainBuf(b *testing.B) {
	b.ReportAllocs()

	var cb bytealg.ChainBuf
	for i := 0; i < b.N; i++ {
		cb.Reset().
			Write(cbBytes).WriteByte('-').
			WriteStr(cbString).WriteByte('-').
			WriteInt(cbInt).WriteByte('-').
			WriteUint(cbUint).WriteByte('-').
			WriteFloat(cbFloat)

		if !bytes.Equal(cb.Bytes(), cbExpect) {
			b.Error("ChainBuf: mismatch result and expectation")
		}
	}
}

func BenchmarkByteSlice(b *testing.B) {
	b.ReportAllocs()

	var buf []byte
	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		buf = append(buf, cbBytes...)
		buf = append(buf, '-')
		buf = append(buf, cbString...)
		buf = append(buf, '-')
		buf = append(buf, strconv.Itoa(int(cbInt))...)
		buf = append(buf, '-')
		buf = append(buf, strconv.Itoa(int(cbUint))...)
		buf = append(buf, '-')
		buf = append(buf, strconv.FormatFloat(cbFloat, 'f', -1, 64)...)

		if !bytes.Equal(buf, cbExpect) {
			b.Error("ByteArray: mismatch result and expectation")
		}
	}
}
