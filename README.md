# Timeslicer

Small utility application to slice time. Give it a start date, end date and an interval and it'll spit out all the timestamps in between start and end in RFC3339 format.

## Installation

```console
go install github.com/tmw/timeslicer
```

## Usage

```console
Usage:
    timeslicer [flags]

Flags:
    -start string
        Start of the timeslicer (default "<now>" eg 2026-01-11T00:00:00Z)

    -end string
        End date (inclusive) (default "<now> + 6 months" eg 2026-05-11T00:00:00Z)

    -interval string
        Interval between chunks (e.g. 1h, 1d, 1m, 1y) (default "1m")
```

## Example

```console
timeslicer\
    -start="2025-01-01T00:00:00Z"\
    -end="2025-05-01T00:00:00Z"\
    -interval="1M"

2025-01-01T00:00:00Z
2025-02-01T00:00:00Z
2025-03-01T00:00:00Z
2025-04-01T00:00:00Z
2025-05-01T00:00:00Z
```

## License

[MIT](./LICENSE)
