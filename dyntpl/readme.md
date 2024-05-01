# DynTpl

Comparison benchmarks between [dyntpl](https://github.com/koykov/dyntpl), [quicktemplate](https://github.com/valyala/quicktemplate), [html/template](https://golang.org/pkg/html/template/) and [text/template](https://golang.org/pkg/text/template/).

```
BenchmarkDyntpl/json/1-8   	             4373738	     269.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/json/10-8             	  901646	      1356 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/json/100-8 	               96657	     12378 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/json/1000-8         	    8516	    123768 ns/op	      12 B/op	       0 allocs/op
BenchmarkDyntpl/xml/1-8             	 6559135         178.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/xml/10-8            	 1513524	     817.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/xml/100-8           	  168139	      7373 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/xml/1000-8          	   16435	     72869 ns/op	      45 B/op	       0 allocs/op
BenchmarkDyntpl/text/1-8            	 7424582	     166.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/text/10-8           	  997195	      1146 ns/op	       0 B/op	       0 allocs/op
BenchmarkDyntpl/text/100-8          	  109776	     10800 ns/op	       1 B/op	       0 allocs/op
BenchmarkDyntpl/text/1000-8         	   10000	    109950 ns/op	       8 B/op	       0 allocs/op
BenchmarkNative/json/1-8            	11909073	     99.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkNative/json/10-8           	 3503720	     339.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkNative/json/100-8          	  430958	      2699 ns/op	       0 B/op	       0 allocs/op
BenchmarkNative/json/1000-8         	   37149	     30720 ns/op	       2 B/op	       0 allocs/op
BenchmarkNative/xml/1-8             	  867600	      1438 ns/op	    4657 B/op	       9 allocs/op
BenchmarkNative/xml/10-8            	  361359	      3271 ns/op	    4657 B/op	       9 allocs/op
BenchmarkNative/xml/100-8           	   53541	     22169 ns/op	    4659 B/op	       9 allocs/op
BenchmarkNative/xml/1000-8          	    5577	    209045 ns/op	    7620 B/op	     909 allocs/op
BenchmarkNative/html/1-8            	 1255239	     942.3 ns/op	     440 B/op	      21 allocs/op
BenchmarkNative/html/10-8           	  236959	      4828 ns/op	    1953 B/op	     102 allocs/op
BenchmarkNative/html/100-8          	   23272	     46188 ns/op	   19252 B/op	    1047 allocs/op
BenchmarkNative/html/1000-8         	    2593	    461264 ns/op	  198173 B/op	   11242 allocs/op
BenchmarkNative/text/1-8            	 3206125	     361.6 ns/op	     120 B/op	       7 allocs/op
BenchmarkNative/text/10-8           	  635608	      1845 ns/op	     352 B/op	      32 allocs/op
BenchmarkNative/text/100-8          	   64558	     17336 ns/op	    2875 B/op	     302 allocs/op
BenchmarkNative/text/1000-8         	    6388	    182174 ns/op	   34065 B/op	    3746 allocs/op
BenchmarkQuickTemplate/json/1-8     	12987396	     79.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/json/10-8    	 4154376	     273.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/json/100-8   	  561513	      2205 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/json/1000-8  	   55358	     21621 ns/op	       1 B/op	       0 allocs/op
BenchmarkQuickTemplate/xml/1-8      	18088328	     71.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/xml/10-8     	 4679110	     232.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/xml/100-8    	  662308	      1670 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/xml/1000-8   	   66312	     18162 ns/op	       1 B/op	       0 allocs/op
BenchmarkQuickTemplate/text/1-8     	27541401	     47.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/text/10-8    	 7403982	     161.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/text/100-8   	  701368	      1938 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickTemplate/text/1000-8  	   81804	     14855 ns/op	       4 B/op	       0 allocs/op
```
