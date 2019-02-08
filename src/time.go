package main

import (
    "fmt"
    "time"
)

func main() {
    // time
    t := time.Now()
    y, m, d := t.Date()
    fmt.Printf("Nanosecond:  %v\n", time.Nanosecond)
    fmt.Printf("Microsecond: %v\n", time.Microsecond)
    fmt.Printf("Millisecond: %v\n", time.Millisecond)
    fmt.Printf("Second:      %v\n", time.Second)
    fmt.Printf("Minute:      %v\n", time.Minute)
    fmt.Printf("Hour:        %v\n", time.Hour)

    fmt.Println("Sleep 2 second...")
    time.Sleep(time.Second * 2)
    fmt.Println("Slept 2 second...")

    fmt.Printf("Date:        %v %v %v\n", y, m, d)
    fmt.Printf("Year:        %v\n", t.Year())
    fmt.Printf("YearDay:     %v\n", t.YearDay())
    fmt.Printf("Month:       %v\n", t.Month())
    fmt.Printf("Day:         %v\n", t.Day())
    fmt.Printf("Hour:        %v\n", t.Hour())
    fmt.Printf("Minute:      %v\n", t.Minute())
    fmt.Printf("Second:      %v\n", t.Second())
    fmt.Printf("Nanosecond:  %v\n", t.Nanosecond())
    fmt.Printf("Weekday:     %v\n", t.Weekday())
    fmt.Printf("Local:       %v\n", t.Local())
    fmt.Printf("Location:    %v\n", t.Location())
    fmt.Printf("UTC:         %v\n", t.UTC())
    fmt.Printf("Unix:        %v\n", t.Unix())
    fmt.Printf("UnixNano:    %v\n", t.UnixNano())
    fmt.Printf("String:      %v\n", t.String())
    zName, zOffset := t.Zone()
    fmt.Printf("Zone:        %v-%v\n", zName, zOffset)
}
