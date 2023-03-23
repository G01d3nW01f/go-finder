package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run search.go <string>")
        os.Exit(1)
    }

    root := "." // 検索を開始するディレクトリ
    target := os.Args[1]
    exclude := os.Args[0] // 自分自身のファイルを除外する

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && path != exclude {
            lines, err := searchInFile(path, target)
            if err != nil {
                return err
            }
            for _, line := range lines {
                fmt.Printf("%s:%s\n", path, line)
            }
        }
        return nil
    })
    if err != nil {
        fmt.Println(err)
    }
}

func searchInFile(file string, target string) ([]string, error) {
    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    var lines []string
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, target) {
            lines = append(lines, line)
        }
    }
    if err := scanner.Err(); err != nil {
        if err == io.EOF {
            return lines, nil
        }
        return nil, err
    }

    return lines, nil
}
