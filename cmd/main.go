package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/arthvm/counter"
	"github.com/arthvm/counter/display"
)

func main() {
	opts := display.Options{}

	flag.BoolVar(
		&opts.ShowWords,
		"w",
		false,
		"Used to toggle whether or not to show the word count",
	)

	flag.BoolVar(
		&opts.ShowLines,
		"l",
		false,
		"Used to toggle whether or not to show the line count",
	)

	flag.BoolVar(
		&opts.ShowBytes,
		"c",
		false,
		"Used to toggle whether or not to show the byte count",
	)

	flag.BoolVar(
		&opts.ShowHeader,
		"header",
		false,
		"Used to toggle whether or not to show a header for the counts",
	)

	flag.Parse()

	log.SetFlags(0)

	wr := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	totals := counter.Counts{}

	filenames := flag.Args()
	didError := false

	for _, filename := range filenames {
		counts, err := counter.CountFile(filename)
		if err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", err)
			continue
		}

		totals = totals.Add(counts)

		counts.Print(wr, opts, filename)
		opts.ShowHeader = false
	}

	if len(filenames) == 0 {
		counter.GetCounts(os.Stdin).Print(wr, opts)
	}

	if len(filenames) > 1 {
		totals.Print(wr, opts, "totals")
	}

	wr.Flush()
	if didError {
		os.Exit(1)
	}
}
