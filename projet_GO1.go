package main

import (
    "fmt"
    "log"
    "strings"

    "github.com/OJ/gobuster"
    "github.com/Ullaakut/nmap"
    "github.com/spf13/cobra"
)

func main() {
    var target string
    var wordlist string
    var threads int

    rootCmd := &cobra.Command{
        Use:   "port-scanner",
        Short: "Scan a target for open ports",
        Run: func(cmd *cobra.Command, args []string) {
            gb := gobuster.New(target)
            gb.Wordlist = wordlist
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
        },
    }

    rootCmd.Flags().StringVarP(&target, "target", "t", "", "The target URL to scan")
    rootCmd.Flags().StringVarP(&wordlist, "wordlist", "w", "/usr/share/wordlists/dirbuster/common.txt", "The path to the wordlist to use")
    rootCmd.Flags().IntVarP(&threads, "threads", "T", 20, "The number of worker threads to use")

    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err
