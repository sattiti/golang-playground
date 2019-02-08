package main

import (
    "crypto/hmac"
    "crypto/md5"
    "crypto/sha256"
    "crypto/sha512"
    "encoding/base64"
    "fmt"
    "io"
)

func main() {
    in := "hello"
    salt := "world"

    md5 := md5.New()
    io.WriteString(md5, in)

    sha_256 := sha256.New()
    sha_256.Write([]byte(in))

    sha_512 := sha512.New()
    sha_512.Write([]byte(in))

    sha_512_256 := sha512.Sum512_256([]byte(in))

    hmac512 := hmac.New(sha512.New, []byte(salt))
    hmac512.Write([]byte(in))

    fmt.Printf("md5:        %x\n", md5.Sum(nil))
    fmt.Printf("sha256:     %x\n", sha_256.Sum([]byte(salt)))
    fmt.Printf("sha512:     %x\n", sha_512.Sum([]byte(salt)))
    fmt.Printf("sha512_256: %x\n", sha_512_256)
    fmt.Printf("hmac512:    %s\n", base64.StdEncoding.EncodeToString(hmac512.Sum(nil)))
}
