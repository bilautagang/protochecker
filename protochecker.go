package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
    "sync"
    "time"
)

// checkProtocol checks if a domain is running on HTTP or HTTPS and writes the result to the output file.
func checkProtocol(domain string, outputFile string, wg *sync.WaitGroup, mu *sync.Mutex) {
    defer wg.Done()

    client := &http.Client{
        Timeout: 5 * time.Second,
    }

    var result string

    // Check HTTPS
    resp, err := client.Head("https://" + domain)
    if err == nil && resp.StatusCode < 400 {
        result = fmt.Sprintf("HTTPS -- %s", domain)
        mu.Lock()
        fmt.Println(result)
        f, _ := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0600)
        defer f.Close()
        fmt.Fprintf(f, "%s,HTTPS\n", domain)
        mu.Unlock()
        return
    }

    // Check HTTP
    resp, err = client.Head("http://" + domain)
    if err == nil && resp.StatusCode < 400 {
        result = fmt.Sprintf("HTTP -- %s", domain)
        mu.Lock()
        fmt.Println(result)
        f, _ := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0600)
        defer f.Close()
        fmt.Fprintf(f, "%s,HTTP\n", domain)
        mu.Unlock()
        return
    }

    // Not Reachable
    result = fmt.Sprintf("Not Reachable -- %s", domain)
    mu.Lock()
    fmt.Println(result)
    f, _ := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0600)
    defer f.Close()
    fmt.Fprintf(f, "%s,Not Reachable\n", domain)
    mu.Unlock()
}

func main() {
    // Print developer information
    fmt.Println("Developed by BilautaGang")

    // Check if the input file was provided
    if len(os.Args) < 2 {
        fmt.Println("Usage: check_protocol <input_file> [output_file]")
        os.Exit(1)
    }

    // Set input and output file names
    inputFile := os.Args[1]
    outputFile := "protocol_check_results.txt"
    if len(os.Args) == 3 {
        outputFile = os.Args[2]
    }

    // Check if the input file exists
    file, err := os.Open(inputFile)
    if err != nil {
        fmt.Printf("Error: Input file '%s' not found.\n", inputFile)
        os.Exit(1)
    }
    defer file.Close()

    // Create or clear the output file
    outFile, err := os.Create(outputFile)
    if err != nil {
        fmt.Printf("Error: Could not create output file '%s'.\n", outputFile)
        os.Exit(1)
    }
    outFile.Close()

    var wg sync.WaitGroup
    var mu sync.Mutex
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        domain := scanner.Text()
        wg.Add(1)
        go checkProtocol(domain, outputFile, &wg, &mu)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input file:", err)
    }

    wg.Wait()
    fmt.Printf("Protocol check results have been saved to %s.\n", outputFile)
}
