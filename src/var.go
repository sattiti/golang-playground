package main

import (
    "fmt"
)

const (
    Sunday    = iota // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)

const (
    _ = iota // use _ to skip 0. blank identifier.
    D        // 1
    E        // 2
    F        // 3
)

func todayis(wday string) {
    fmt.Printf("Today is %s\n", wday)
}

func main() {
    // 小文字はパッケージ内のみ参照可。
    // 大文字はパッケージ外の参照可。(exported name)

    // var
    var today int = 0

    switch today {
    case Monday:
        todayis("Monday")
    case Tuesday:
        todayis("Tuesday")
    case Wednesday:
        todayis("Wednesday")
    case Thursday:
        todayis("Thursday")
    case Friday:
        todayis("Friday")
    case Saturday:
        todayis("Saturday")
    default:
        todayis("Sunday")
    }

    // utf8 input text as a raw literal.
    r := `日本語`
    fmt.Printf("value: %v\n", r)
    fmt.Printf("type:  %T\n", r)

    // _ underscore
    // 戻り値は使用しない場合、_ に代入する。
    _ = &r
}
