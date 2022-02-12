You might need to do a "go mod init bhg-scanner" or "go mod tidy"

### Things I Added

- The `out.csv` file is the captured output from running `main.go`
- There are two extra tests, one for specific port targets and the other for closed ports.
- I added a struct (`Ports`) to make storage easier.
- An example of usage can be found on line 7 of `scanner.go`
- `main.go` will print the amount of time `PortScanner` took to execute.
- A target address and ports can now be specified.
- When timed, `PortScanner` took ~11s.