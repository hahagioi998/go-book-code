package main

import (
	"fmt"
	"net"
)

type DNSError struct {
	//...
}

func (e *DNSError) Error() string {
	//...
	return "error"
}
func (e *DNSError) Timeout() bool {
	//...
	return true
}
func (e *DNSError) Temporary() bool {
	//...
	return true
}

func main() {
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}
