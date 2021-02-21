# CByteBuf

Comparison benchmarks between [cbytebuf](https://github.com/koykov/cbytebuf),
[bytebufferpool](https://github.com/valyala/bytebufferpool), [bytes.Buffer](https://golang.org/pkg/bytes/#Buffer) and
byte slices ([]byte).

## Benchmarks

```
BenchmarkByteArray_Append-8             	  767320	      1449 ns/op	    2040 B/op	       8 allocs/op
BenchmarkByteArray_AppendLong-8         	    1557	    754013 ns/op	 4646288 B/op	      25 allocs/op
BenchmarkByteBufferNative_Write-8       	  517546	      2376 ns/op	    2416 B/op	       5 allocs/op
BenchmarkByteBufferNative_WriteLong-8   	    3441	    346512 ns/op	 1646722 B/op	      10 allocs/op
BenchmarkByteBufferPool_Write-8         	  904567	      1335 ns/op	       0 B/op	       0 allocs/op
BenchmarkByteBufferPool_WriteLong-8     	    1555	    754847 ns/op	 4667398 B/op	      29 allocs/op
BenchmarkCByteBuf_Write-8               	  380574	      3171 ns/op	       0 B/op	       0 allocs/op
BenchmarkCByteBuf_WriteLong-8           	    2631	    454779 ns/op	       0 B/op	       0 allocs/op
```
