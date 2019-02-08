package main

import(
    "fmt"
)

func main(){
    //var ary = []int{0, 1, 2, 3, 4, 5, 6}
    ary := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

    for i := 0; i < len(ary); i++{
        fmt.Printf("ary[%d] == %s\n", i, ary[i])
    }

    fmt.Println("----")

    fmt.Printf("ary[0:2] == %s\n", ary[0:2])
    fmt.Printf("ary[2:] == %s\n", ary[2:])
    fmt.Printf("ary[:2] == %s\n", ary[:2])

    fmt.Println("----")
    fmt.Printf("len(ary) == %d\n", len(ary))
    
    // cap は容量
    fmt.Printf("cap(ary) == %d\n", cap(ary))


    // make 関数で生成する。
    // make(型, [長さ], [容量])
    ary2 := make([]string, 7)

    fmt.Println("----------")
    fmt.Printf("ary2 == %#v\n", ary2)
}