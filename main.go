package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	headSize = flag.Int("h", 0, "unmask head size")
	tailSize = flag.Int("t", 0, "unmask tail size")
)

func main() {
	flag.Parse()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%s=%s\n", pair[0], MaskedStr(pair[1], *headSize, *tailSize))
	}
}

func MaskedStr(origin string, hs, ts int) string {
	head := origin
	tail := origin
	if len(origin) >= hs {
		head = origin[:hs]
	}
	if len(origin) >= ts {
		tail = origin[len(origin)-ts:]
	}
	masked := ""
	if len(origin)-hs-ts > 0 {
		masked = strings.Repeat("*", len(origin)-hs-ts)
	}
	return fmt.Sprintf("%s%s%s", head, masked, tail)
}
