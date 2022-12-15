package main

import (
    "fmt"
    "log"
    "os"

    "github.com/spf13/cobra"

    "github.com/OJ/gobuster"
    "github.com/OJ/gobuster/v3/libgobuster"
)

var (
    rootCmd = &cobra.Command{
        Use:   "my-cli",
        Short: "My command-line interface",
        Long:  "A sample CLI application using gobuster and spf13/cobra",
        Run: func(cmd *cobra.Command, args []string) {
            // Print usage information if no arguments are provided
            if len(args) == 0 {
                cmd.Usage()
                return
            }

            // Get the target URL and wordlist from the arguments
            url := args[0]
            wordlist := args[1]

            // Enumerate directories and files on the web server using gobuster
            enumerateWebServer(url, wordlist)
        },
    }
)

func main() {
    // Define flags for the root command
    rootCmd.Flags().StringP("target", "t", "", "target host or IP address")
    rootCmd.Flags().StringSliceP("ports", "p", []string{}, "target port(s)")
    rootCmd.Flags().IntP("workers", "w", 10, "number of concurrent workers")
    rootCmd.Flags().BoolP("quiet", "q", false, "don't show banner and logo")

    // Parse command-line arguments
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func enumerateWebServer(url string, wordlist string) {
    // Create a new gobuster instance
    gb := gobuster.New(url, wordlist, &libgobuster.Options{
        Threads:  10,
        Quiet:    true,
        Verbose:  false,
        NoColor:  false,
        ShowIPs:  false,
        Recursive: false,
        WildcardForced: false,
    })

    // Enumerate directories and files on the web server
    err := gb.Start()
    if err != nil {
        log.Fatal(err)
    }

    // Print the results
    for _, result := range gb.Results() {
        fmt.Println(result.String())
    }
}
