// denread prints local files inside one root.
//
//	denread FILE
//	denread -batch FILE FILE
//	denread -overview FILE
//
// Lines are numbered. A batch is several files in one call.
// An overview names headings, symbols, and related files that exist.
// It does not copy the body. A path outside the root is refused.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tcpsystems/den/internal/read"
)

func main() {
	batch := flag.Bool("batch", false, "print each named file with line numbers")
	overview := flag.Bool("overview", false, "print a content-light picture of one file")
	root := flag.String("root", ".", "directory the paths must stay inside")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "usage: denread [-root DIR] FILE")
		fmt.Fprintln(os.Stderr, "       denread [-root DIR] -batch FILE...")
		fmt.Fprintln(os.Stderr, "       denread [-root DIR] -overview FILE")
		os.Exit(2)
	}
	if *overview {
		if len(args) != 1 || *batch {
			fmt.Fprintln(os.Stderr, "denread: -overview takes one file")
			os.Exit(2)
		}
		ov, err := read.Condense(*root, args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		read.WriteOverview(os.Stdout, ov)
		return
	}
	var err error
	if *batch {
		err = read.Batch(os.Stdout, *root, args)
	} else if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "denread: one file, or -batch")
		os.Exit(2)
	} else {
		err = read.Numbered(os.Stdout, *root, args[0])
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, strings.TrimSpace(err.Error()))
		os.Exit(1)
	}
}
