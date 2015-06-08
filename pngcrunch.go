// github.com/rubicks/pngcrunch/pngcrunch.go

package main

import (
	"flag"
	"fmt"
	"image/png"
	"io"
	"os"
)

var help bool
var version bool
var enc png.Encoder = png.Encoder{CompressionLevel: png.BestCompression}

func init() {
	flag.BoolVar(&help, "help", false, "print this message and exit")
	flag.BoolVar(&version, "version", false, "print version and exit")
}

func rdr(in string) (io.Reader, error) {
	if "-" == in {
		return os.Stdin, nil
	}
	return os.Open(in)
}

func wtr(out string) (io.Writer, error) {
	if "-" == out {
		return os.Stdout, nil
	}
	return os.Create(out)
}

func crunch(in, out string) (ret int, err error) {
	rdr, err := rdr(in)
	if nil != err {
		return
	}
	wtr, err := wtr(out)
	if nil != err {
		return
	}
	img, err := png.Decode(rdr)
	if nil != err {
		return
	}
	err = enc.Encode(wtr, img)

	return
}

func usage(out io.Writer) {
	fmt.Fprintf(out, "usage: %s [input file] [output file]\n", os.Args[0])
	fmt.Fprintf(out, "  '-' input file means stdin, the default\n")
	fmt.Fprintf(out, "  '-' output file means stdout, the default\n")
}

func fixargs(args []string) []string {
	return append(args, []string{"-", "-"}...)[:2]
}

func main() {
	flag.Parse()
	if help {
		usage(os.Stdout)
		flag.PrintDefaults()
		os.Exit(0)
	}
	if version {
		fmt.Printf("%s v0.1.0\n", os.Args[0])
		os.Exit(0)
	}

	args := fixargs(flag.Args())

	crunch(args[0], args[1])
}
