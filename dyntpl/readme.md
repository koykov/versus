# DynTpl

Comparison benchmarks between [dyntpl](https://github.com/koykov/dyntpl), [quicktemplate](https://github.com/valyala/quicktemplate), [html/template](https://golang.org/pkg/html/template/) and [text/template](https://golang.org/pkg/text/template/).

```
BenchmarkMarshalJSONDyntpl1-8                2330160           508 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalJSONDyntpl10-8                446022          2581 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalJSONDyntpl100-8                51913         23310 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalJSONDyntpl1000-8                4852        233650 ns/op          16 B/op          0 allocs/op
BenchmarkMarshalXMLDyntpl1-8                 2393229           500 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalXMLDyntpl10-8                 452361          2666 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalXMLDyntpl100-8                 49423         22916 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalXMLDyntpl1000-8                 5077        243913 ns/op          14 B/op          0 allocs/op
BenchmarkDyntpl1-8                           3113560           392 ns/op           0 B/op          0 allocs/op
BenchmarkDyntpl10-8                           716868          1645 ns/op           0 B/op          0 allocs/op
BenchmarkDyntpl100-8                           77823         15318 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalJSONNative1-8                9163930           130 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalJSONNative10-8               2272581           526 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalJSONNative100-8               265914          4550 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalJSONNative1000-8               24528         48666 ns/op           6 B/op          0 allocs/op
BenchmarkMarshalXMLNative1-8                 1924874           618 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalXMLNative10-8                 431286          2812 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalXMLNative100-8                 47599         25205 ns/op           1 B/op          0 allocs/op
BenchmarkMarshalXMLNative1000-8                 4371        271960 ns/op        2906 B/op        900 allocs/op
BenchmarkHTMLTemplate1-8                     1140668          1077 ns/op         576 B/op         21 allocs/op
BenchmarkHTMLTemplate10-8                     204366          5248 ns/op        2673 B/op        111 allocs/op
BenchmarkHTMLTemplate100-8                     22550         52723 ns/op       26446 B/op       1146 allocs/op
BenchmarkTextTemplate1-8                     2924638           399 ns/op         128 B/op          7 allocs/op
BenchmarkTextTemplate10-8                     553068          2420 ns/op         432 B/op         41 allocs/op
BenchmarkTextTemplate100-8                     56626         21288 ns/op        3674 B/op        401 allocs/op
BenchmarkMarshalJSONQuickTemplate1-8        13577886            85.0 ns/op         0 B/op          0 allocs/op
BenchmarkMarshalJSONQuickTemplate10-8        2826133           401 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalJSONQuickTemplate100-8        329229          3594 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalJSONQuickTemplate1000-8        30543         40436 ns/op           2 B/op          0 allocs/op
BenchmarkMarshalXMLQuickTemplate1-8         16197139            69.0 ns/op         0 B/op          0 allocs/op
BenchmarkMarshalXMLQuickTemplate10-8         3934564           320 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalXMLQuickTemplate100-8         465794          2576 ns/op           0 B/op          0 allocs/op
BenchmarkMarshalXMLQuickTemplate1000-8         42067         28501 ns/op           1 B/op          0 allocs/op
BenchmarkQuickTemplate1-8                   19076178            57.4 ns/op         0 B/op          0 allocs/op
BenchmarkQuickTemplate10-8                   4996420           240 ns/op           0 B/op          0 allocs/op
BenchmarkQuickTemplate100-8                   547682          2218 ns/op           0 B/op          0 allocs/op

```
