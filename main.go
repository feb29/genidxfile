package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	line = flag.Int("line", 10, "line")
	size = flag.Int("size", 6500, "size")
	max  = flag.Int("max", 15000000, "max")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()

	for l := 1; l <= *line; l++ {
		slice := make([]string, *size)
		for i := 0; i < *size; i++ {
			n := rand.Int31n(int32(*max))
			slice[i] = fmt.Sprint(n)
		}
		fmt.Fprintf(os.Stdout, "%d\t%s\n", l, strings.Join(slice, "\t"))
	}
}
