# stratum-ping

## Abstract

Often the performance of various mining pools is verified using the built-in `ping` utility. While this approach is feasible, it is more important to measure the actual response time from the pool, as it may be busy or using some geolocation-based forwarding solution that brings the connecting endpoint closer to the user but then it still takes some significant time for the actual data to be transferred.

Therefore we find useful a mining-specific approach when we measure the amount of time required to connect and successfully pass authentication using the Stratum protocol. This provides more accurate readings in regards of what the mining software will actually do when mining.

We use this tool at 2Miners internally to measure the performance of our pools. It is capable of pinging through IPv4 and IPv6 with or without TLS.

**TL;DR**: The ping to the pool server's box is not as significant as the actual response time through Stratum.

## Usage

```
Usage of ./stratum-ping:
  -6    use ipv6
  -c int
        stop after <count> replies (default 5)
  -p string
        pass (default "x")
  -t string
        stratum type: stratum1, stratum2 (default "stratum2")
  -tls
        use TLS
  -u string
        login (default "0x63a14c53f676f34847b5e6179c4f5f5a07f0b1ed")

```

## Example Usage

### IPv4 without TLS:
```
# ./stratum-ping eth.2miners.com:2020
PING stratum eth.2miners.com (51.89.64.65) port 2020
eth.2miners.com (51.89.64.65): seq=0, time=14.239754ms
eth.2miners.com (51.89.64.65): seq=1, time=14.318485ms
eth.2miners.com (51.89.64.65): seq=2, time=16.103118ms
eth.2miners.com (51.89.64.65): seq=3, time=15.77519ms
eth.2miners.com (51.89.64.65): seq=4, time=14.223268ms

--- eth.2miners.com ping statistics ---
5 packets transmitted, 5 received, 0% packet loss, time 5.150504495s
min/avg/max = 14.223268ms, 14.931963ms, 16.103118ms
```

### IPv4 using TLS:
```
# ./stratum-ping -tls eth.2miners.com:12020
PING stratum eth.2miners.com (51.195.88.15) TLS port 12020
eth.2miners.com (51.195.88.15): seq=0, time=308.065µs
eth.2miners.com (51.195.88.15): seq=1, time=165.527µs
eth.2miners.com (51.195.88.15): seq=2, time=192.482µs
eth.2miners.com (51.195.88.15): seq=3, time=191.818µs
eth.2miners.com (51.195.88.15): seq=4, time=169.952µs

--- eth.2miners.com ping statistics ---
5 packets transmitted, 5 received, 0% packet loss, time 5.021413961s
min/avg/max = 165.527µs, 205.568µs, 308.065µs
```

### IPv6 without TLS:
```
# ./stratum-ping -6 eth.2miners.com:2020
PING stratum eth.2miners.com (2001:41d0:700:3441::) port 2020
eth.2miners.com (2001:41d0:700:3441::): seq=0, time=176.611µs
eth.2miners.com (2001:41d0:700:3441::): seq=1, time=177.769µs
eth.2miners.com (2001:41d0:700:3441::): seq=2, time=185.733µs
eth.2miners.com (2001:41d0:700:3441::): seq=3, time=166.971µs
eth.2miners.com (2001:41d0:700:3441::): seq=4, time=171.775µs

--- eth.2miners.com ping statistics ---
5 packets transmitted, 5 received, 0% packet loss, time 5.002555972s
min/avg/max = 166.971µs, 175.771µs, 185.733µs
```

## Copyright

Copyright ©2021, [2Miners.com](https://2miners.com). All rights reserved.
