package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	file  = flag.String("file", "", "output file")
	label = flag.String("label", "", "string label")
	sep   = flag.String("sep", "\t", "separator")
	line  = flag.Int("line", 10, "line")
	size  = flag.Int("size", 20, "size")
	max   = flag.Int("max", 10000, "max")
)

func main() {
	flag.Parse()

	if *line <= 1 || *size <= 1 {
		return
	}

	rand.Seed(time.Now().UnixNano())

	buf := make([]string, *size)

	var out *bufio.Writer
	if *file != "" {
		open, err := os.Create(*file)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		out = bufio.NewWriter(open)
		defer open.Close()
	} else {
		out = bufio.NewWriter(os.Stdout)
	}
	defer out.Flush()

	for l := 1; l <= *line; l++ {
		i, j := 0, 0
		for ; i < *size; i++ {
			n := rand.Int31n(int32(*max))
			buf[i] = fmt.Sprint(n)
			j++
		}
		if *label != "" {
			fmt.Fprintf(out, "%s\t%s\n", fmt.Sprintf("%s-%d", *label, l), strings.Join(buf[:j], *sep))
		} else {
			fmt.Fprintf(out, "%d\t%s\n", l, strings.Join(buf[:j], *sep))
		}
	}
}
