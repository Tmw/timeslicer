package main

import (
	"testing"
	"time"
)

func TestParseInterval(t *testing.T) {
	tests := map[string]Interval{
		"1m":  {unit: intervalUnitMinute, val: 1},
		"2M":  {unit: intervalUnitMonth, val: 2},
		"7d":  {unit: intervalUnitDay, val: 7},
		"2y":  {unit: intervalUnitYear, val: 2},
		"4w":  {unit: intervalUnitWeek, val: 4},
		"13h": {unit: intervalUnitHour, val: 13},
	}

	for input, want := range tests {
		t.Run("parsing "+input, func(t *testing.T) {
			got, err := parseInterval(input)
			if err != nil {
				t.Error("expected no error, got error:", err)
			}

			if got.unit != want.unit {
				t.Errorf("unit mismatch. Expected: %c received %c", want.unit, got.unit)
			}

			if got.val != want.val {
				t.Errorf("val mismatch. Expected: %d received %d", want.val, got.val)
			}
		})
	}
}

func TestApply(t *testing.T) {
	start := time.Now().UTC()

	tests := map[Interval]time.Time{
		{unit: intervalUnitDay, val: 1}:     start.AddDate(0, 0, 1),
		{unit: intervalUnitDay, val: 7}:     start.AddDate(0, 0, 7),
		{unit: intervalUnitMonth, val: 3}:   start.AddDate(0, 3, 0),
		{unit: intervalUnitYear, val: 5}:    start.AddDate(5, 0, 0),
		{unit: intervalUnitHour, val: 6}:    start.Add(time.Hour * 6),
		{unit: intervalUnitMinute, val: 30}: start.Add(time.Minute * 30),
	}

	for interval, want := range tests {
		t.Run("apply "+interval.String(), func(t *testing.T) {
			got := interval.Apply(start)
			if got != want {
				t.Errorf("incorrect apply result. expected %+v but got %+v", want, got)
			}
		})
	}
}
