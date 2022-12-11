package main

import (
    "fmt"
    "log"
    "os/exec"
    "strings"

    "github.com/OJ/gobuster"
    "github.com/Ullaakut/nmap"
    "github.com/spf13/cobra"
)

func main() {
    var target string
    var threads int

    rootCmd := &cobra.Command{
        Use:   "port-scanner",
        Short: "Scan a target for open ports",
        Run: func(cmd *cobra.Command, args []string) {
            gb := gobuster.New(target)
            gb.Wordlist = "/usr/share/wordlists/dirbuster/common.txt"
            gb.Threads = threads

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

            if out, err := exec.Command("zap-cli", "--zap-url", "http://localhost:8080", "status").Output(); err != nil {
                log.Fatal(err)
            } else {
                fmt.Printf("OWASP ZAP status: %s\n", out)
            }

            if out, err := exec.Command("openvas-check-setup").Output();
