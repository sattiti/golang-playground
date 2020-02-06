package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <hello.c>

void cprint(char *s){
  printf(s);
}
*/
import (
  "C"
)

func main() {
  C.hello()
  s := C.CString("hello world again.")
  defer C.free(s)
  C.cprint(s)
}
