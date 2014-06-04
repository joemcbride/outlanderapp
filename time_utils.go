package main

import (
    "errors"
    "time"
)

var ISO8601Full string = "2006-01-02T15:04:05-07:00"

func ISO8601FullFromNow () (string, error) {
    return ISO8601FullFromTime(time.Now())
}

func ISO8601FullFromTime (t time.Time) (string, error) {
    i := t.Format(ISO8601Full)
    if i == "" {
        return "", errors.New("Unable to format as: "+ISO8601Full)
    }
    return i, nil
}