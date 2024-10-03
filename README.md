# Port Scanner

This is a simple concurrent port scanner written in Go that scans all ports in a specified range (default: 1-65535) for a given host using TCP or any other specified protocol. It uses goroutines for concurrency to improve scanning speed and efficiency.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Flags](#flags)
- [Example](#example)
- [How it Works](#how-it-works)
- [Contributing](#contributing)
- [License](#license)

## Features

- Scans all ports (1 to 65535) on a specified host.
- Supports custom protocols (default: `TCP`).
- Concurrent scanning with adjustable worker pools.
- Configurable timeout for connection attempts.
- Outputs open ports with their corresponding protocols.

## Installation

To use this tool, you need to have Go installed on your machine. If you don't have Go installed, you can download and install it from [the official website](https://golang.org/dl/).

Once Go is installed, clone this repository and navigate to the project directory:

```bash
git clone https://github.com/ashissharma97/port-scanner.git
cd port-scanner
```

You can build the tool using:

```bash
go build -o portscanner
```

## Usage

After building the program, you can run it with various flags to customize the host, protocol, worker count, and request timeout.

```bash
./portscanner [flags]
```

## Flags

The following flags can be passed to configure the scanner:

- `-host`: Specifies the destination host to scan. Default: `google.com`.
- `-protocol`: Specifies the protocol to use for scanning (e.g., `tcp`, `udp`). Default: `tcp`.
- `-requestTimeout`: Time in seconds to wait for a response from the port. Default: `2`.
- `-maxWorkers`: The number of concurrent workers (goroutines) used to perform the scan. Default: `100`.

## Example

To scan ports on `example.com` using the `tcp` protocol, with a timeout of 3 seconds and 50 workers, use the following command:

```bash
./portscanner -host example.com -protocol tcp -requestTimeout 3 -maxWorkers 50
```

Sample output:
```
Host:  example.com
Protocol:  tcp
Max Workers:  50
Timeout:  3
Port is open:  80 Protocol:  tcp
Port is open:  443 Protocol:  tcp
```

## How it Works

1. The program initializes flags for the host, protocol, number of workers, and request timeout.
2. A worker pool is created with the specified number of workers (goroutines), each performing port scans.
3. The ports to be scanned (1–65535) are fed to workers via a channel (`portsChan`).
4. For each port, a connection attempt is made using the specified protocol and timeout.
5. If a port is open, the result is passed to the `resultsChan`, and open ports are printed to the console.
6. The program finishes once all ports are scanned.

## Contributing

If you'd like to contribute, feel free to fork the repository, make changes, and submit a pull request. Any improvements, bug fixes, or suggestions are welcome!

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

Enjoy scanning your ports responsibly! 😊