package main

import (
    "fmt"
    "log"

    "github.com/OJ/gobuster"
)

func main() {
    gb := gobuster.New("https://www.example.com")
    gb.Wordlist = "/usr/share/wordlists/dirbuster/common.txt"
    gb.Threads = 20

    if err := gb.Start(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Open ports: %v\n", gb.OpenPorts())
}
