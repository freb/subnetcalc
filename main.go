package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/brotherpowers/ipsubnet"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("requires at least one subnet argument")
		os.Exit(1)
	}
	total := 0
	for _, netaddr := range os.Args[1:] {
		fields := strings.Split(netaddr, "/")
		if len(fields) != 2 {
			fmt.Println("malformed CIDR address:", netaddr)
			continue
		}
		ip := fields[0]
		size, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Printf("invalid subnet size for %s: %s\n", netaddr, err)
			continue
		}
		sub := ipsubnet.SubnetCalculator(ip, size)
		fmt.Println(netaddr)
		fmt.Println("------------------")
		fmt.Println("Total IPs:        ", sub.GetNumberIPAddresses())
		fmt.Println("usage IP Range:   ", sub.GetIPAddressRange()[0], "-", sub.GetIPAddressRange()[1])
		fmt.Println("Network Address:  ", sub.GetNetworkPortion())
		fmt.Println("Broadcast Address:", sub.GetBroadcastAddress())
		fmt.Println("Subnet Mask:      ", sub.GetSubnetMask())
		fmt.Println()
		total += sub.GetNumberIPAddresses()
	}
	fmt.Println("Total: ", total)
}
