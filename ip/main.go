package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arg := firstArg(os.Args[1:])
	if arg == "" {
		fmt.Println("no ip address")
		return
	}
	ip := parseIP(arg)
	fmt.Println(ip)
}

func firstArg(args []string) (arg string) {
	if len(args) > 0 {
		arg = args[0]
	}
	return
}

func parseIP(ipStr string) string {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "invalid ip address"
	}
	return ip.String()
}
