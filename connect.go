package main

import (
	"log"
	"strings"

	"github.com/micro/mdns"
)

const (
	googleCastServiceName = "_googlecast._tcp"
	googleHomeModelInfo   = "md=Google Home"
)

type GoogleHomeInfo struct {
	Ip   string
	Port int
}

func LookupHomeIP() []*GoogleHomeInfo {
	entriesCh := make(chan *mdns.ServiceEntry, 4)

	results := []*GoogleHomeInfo{}
	go func() {
		for entry := range entriesCh {
			log.Printf("[INFO] ServiceEntry detected: [%s:%d]%s", entry.AddrV4, entry.Port, entry.Name)
			for _, field := range entry.InfoFields {
				if strings.HasPrefix(field, googleHomeModelInfo) {
					results = append(results, &GoogleHomeInfo{entry.AddrV4.String(), entry.Port})
				}
			}
		}
	}()

	mdns.Lookup(googleCastServiceName, entriesCh)
	close(entriesCh)

	return results
}
