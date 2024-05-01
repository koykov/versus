# DynTpl

Comparison benchmarks between [dyntpl](https://github.com/koykov/dyntpl), [quicktemplate](https://github.com/valyala/quicktemplate), [html/template](https://golang.org/pkg/html/template/) and [text/template](https://golang.org/pkg/text/template/).

```
BenchmarkDyntpl/json/1-8   	             4234180	     268.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/json/10-8  	              912177	      1294 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/json/100-8 	              100958	     11716 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/json/1000-8         	   10177	    117188 ns/op	      10 B/op	       0 allocs/op
BenchmarkDyntpl/xml/1-8             	 6889815	     184.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/xml/10-8            	 1628126	     723.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/xml/100-8           	  172864	      6559 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/xml/1000-8          	   17310	     67661 ns/op	       6 B/op	       0 allocs/op
BenchmarkDyntpl/text/1-8                 9159510         124.9 ns/op           0 B/op          0 allocs/op
BenchmarkDyntpl/text/10-8                1703347         703.1 ns/op           0 B/op          0 allocs/op
BenchmarkDyntpl/text/100-8                177800          6436 ns/op           0 B/op          0 allocs/op
BenchmarkDyntpl/text/1000-8                18267         71654 ns/op          32 B/op          0 allocs/op
BenchmarkNative/json/1-8            	12114262	     99.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkNative/json/10-8           	 3485714	     347.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkNative/json/100-8          	  436262	      2679 ns/op	       0 B/op	       0 allocs/op
BenchmarkNative/json/1000-8         	   39279	     30788 ns/op	       2 B/op	       0 allocs/op
BenchmarkNative/xml/1-8             	  905347	      1464 ns/op	    4657 B/op	       9 allocs/op
BenchmarkNative/xml/10-8            	  332524	      3281 ns/op	    4657 B/op	       9 allocs/op
BenchmarkNative/xml/100-8           	   53130	     21190 ns/op	    4660 B/op	       9 allocs/op
BenchmarkNative/xml/1000-8          	    5442	    209890 ns/op	    7566 B/op	     909 allocs/op
BenchmarkNative/html/1-8            	 1209238	     973.6 ns/op	     440 B/op	      21 allocs/op
BenchmarkNative/html/10-8           	  210406	      4792 ns/op	    1954 B/op	     102 allocs/op
BenchmarkNative/html/100-8          	   25087	     46341 ns/op	   19251 B/op	    1047 allocs/op
BenchmarkNative/html/1000-8         	    2245	    469354 ns/op	  198190 B/op	   11242 allocs/op
BenchmarkNative/text/1-8            	 3059950	     370.8 ns/op	     120 B/op	       7 allocs/op
BenchmarkNative/text/10-8           	  594804	      1834 ns/op	     352 B/op	      32 allocs/op
BenchmarkNative/text/100-8          	   66560	     18003 ns/op	    2875 B/op	     302 allocs/op
BenchmarkNative/text/1000-8         	    6274	    177844 ns/op	   34078 B/op	    3746 allocs/op
BenchmarkQuickTemplate/json/1-8     	15303483	     78.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/json/10-8    	 4466312	     276.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/json/100-8   	  569451	      2031 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/json/1000-8  	   52537	     21645 ns/op	       1 B/op	       0 allocs/op
BenchmarkQuickTemplate/xml/1-8      	20859148	     66.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/xml/10-8     	 5966094	     207.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/xml/100-8    	  674106	      1765 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/xml/1000-8   	   66914	     17730 ns/op	       1 B/op	       0 allocs/op
BenchmarkQuickTemplate/text/1-8     	24687188	     43.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/text/10-8    	 7081336	     171.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/text/100-8   	  745024	      1488 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/text/1000-8  	   84392	     14478 ns/op	       0 B/op	       0 allocs/op
```
