package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"encoding/json"
	"net"
	"strconv"
	"strings"
	"time"
)

type StratumPinger struct {
	login   string
	pass    string
	count   int
	ipv6    bool
	host    string
	port    string
	addr    *net.IPAddr
	proto   string
	tls     bool
}

func main() {
	argLogin := flag.String("u", "0x63a14c53f676f34847b5e6179c4f5f5a07f0b1ed", "login")
	argPass := flag.String("p", "x", "pass")
	argCount := flag.Int("c", 5, "stop after <count> replies")
	argV6 := flag.Bool("6", false, "use ipv6")
	argProto := flag.String("t", "stratum2", "stratum type: stratum1, stratum2")
	argTLS := flag.Bool("tls", false, "use TLS")

	flag.Parse()

	argServer := flag.Arg(0)

	if len(argServer) == 0 {
		fmt.Printf("Stratum server cannot be empty\n\n")
		return
	}

	split := strings.Split(argServer, ":")
	if len(split) != 2 {
		fmt.Printf("Invalid host/port specified\n\n")
		return
	}

	if *argCount <= 0 || *argCount > 20000 {
		fmt.Printf("Invalid count specified\n\n")
		return
	}

	portNum, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil || portNum <= 0 || portNum >= 65536 {
		fmt.Printf("Invalid port specified\n\n")
		return
	}

	switch *argProto {
		case "stratum1": 
			fallthrough
		case "stratum2":
			break
		default:
			fmt.Printf("Invalid stratum type specified\n\n")
			return
	}

	pinger := StratumPinger{
		login: *argLogin,
		pass:  *argPass,
		count: *argCount,
		host:  split[0],
		port:  split[1],
		ipv6:  *argV6,
		proto: *argProto,
		tls:   *argTLS,
	}

	if err := pinger.Do(); err != nil {
		fmt.Printf("%s\n\n", err)
	}
}

func (p *StratumPinger) Do() error {
	if err := p.Resolve(); err != nil {
		return err
	}

	creds := ""
	if p.proto == "stratum1" {
		creds = " with credentials: " + p.login + ":" + p.pass
	}
	tls := ""
	if p.tls {
		tls = " TLS"
	}
	fmt.Printf("PING stratum %s (%s)%s port %s%s\n", p.host, p.addr.String(), tls, p.port, creds)

	min := time.Duration(time.Hour)
	max := time.Duration(0)
	avg := time.Duration(0)
	avgCount := 0
	success := 0
	start := time.Now()

	for i := 0; i < p.count; i++ {
		elapsed, err := p.DoPing()
		if err != nil {
			fmt.Printf("%s (%s): seq=%d, %s\n", p.host, p.addr.String(), i, err)
		} else {
			fmt.Printf("%s (%s): seq=%d, time=%s\n", p.host, p.addr.String(), i, elapsed.String())
			if elapsed > max {
				max = elapsed
			}
			if elapsed < min {
				min = elapsed
			}
			avg += elapsed
			avgCount++
			success++
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("\n--- %s ping statistics ---\n", p.host)
	loss := 100 - int64(float64(success) / float64(p.count) * 100.0)
	fmt.Printf("%d packets transmitted, %d received, %d%% packet loss, time %s\n", p.count, success, loss, time.Since(start))
	if success > 0 {
		fmt.Printf("min/avg/max = %s, %s, %s\n", min.String(), (avg / time.Duration(avgCount)).String(), max.String())
	}
	return nil
}

func (p *StratumPinger) Resolve() error {
	var err error
	network := "ip4"

	if p.ipv6 {
		network = "ip6"
	}

	p.addr, err = net.ResolveIPAddr(network, p.host)
	if err != nil {
		return fmt.Errorf("Failed to resolve host name: %s", err)
	}
	return nil
}

func (p *StratumPinger) DoPing() (time.Duration, error) {
	var dial string
	var network string

	if p.ipv6 {
		network = "tcp6"
		dial = "[" + p.addr.IP.String() + "]:" + p.port
	} else {
		network = "tcp4"
		dial = p.addr.IP.String() + ":" + p.port
	}

	var err error
	var conn net.Conn
	if p.tls {
		cfg :=  &tls.Config{InsecureSkipVerify: true}
		conn, err = tls.Dial(network, dial, cfg)
	} else {
		conn, err = net.Dial(network, dial)
	}
	if err != nil {
		return 0, err
	}

	enc := json.NewEncoder(conn)
	buff := bufio.NewReaderSize(conn, 1024)

	readTimeout := 10 * time.Second
	writeTimeout := 10 * time.Second

	conn.SetWriteDeadline(time.Now().Add(writeTimeout))

	var req map[string]interface{}

	switch p.proto {
		case "stratum1":
			req = map[string]interface{}{"id":1, "jsonrpc": "2.0", "method": "eth_submitLogin", "params": []string{p.login,p.pass}}
		case "stratum2":
			req = map[string]interface{}{"id": 1, "method": "mining.subscribe", "params": []string{"stratum-ping/1.0.0", "EthereumStratum/1.0.0"}}
	}

	start := time.Now()
	if err = enc.Encode(&req); err != nil {
		return 0, err
	}
	conn.SetReadDeadline(time.Now().Add(readTimeout))
	if _, _, err = buff.ReadLine(); err != nil {
		return 0, err
	}
	elapsed := time.Since(start)
	conn.Close()

	return elapsed, nil
}
