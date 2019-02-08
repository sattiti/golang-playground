package main

import (
    "bufio.Scanner"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

// io
// すべての io の基本 interface
// io.Reader
// io.Writer

// 標準入出力、エラー
// os.Stdin
// os.Stdout
// os.Stderr

// メモリに書き込む
// bytes.Buffer(struct)

// メモリから読込む
// bytes.Reader(struct)

// 1行ずつ読み込む
// bufio.NewScanner

// すべてを読み込んで、[]byte を作成する。文字列にしたい場合、キャストする。string(b)
// io/ioutil.ReadAll

// ファイルからすべてを読み込む。
// io/ioutil.ReadFile

// すべてをファイルに書込む。
// io/ioutil.WriteFile

// io.Reader から io.Writer にデータをコピー。
// io.Copy

func readfile(path string) {
    contents, err := ioutil.readFile(path)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", contents)
}

func readlines(path string) {
    f, err := os.OpenFile("FILE", os.O_RDONLY, 0775)
    if err != nil {
        log.Fatal(err)
    }

    s := bufio.NewScanner(f)
    if s.Err() != nil {
        log.Fatal(s.Err())
    }

    for s.Scan() {
        fmt.Println(s.Text())
    }
}

func main() {
    filePath := "./res/index.html"
    readfile(filePath)
    readlines(filePath)
}
