package main

import (
	"fmt"
	"time"
)

type intervalUnit byte

const (
	intervalUnitMinute = 'm'
	intervalUnitHour   = 'h'
	intervalUnitDay    = 'd'
	intervalUnitWeek   = 'w'
	intervalUnitMonth  = 'M'
	intervalUnitYear   = 'y'
)

type Interval struct {
	val  int
	unit intervalUnit
}

func (i Interval) String() string {
	return fmt.Sprintf("%d%c", i.val, i.unit)
}

func (i Interval) Apply(d time.Time) time.Time {
	switch i.unit {

	case intervalUnitMinute:
		return d.Add(time.Minute * time.Duration(i.val))

	case intervalUnitHour:
		return d.Add(time.Hour * time.Duration(i.val))

	case intervalUnitDay:
		return d.AddDate(0, 0, i.val)

	case intervalUnitWeek:
		return d.AddDate(0, 0, 7*i.val)

	case intervalUnitMonth:
		return d.AddDate(0, i.val, 0)

	case intervalUnitYear:
		return d.AddDate(i.val, 0, 0)

	default:
		return d
	}
}

func parseInterval(interval string) (Interval, error) {
	var res Interval

	if _, err := fmt.Sscanf(interval, "%d%c", &res.val, &res.unit); err != nil {
		return res, fmt.Errorf("error parsing interval: %w", err)
	}

	return res, nil
}
