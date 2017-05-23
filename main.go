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
	label = flag.String("label", "", "string label")
	line  = flag.Int("line", 10, "line number")
	size  = flag.Int("size", 6500, "size of set")
	max   = flag.Int("max", 15000000, "max value")
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
		if *label != "" {
			fmt.Fprintf(os.Stdout, "%s\t%s\n", fmt.Sprintf("%s-%d", *label, l), strings.Join(slice, "\t"))
		} else {
			fmt.Fprintf(os.Stdout, "%d\t%s\n", l, strings.Join(slice, "\t"))
		}
	}
}
