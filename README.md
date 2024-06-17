# Protochecker

`protochecker` is a Go-based tool that checks if a list of domains is running on HTTP or HTTPS and saves the results to an output file. It uses concurrent processing to efficiently handle large lists of domains.

## Features

- Checks if domains are running on HTTP or HTTPS.
- Saves results to an output file.
- Concurrent processing for fast execution.

## Requirements

- Go 1.16 or higher

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/bilautagang/protochecker.git
    cd protochecker
    ```

2. Build the Go program:
    ```bash
    go build check_protocol.go
    ```

## Usage

Run the program with the input file containing the list of domains and an optional output file name:

```bash
./check_protocol <input_file> [output_file]
