# Timeslicer

Small utility application to slice time. Give it a start date, end date and an interval and it'll spit out all the timestamps in between start and end in RFC3339 format.

## Usage

```console
Usage:
    timeslicer [flags]

Flags:
    -start string
        Start of the date chunker (default "<now>")
    -interval string
        Interval between chunks (e.g. 1h, 1d, 1m, 1y) (default "1m")
    -end string
        End date (inclusive) (default "<then>")
```

## Installation

```console
go install github.com/tmw/timeslicer
```
