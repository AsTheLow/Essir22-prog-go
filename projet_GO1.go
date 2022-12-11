package main

import (
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/OJ/gobuster"
    "github.com/Ullaakut/nmap"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Usage: go run main.go <TARGET_URL>")
    }

    target := os.Args[1]

    gb := gobuster.New(target)
    gb.Wordlist = "/usr/share/wordlists/dirbuster/common.txt"
    gb.Threads = 20

    if err := gb.Start(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Open ports (gobuster): %v\n", gb.OpenPorts())

    scanner, err := nmap.NewScanner(
        nmap.WithTargets(strings.Split(target, ",")...),
        nmap.WithPorts("1-65535"),
        nmap.WithServiceInfo(),
    )
    if err != nil {
        log.Fatal(err)
    }

    result, err := scanner.Run()
    if err != nil {
        log.Fatal(err)
    }

    openPorts := []string{}
    for _, host := range result.Hosts {
        if len(host.Ports) == 0 || len(host.Addresses) == 0 {
            continue
        }

        address := host.Addresses[0]
        for _, port := range host.Ports {
            if port.State.State == "open" {
                openPorts = append(openPorts, fmt.Sprintf("%s:%d/%s", address, port.ID, port.Service.Name))
            }
        }
    }

    fmt.Printf("Open ports (nmap): %v\n", openPorts)
}
