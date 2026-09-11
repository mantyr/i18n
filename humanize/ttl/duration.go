package ttl

import (
	"errors"
	"time"
)

type Duration struct {
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func (d Duration) Duration() time.Duration {
	return time.Duration(d.Days*24)*time.Hour +
		time.Duration(d.Hours)*time.Hour +
		time.Duration(d.Minutes)*time.Minute +
		time.Duration(d.Seconds)*time.Second
}

func Split(d time.Duration) (Duration, error) {
	if d < 0 {
		return Duration{}, errors.New("expected positive number but actual negative")
	}
	days := int(d / (24 * time.Hour))
	d -= time.Duration(days) * 24 * time.Hour
	hours := int(d / time.Hour)
	d -= time.Duration(hours) * time.Hour
	minutes := int(d / time.Minute)
	d -= time.Duration(minutes) * time.Minute
	seconds := int(d / time.Second)

	return Duration{
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}, nil
}
