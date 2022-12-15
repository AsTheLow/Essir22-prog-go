package main

import (
        "fmt"
        "os"

        "github.com/spf13/cobra"
        "github.com/OJ/gobuster"
        "github.com/OJ/gobuster/v3/libgobuster"
        "github.com/OJ/gobuster/v3/libgobuster/plugins/common"
)

var (
        target string
        ports  string
        workers int
        quiet  bool
        help   bool
)

func init() {
        cobra.OnInitialize(initConfig)
        rootCmd.PersistentFlags().StringVarP(&target, "target", "t", "", "Target host or IP address")
        rootCmd.PersistentFlags().StringVarP(&ports, "ports", "p", "", "Comma-separated list of ports to scan")
        rootCmd.PersistentFlags().IntVarP(&workers, "workers", "w", 10, "Number of workers to use for port scanning")
        rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Only show open ports")
        rootCmd.PersistentFlags().BoolVarP(&help, "help", "h", false, "Show help message")
}

func initConfig() {
        // Initialize configuration here if necessary
}

var rootCmd = &cobra.Command{
        Use:   "scanner",
        Short: "Scan target host for open ports and web server directories and files",
        Long:  `Scan target host for open ports and web server directories and files`,
        Run: func(cmd *cobra.Command, args []string) {
                // Check if target host was specified
                if target == "" {
                        fmt.Println("Error: Target host not specified")
                        os.Exit(1)
                }

                // Check if ports were specified
                if ports == "" {
                        // If not, scan all ports
                        ports = "1-65535"
                }

                // Scan target host for open ports
                fmt.Println("Scanning target host for open ports...")
                nmapOutput, err := libgobuster.Nmap(target, ports, workers)
                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }

                // Parse nmap output and print open ports
                openPorts := common.ParseNmap(nmapOutput)
                if len(openPorts) == 0 {
                        fmt.Println("No open ports found")
                        os.Exit(0)
                }

                fmt.Println("Open ports:")
                for _, port := range openPorts {
                        fmt.Println(port)
                }

                // Enumerate directories and files on web server, if port 80 or 443 is open
                if common.StringInSlice("80", openPorts) || common.StringInSlice("443", openPorts) {
                        fmt.Println("\nEnumerating directories and files on web server...")

                        // Create
