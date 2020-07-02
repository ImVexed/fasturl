## Fast URL
FastURL is a Go URL parser using a [Ragel](http://www.colm.net/open-source/ragel/) state-machine instead of regex, or the built in standard library `url.Parse`.

## Why?
*S P E E D*

## Benchmarks
## ns/op
![](/_images/ns.svg)
## B/op
![](/_images/b.svg)

## Raw:
```
goos: linux
goarch: amd64
pkg: github.com/ImVexed/fasturl
BenchmarkRegex-32         464509              2557 ns/op             530 B/op          3 allocs/op
BenchmarkRagel-32        5350304               225 ns/op              96 B/op          1 allocs/op
BenchmarkStd-32          2225313               537 ns/op             128 B/op          1 allocs/op
PASS
ok      github.com/ImVexed/fasturl      4.405s
```

## How does this work?
Lots of goto's and determinism, feel free to zoom
![](/_images/graph.svg)


## Credits
[maximecaron](https://github.com/maximecaron) - Creating the initial ragael parser
