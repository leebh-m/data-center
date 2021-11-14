package main

import (
	"data-center/collect"
	"flag"
)

func main() {
	collect.Start(*flag.Int("port", 4000, "Set port of the server"))
}
