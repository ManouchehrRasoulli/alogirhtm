package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	MaxTTL       = 30
	Timeout      = time.Second * 3
	ProtocolICMP = 1
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <host>\n", os.Args[0])
		os.Exit(1)
	}
	dest := os.Args[1]

	ipAddr, err := net.ResolveIPAddr("ip4", dest)
	if err != nil {
		log.Fatal("Failed to resolve IP:", err)
	}

	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatal("Failed to listen for ICMP packets:", err)
	}
	defer conn.Close()

	for ttl := 1; ttl <= MaxTTL; ttl++ {
		p := ipv4.NewPacketConn(nil)
		p.SetTTL(ttl)

		wm := icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Seq:  ttl,
				Data: []byte("HELLO"),
			},
		}

		wb, err := wm.Marshal(nil)
		if err != nil {
			log.Fatal("Marshal error:", err)
		}

		start := time.Now()
		_, err = conn.WriteTo(wb, &net.IPAddr{IP: ipAddr.IP})
		if err != nil {
			log.Printf("%2d  *\n", ttl)
			continue
		}

		conn.SetReadDeadline(time.Now().Add(Timeout))
		rb := make([]byte, 1500)
		n, peer, err := conn.ReadFrom(rb)
		if err != nil {
			fmt.Printf("%2d  * (timeout)\n", ttl)
			continue
		}

		rm, err := icmp.ParseMessage(ProtocolICMP, rb[:n])
		if err != nil {
			log.Println("Failed to parse ICMP message:", err)
			continue
		}

		duration := time.Since(start)

		switch rm.Type {
		case ipv4.ICMPTypeTimeExceeded:
			fmt.Printf("%2d  %s  %v\n", ttl, peer, duration)
		case ipv4.ICMPTypeEchoReply:
			fmt.Printf("%2d  %s  %v (destination reached)\n", ttl, peer, duration)
			return
		default:
			fmt.Printf("%2d  %s  %v (unknown ICMP type: %v)\n", ttl, peer, duration, rm.Type)
		}
	}
}
