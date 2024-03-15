package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	fs := flag.NewFlagSet("xd", flag.ContinueOnError)

	n := fs.Int("n", 0, "how many bytes to process (default 0, unlimited)")
	s := fs.Int("s", 0, "how many bytes to skip (default 0)")
	file := fs.String("f", "", "file to read (default stdin)")
	outFile := fs.String("o", "", "file to write (default stdout)")

	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	in := io.Reader(os.Stdin)
	out := io.Writer(os.Stdout)

	if !fromStdin() {
		if *file == "" {
			panic("no input")
		}
		f, err := os.Open(*file)
		if err != nil {
			panic(fmt.Sprintf("Failed to open %s: %s", *file, err))
		}
		defer f.Close()

		in = f
	}
	if *outFile != "" {
		f, err := os.Create(*outFile)
		if err != nil {
			panic(fmt.Sprintf("Failed to create %s: %s", *outFile, err))
		}
		defer f.Close()

		out = f
	}

	if *s != 0 {
		io.CopyN(io.Discard, in, int64(*s))
	}
	if *n != 0 {
		in = io.LimitReader(in, int64(*n))
	}
	dumper := hex.Dumper(out)
	io.Copy(dumper, in)
}

func fromStdin() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Size() > 0 || fi.Mode()&os.ModeNamedPipe != 0
}
