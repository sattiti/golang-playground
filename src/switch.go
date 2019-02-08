package main

import(
    "fmt"
    "runtime"
)

func main(){
    var rt string

    // auto break.
    // break せずに通したい場合は、case の最後に fallthrough 文を書く。
    switch runtime.GOOS{
        case "darwin":
            rt = "OSX"
            fallthrough
            
        case "linux":
            rt = "Linux"
        default:
            rt = os
    }

    fmt.Println(rt)
}
