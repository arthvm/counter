package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"text/tabwriter"

	"github.com/arthvm/counter"
	"github.com/arthvm/counter/display"
)

type FileCountsResult struct {
	counts   counter.Counts
	filename string
}

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

	wg := sync.WaitGroup{}
	wg.Add(len(filenames))

	ch := make(chan FileCountsResult)

	for _, filename := range filenames {
		go func() {
			defer wg.Done()

			counts, err := counter.CountFile(filename)
			if err != nil {
				didError = true
				fmt.Fprintln(os.Stderr, "counter:", err)
				return
			}

			ch <- FileCountsResult{
				counts:   counts,
				filename: filename,
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		totals = totals.Add(res.counts)
		res.counts.Print(wr, opts, res.filename)

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
