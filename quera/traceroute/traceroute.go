package main

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s host\n", os.Args[0])
		return
	}

	host := os.Args[1]
	hostIp, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		log.Fatalf("Failed to resolve host IP address: %s\n", err)
		return
	}

	log.Printf("host IP: %s\n", hostIp.String())

	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatalf("Failed to create ICMP listener: %s\n", err)
		return
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatalf("Failed to close ICMP listener: %s\n", err)
		}
	}()

	// max ttl is 30 hops
	for ttl := 1; ttl <= 30; ttl++ {
		packet := ipv4.NewPacketConn(conn.IPv4PacketConn().PacketConn)
		_ = packet.SetTTL(ttl)

		webMessage := icmp.Message{
			Type:     ipv4.ICMPTypeEcho,
			Code:     0,
			Checksum: 0,
			Body: &icmp.Echo{
				ID:   os.Getppid() & 0xffff,
				Seq:  ttl,
				Data: []byte("HELLO"),
			},
		}

		webByte, err := webMessage.Marshal(nil)
		if err != nil {
			log.Fatalf("Failed to marshal web message: %s\n", err)
			return
		}

		start := time.Now()
		_, err = conn.WriteTo(webByte, &net.IPAddr{IP: hostIp.IP})
		if err != nil {
			log.Printf("%2d  *\n", ttl)
			continue
		}

		_ = conn.SetReadDeadline(time.Now().Add(3 * time.Second))
		rb := make([]byte, 1500)
		n, peer, err := conn.ReadFrom(rb)
		if err != nil {
			fmt.Printf("%2d  * (timeout)\n", ttl)
			continue
		}

		rm, err := icmp.ParseMessage(1, rb[:n])
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
