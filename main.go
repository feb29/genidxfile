package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	file = flag.String("file", "", "output file")
	line = flag.Int("line", 10, "line")
	size = flag.Int("size", 65000, "size")
	max  = flag.Int("max", 15000000, "max")
)

func main() {
	flag.Parse()

	if *line <= 1 || *size <= 1 {
		return
	}

	rand.Seed(time.Now().UnixNano())

	buf := make([]string, *size)

	var out io.Writer
	if *file != "" {
		open, err := os.Create(*file)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		out = bufio.NewWriter(open)
	} else {
		out = bufio.NewWriter(os.Stdout)
	}

	for l := 1; l <= *line; l++ {
		i, j := 0, 0
		for ; i < *size; i++ {
			n := rand.Int31n(int32(*max))
			buf[i] = fmt.Sprint(n)
			j++
		}
		fmt.Fprintf(out, "%d\t%s\n", l, strings.Join(buf[:j], "\t"))
	}
}
