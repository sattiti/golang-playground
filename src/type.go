package main

import (
    "fmt"
    "reflect"
    "math/cmplx"
)

func showType(typeName string, varName interface{}) {
    fmt.Printf("= %s\n", typeName)
    fmt.Println("----")
    fmt.Printf("value: %V\n", varName)
    fmt.Printf("type:  %v\n", reflect.TypeOf(varName))
    // fmt.Printf("type:  %T\n", varName)
    fmt.Println("----\n")
}

func main() {
    // data types
    fmt.Println("\n-- DataTypes --")

    // interface any type
    var t interface{}
    showType("interface", t)

    // bool
    t = true
    showType("bool", t)

    // string
    t = "hello world"
    showType("string", t)

    // nil
    t = nil
    showType("nil", t)

    // int
    t = 10
    showType("int", t)

    // int8 -128..127
    t = int8(127)
    showType("int8", t)

    // int16 -32768..32767
    t = int16(32767)
    showType("int16", t)

    // int32 -2147483648..2147483647
    // alias: rune (unicode code)
    t = int32(2147483647)
    showType("int32/rune", t)

    // int64 -9223372036854775808..9223372036854775807
    t = int64(9223372036854775807)
    showType("int64", t)

    // uint
    t = uint(10)
    showType("uint", t)

    // uint8 0..255
    // alias: byte
    t = byte(255)
    showType("byte/uint8", t)

    // uint16 0..65535
    t = uint16(65535)
    showType("uint16", t)

    // uint32
    t = uint32(4294967295)
    showType("uint32", t)

    // uint64 0..18446744073709551615
    t = uint64(18446744073709551615)
    showType("uint64", t)

    // uintptr uint32/uint64 (pointer
    t = uintptr(123)
    showType("uintptr", t)

    // float32 IEEE-754
    t = float32(184.46744073)
    showType("float32", t)

    // float64 IEEE-754
    t = float64(184.46744073)
    showType("float64", t)

    // complex64 実数部・虚数部を float32 で表現する複素数型
    t = complex64(cmplx.Sqrt(-5 + 12i))
    showType("complex64", t)
    
    // complex128 実数部・虚数部を float64 で表現する複素数型
    t = cmplx.Sqrt(-5 + 12i)
    showType("complex128", t)
    
    // Array types
    t = make([]string, 10)
    t = [10]string{}
    showType("array", t)

    // Slice types
    t = [4]int{1, 2, 3, 4}
    showType("slice", t)

    // Struct types
    t = struct{}{}
    showType("struct", t)

    // Pointer types
    t = &struct{}{}
    showType("struct pointer", t)

    // Function types
    t = func() {}
    showType("func", t)

    // Map types
    t = make(map[string]string)
    t = map[string]string{}
    showType("map", t)

    // Channel types
    t = make(chan int, 0)
    showType("channel", t)

    // type cast
    var ci int = 40
    fmt.Printf("cast: %f\n", float64(ci))
}
