package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	udplistener, err := net.ListenUDP("udp", &net.UDPAddr{Port: 9310})
	if err != nil {
		log.Panic(err)
	}
	defer udplistener.Close()
	fmt.Println("server open")

	peers := make([]*net.UDPAddr, 2, 2)
	buf := make([]byte, 256)

	n, addr, err := udplistener.ReadFromUDP(buf)
	if err != nil {
		log.Panic("Failed to readfromudp")
	}
	peers[0] = addr
	fmt.Println("read %d size, from %s, msg:%s\n", n, addr.String(), buf[:n])

	n, addr, err = udplistener.ReadFromUDP(buf)
	if err != nil {
		log.Panic("Failed to readfromudp")
	}
	peers[1] = addr
	fmt.Println("read %d size, from %s, msg:%s\n", n, addr.String(), buf[:n])

	udplistener.WriteToUDP([]byte(peers[0].String()), peers[1])
	udplistener.WriteToUDP([]byte(peers[1].String()), peers[0])

	fmt.Println("server closing in 10 seconds")
	time.Sleep(time.Second * 10)
}
