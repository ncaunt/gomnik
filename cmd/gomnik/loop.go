package main

import (
	"net"
	"time"

	"github.com/ncaunt/gomnik"
)

func loop(ticker *time.Ticker, m *gomnik.Metrics) (err error) {
	req := gomnik.NewRequest(*serial)
	b, err := req.Bytes()

	for range ticker.C {
		conn, err2 := net.Dial("tcp", *addr)
		if err2 != nil {
			continue
		}

		conn.Write(b)

		buffer := make([]byte, 1024)
		conn.Read(buffer)
		conn.Close()

		r, err := gomnik.NewResponse(buffer)
		if err != nil {
			continue
		}

		m.SetFromResponse(r)
	}

	return
}
