package main

import (
  "syscall"
  "unsafe"
)

func run(f func()) {
  fmt.Printf("----\n")
  f()
  fmt.Printf("----\n\n")
}

func throwError(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  run(func() {
    dll, rrr := syscall.LoadDLL("user32.dll")
    throwError(err)
    defer dll.Release()

    proc, err := dll.FindProc("MessageBoxW")
    throwError(err)

    proc.Call(
      0,
      uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("ハローworld"))),
      uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("MessageBoxTitle"))),
      0,
    )
  })
}
