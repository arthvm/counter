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
	args := display.NewOptionsArgs{}

	flag.BoolVar(
		&args.ShowWords,
		"w",
		false,
		"Used to toggle whether or not to show the word count",
	)

	flag.BoolVar(
		&args.ShowLines,
		"l",
		false,
		"Used to toggle whether or not to show the line count",
	)

	flag.BoolVar(
		&args.ShowBytes,
		"c",
		false,
		"Used to toggle whether or not to show the byte count",
	)

	flag.BoolVar(
		&args.ShowHeader,
		"header",
		false,
		"Used to toggle whether or not to show a header for the counts",
	)

	flag.Parse()

	opts := display.NewOptions(args)

	log.SetFlags(0)

	wr := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	totals := counter.Counts{}

	filenames := flag.Args()
	didError := false

	ch := counter.CountFiles(filenames)

	results := make([]counter.FileCountsResult, len(filenames))

	for res := range ch {
		results[res.Idx] = res
	}

	for _, res := range results {
		if res.Err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", res.Err)
			continue
		}

		totals = totals.Add(res.Counts)
		res.Counts.Print(wr, opts, res.Filename)

		args.ShowHeader = false
		opts = display.NewOptions(args)
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
