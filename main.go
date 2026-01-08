package main

import (
	"flag"
	"fmt"
	"iter"
	"os"
	"time"
)

var (
	now  = time.Now().UTC()
	then = now.AddDate(0, 5, 0)

	nowFormatted  = now.Format(time.RFC3339)
	thenFormatted = then.Format(time.RFC3339)

	startStr    = flag.String("start", nowFormatted, "Start of the date chunker")
	intervalStr = flag.String("interval", "1m", "Interval: 1h, 1d, 1m, 1y, etc")
	endStr      = flag.String("end", thenFormatted, "end date (inclusive)")
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()

	start, err := time.Parse(time.RFC3339, *startStr)
	if err != nil {
		return fmt.Errorf("error parsing start date: %w", err)
	}

	end, err := time.Parse(time.RFC3339, *endStr)
	if err != nil {
		return fmt.Errorf("error parsing end date: %w", err)
	}

	interval, err := parseInterval(*intervalStr)
	if err != nil {
		return fmt.Errorf("error parsing duration: %w", err)
	}

	for date := range generate(start, end, interval) {
		fmt.Println(date.Format(time.RFC3339))
	}

	return nil
}

func generate(start, end time.Time, interval Interval) iter.Seq[time.Time] {
	return iter.Seq[time.Time](func(yield func(time.Time) bool) {
		cur := start
		for {
			if !yield(cur) {
				break
			}

			cur = interval.Apply(cur)
			if cur.After(end) {
				break
			}
		}
	})
}
