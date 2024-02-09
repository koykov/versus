## Core i7-7700HQ (no pool)

```
BenchmarkParseFastJSON/small-8         	 9719614	       115.7 ns/op	1642.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseFastJSON/medium-8        	 1264318	       933.6 ns/op	2494.74 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseFastJSON/large-8         	   93084	     11317 ns/op	2484.51 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseFastJSON/canada-8        	     626	   1860054 ns/op	1210.21 MB/s	       7 B/op	       0 allocs/op
BenchmarkParseFastJSON/citm-8          	    1654	    714085 ns/op	2418.77 MB/s	       2 B/op	       0 allocs/op
BenchmarkParseFastJSON/twitter-8       	    4732	    266328 ns/op	2371.19 MB/s	       0 B/op	       0 allocs/op

BenchmarkParseJsonvector/small-8     	 4942929	       236.0 ns/op	 805.03 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseJsonvector/medium-8    	  619875	      1910 ns/op	1219.48 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseJsonvector/large-8     	   41522	     28695 ns/op	 979.90 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseJsonvector/canada-8    	     154	   7097310 ns/op	 317.17 MB/s	      30 B/op	       0 allocs/op
BenchmarkParseJsonvector/citm-8      	     687	   1686870 ns/op	1023.91 MB/s	       3 B/op	       0 allocs/op
BenchmarkParseJsonvector/twitter-8   	    1897	    614699 ns/op	1027.35 MB/s	       2 B/op	       0 allocs/op

BenchmarkParseSimdjson/small-8         	  110518	     10390 ns/op	  18.29 MB/s	  108808 B/op	       9 allocs/op
BenchmarkParseSimdjson/medium-8        	   93631	     12289 ns/op	 189.52 MB/s	  114872 B/op	      11 allocs/op
BenchmarkParseSimdjson/large-8         	   22711	     50925 ns/op	 552.15 MB/s	  244250 B/op	      14 allocs/op
BenchmarkParseSimdjson/canada-8        	     222	   4812389 ns/op	 467.76 MB/s	 3040797 B/op	      10 allocs/op
BenchmarkParseSimdjson/citm-8          	    1141	   1114941 ns/op	1549.14 MB/s	 2721306 B/op	      11 allocs/op
BenchmarkParseSimdjson/twitter-8       	    2259	    483359 ns/op	1306.51 MB/s	 1828377 B/op	      13 allocs/op
```

## Core i7-7700HQ (pool)

```
BenchmarkParseFastJSON/small-8         	10463110	       117.9 ns/op	1611.87 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseFastJSON/medium-8        	 1295306	       989.3 ns/op	2354.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseFastJSON/large-8         	  102326	     12002 ns/op	2342.70 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseFastJSON/canada-8        	     619	   1995162 ns/op	1128.25 MB/s	       7 B/op	       0 allocs/op
BenchmarkParseFastJSON/citm-8          	    1737	    691977 ns/op	2496.04 MB/s	       2 B/op	       0 allocs/op
BenchmarkParseFastJSON/twitter-8       	    4844	    247732 ns/op	2549.18 MB/s	       0 B/op	       0 allocs/op

BenchmarkParseJsonvector/small-8       	 5008723	       236.0 ns/op	 804.94 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseJsonvector/medium-8      	  634435	      1988 ns/op	1171.43 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseJsonvector/large-8       	   41658	     30143 ns/op	 932.81 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseJsonvector/canada-8      	     158	   7295810 ns/op	 308.54 MB/s	      10 B/op	       0 allocs/op
BenchmarkParseJsonvector/citm-8        	     589	   1731708 ns/op	 997.40 MB/s	       3 B/op	       0 allocs/op
BenchmarkParseJsonvector/twitter-8     	    1640	    623693 ns/op	1012.54 MB/s	       1 B/op	       0 allocs/op

BenchmarkParseSimdjson/small-8         	 3942674	       307.9 ns/op	 617.05 MB/s	      16 B/op	       1 allocs/op
BenchmarkParseSimdjson/medium-8        	  611296	      1824 ns/op	1276.92 MB/s	      16 B/op	       1 allocs/op
BenchmarkParseSimdjson/large-8         	   53648	     22228 ns/op	1264.97 MB/s	      64 B/op	       3 allocs/op
BenchmarkParseSimdjson/canada-8        	     242	   4602677 ns/op	 489.07 MB/s	      88 B/op	       3 allocs/op
BenchmarkParseSimdjson/citm-8          	    1453	    795852 ns/op	2170.26 MB/s	      65 B/op	       3 allocs/op
BenchmarkParseSimdjson/twitter-8       	    3672	    299870 ns/op	2105.96 MB/s	      64 B/op	       3 allocs/op
```

## Xeon Silver 4214

```
BenchmarkParseFastJSON/small-48         	30428878	        34.01 ns/op	5587.31 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseFastJSON/medium-48        	 5725995	       196.0 ns/op	11884.31 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseFastJSON/large-48         	  544407	      2244 ns/op	12532.36 MB/s	      10 B/op	       0 allocs/op
BenchmarkParseFastJSON/canada-48        	    1246	   1075575 ns/op	2092.88 MB/s	  654633 B/op	    1009 allocs/op
BenchmarkParseFastJSON/citm-48          	    3562	    349464 ns/op	4942.43 MB/s	   66965 B/op	     103 allocs/op
BenchmarkParseFastJSON/twitter-48       	   16213	     73763 ns/op	8561.41 MB/s	       1 B/op	       0 allocs/op

BenchmarkParseJsonvector/small-48       	25897755	        47.86 ns/op	3969.67 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseJsonvector/medium-48      	 3483019	       341.1 ns/op	6826.91 MB/s	       0 B/op	       0 allocs/op
BenchmarkParseJsonvector/large-48       	  235570	      5116 ns/op	5495.73 MB/s	      29 B/op	       0 allocs/op
BenchmarkParseJsonvector/canada-48      	      56	  24052117 ns/op	  93.59 MB/s	173683897 B/op	  125471 allocs/op
BenchmarkParseJsonvector/citm-48        	    1450	    816113 ns/op	2116.38 MB/s	  413288 B/op	     339 allocs/op
BenchmarkParseJsonvector/twitter-48     	    4339	    244853 ns/op	2579.15 MB/s	       6 B/op	       0 allocs/op

BenchmarkParseSimdjson/small-48         	12611644	        96.59 ns/op	1967.05 MB/s	      16 B/op	       1 allocs/op
BenchmarkParseSimdjson/medium-48        	 2566564	       476.9 ns/op	4883.43 MB/s	      16 B/op	       1 allocs/op
BenchmarkParseSimdjson/large-48         	  154791	      7638 ns/op	3681.27 MB/s	      83 B/op	       3 allocs/op
BenchmarkParseSimdjson/canada-48        	    1282	    935275 ns/op	2406.83 MB/s	      78 B/op	       3 allocs/op
BenchmarkParseSimdjson/citm-48          	    5502	    219826 ns/op	7857.13 MB/s	      66 B/op	       3 allocs/op
BenchmarkParseSimdjson/twitter-48       	    9332	    130020 ns/op	4857.05 MB/s	     261 B/op	       3 allocs/op
```
