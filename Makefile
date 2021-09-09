benchmarks: raw json

raw:
	grep -v -E "(Benchmark|PASS|ok)" benchmarks.txt > site/benchmarks/raw/head.txt
	grep fib10 benchmarks.txt > site/benchmarks/raw/fib10.txt
	grep fib35 benchmarks.txt > site/benchmarks/raw/fib35.txt
	grep closures benchmarks.txt > site/benchmarks/raw/closures.txt
	grep iterations benchmarks.txt > site/benchmarks/raw/iterations.txt
	grep strings benchmarks.txt > site/benchmarks/raw/strings.txt

json:
	cat benchmarks.txt | bench2json -p=fib35 > site/benchmarks/fib35.json
	cat benchmarks.txt | bench2json > site/benchmarks/all.json
