/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

var (
	url   string
	count int
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping is a network utility used to test the reachability and measure the response time of a host on an IP network.",
	Long: `The 'ping' command is a network utility used to test the connectivity between your local machine and a specified host.

It sends a series of ICMP (Internet Control Message Protocol) Echo Request packets to the target host and waits for Echo Reply packets to determine if the host is reachable.

This command is useful for diagnosing network connectivity issues, measuring network latency, and checking the overall status of network interfaces.

Examples:
  ping -u google.com
  ping -c 4 -u 192.168.1.1`,

	Run: func(cmd *cobra.Command, args []string) {
		pingHost()
	},
}

func pingHost() {
	addr, err := net.ResolveIPAddr("ip4", url)
	if err != nil {
		fmt.Printf("Could not resolve host: %v\n", err)
		return
	}

	conn, err := icmp.ListenPacket("ipv4:icmp", "0.0.0.0")
	if err != nil {
		fmt.Printf("Error creating ICMP connection: %v\n", err)
		return
	}
	defer conn.Close()

	for i := 0; i < count; i++ {
		start := time.Now()
		sendEchoRequest(i, conn, addr)
		receiveEchoReply(conn, start, addr)
	}
}

func sendEchoRequest(i int, conn *icmp.PacketConn, addr *net.IPAddr) {
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  i,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}
	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		fmt.Printf("Error marshalling ICMP message: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.WriteTo(msgBytes, addr)
	if err != nil {
		fmt.Printf("Error sending ICMP message: %v\n", err)
		os.Exit(1)
	}
}

func receiveEchoReply(conn *icmp.PacketConn, start time.Time, addr *net.IPAddr) {
	reply := make([]byte, 1500)
	err := conn.SetReadDeadline(time.Now().Add(3 * time.Second)) // Timeout after 3 seconds
	if err != nil {
		fmt.Printf("Error setting read deadline: %v\n", err)
		os.Exit(1)
	}

	n, _, err := conn.ReadFrom(reply)
	if err != nil {
		fmt.Printf("Request timed out.\n")
		return
	}

	duration := time.Since(start)
	rm, err := icmp.ParseMessage(1, reply[:n])
	if err != nil {
		fmt.Printf("Error parsing ICMP message: %v\n", err)
		os.Exit(1)
	}

	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		fmt.Printf("Reply from %s: time=%v\n", addr.String(), duration)
	default:
		fmt.Printf("Got %+v from %v; expected echo reply\n", rm, addr)
	}

}

func init() {

	pingCmd.Flags().StringVarP(&url, "url", "u", "", "The URL to ping")
	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err.Error())
	}
	pingCmd.Flags().IntVarP(&count, "count", "c", 5, "Stop after sending the specified number of requests")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
