package main

import (
	"bytes"
	"slices"
	"strings"
	"testing"
	"time"
)

func TestPrintAll(t *testing.T) {
	var (
		buf   bytes.Buffer
		start = time.Now().UTC().Add(-5 * time.Hour)
		end   = time.Now().UTC().Add(-1 * time.Hour)
	)

	interval, err := parseInterval("1h")
	if err != nil {
		t.Error("caught error", err)
	}

	printAll(&buf, start, end, interval)

	iter := strings.Lines(buf.String())
	lines := slices.Collect(iter)
	if len(lines) != 5 {
		t.Errorf("expected 5 lines, got %d", len(lines))
	}
}
