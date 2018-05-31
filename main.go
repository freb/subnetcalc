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
	utotal := 0 //usable
	for _, netaddr := range os.Args[1:] {

		// testing. I should switch to the net.ParseCIDR parser.
		// then use ipnet.Mask.Size() and use the second return, bits. guess i still
		// need to convert to a string first...
		//
		// /32 doesn't display properly, but in ipsubnet lib.

		// aip, aipnet, _ := net.ParseCIDR(netaddr)
		// fmt.Println(aip, aipnet)
		// aones, abits := aipnet.Mask.Size()
		// fmt.Println(aones, abits)

		// end testing

		fields := strings.Split(netaddr, "/")
		if len(fields) == 1 {
			fields = append(fields, "32")
		} else if len(fields) > 2 {
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

		if size == 32 {
			fmt.Println(netaddr)
			fmt.Println("------------------")
			fmt.Println("Total IPs:        ", 1)
			fmt.Println("Usable IPs:       ", 1)
			fmt.Println("usage IP Range:   ", sub.GetIPAddressRange()[0], "-", sub.GetIPAddressRange()[0])
			fmt.Println("Network Address:  ", sub.GetNetworkPortion())
			fmt.Println("Broadcast Address:", sub.GetNetworkPortion())
			fmt.Println("Subnet Mask:      ", sub.GetSubnetMask())
			fmt.Println()
			total += 1
			utotal += 1
			continue
		}

		fmt.Println(netaddr)
		fmt.Println("------------------")
		fmt.Println("Total IPs:        ", sub.GetNumberIPAddresses())
		fmt.Println("Usable IPs:       ", sub.GetNumberAddressableHosts())
		fmt.Println("usage IP Range:   ", sub.GetIPAddressRange()[0], "-", sub.GetIPAddressRange()[1])
		fmt.Println("Network Address:  ", sub.GetNetworkPortion())
		fmt.Println("Broadcast Address:", sub.GetBroadcastAddress())
		fmt.Println("Subnet Mask:      ", sub.GetSubnetMask())
		fmt.Println()
		total += sub.GetNumberIPAddresses()
		utotal += sub.GetNumberAddressableHosts()
	}
	fmt.Println("Total:        ", total)
	fmt.Println("Usable Total: ", utotal)
}
