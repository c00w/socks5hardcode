package main

import (
	"log"
	"net"
    "io"

	"golang.org/x/net/proxy"
)

func main() {

	d, err := proxy.SOCKS5("tcp", "localhost:9050", nil, &net.Dialer{})
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", ":47854")
	if err != nil {
		log.Fatal(err)
	}

	for {
		tc, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
        log.Print("accepted")

        sc, err := d.Dial("tcp", "2k3ffkhtbvxromr3.onion:22")
		if err != nil {
			log.Print(err)
			continue
		}

		go io.Copy(tc, sc)
		go io.Copy(sc, tc)
        log.Print("copying")

	}

}
