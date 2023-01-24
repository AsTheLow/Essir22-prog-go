package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var target string
var ports string
var workers int
var quiet bool

var rootCmd = &cobra.Command{
	Use:   "portscan",
	Short: "A simple port scanner",
	Long:  `A simple port scanner that can scan a range of ports on a target IP`,
	Run: func(cmd *cobra.Command, args []string) {
		scanPorts(target, ports, workers, quiet)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&target, "target", "t", "", "the target IP to scan")
	rootCmd.MarkFlagRequired("target")

	rootCmd.Flags().StringVarP(&ports, "ports", "p", "", "the range of ports to scan (examples: 1024-65535, all)")
	rootCmd.MarkFlagRequired("ports")

	rootCmd.Flags().IntVarP(&workers, "workers", "w", 1, "the number of workers to use for scanning in parallel")

	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "don't log, only show results")
}

func main() {
	rootCmd.Execute()
}

func scanPorts(target string, ports string, workers int, quiet bool) {
	portsList := parsePorts(ports)
	var wg sync.WaitGroup
	wg.Add(len(portsList))

	for _, port := range portsList {
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", target, port)
			conn, err := net.DialTimeout("tcp", address, time.Second*5)
			if err != nil {
				if !quiet {
					fmt.Printf("Port %d is closed\n", port)
				}
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open\n", port)
		}(port)
	}
	wg.Wait()
}
func parsePorts(ports string) []int {
	var portsList []int
	if ports == "all" {
		for i := 1; i <= 65535; i++ {
			portsList = append(portsList, i)
		}
	} else {
		parts := strings.Split(ports, "-")
		if len(parts) != 2 {
			log.Fatalf("Invalid port range: %s", ports)
		}
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Invalid start port: %s", parts[0])
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Invalid end port: %s", parts[1])
		}
		if end < start {
			log.Fatalf("Invalid port range: %s", ports)
		}
		for i := start; i <= end; i++ {
			portsList = append(portsList, i)
		}
	}
	return portsList
}
