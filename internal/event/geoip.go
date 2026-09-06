package event

import (
	"net"

	"github.com/oschwald/geoip2-golang"
)

type GeoIP struct {
	db *geoip2.Reader
}

type LocationInfo struct {
	Country string
	City    string
}

func NewGeoIP(path string) (*GeoIP, error) {
	db, err := geoip2.Open(path)
	if err != nil {
		return nil, err
	}

	return &GeoIP{
		db: db,
	}, nil
}

func (g *GeoIP) Lookup(ip string) LocationInfo {
	address := net.ParseIP(ip)
	if address == nil {
		return LocationInfo{}
	}

	record, err := g.db.City(address)
	if err != nil {
		return LocationInfo{}
	}

	return LocationInfo{
		Country: record.Country.Names["en"],
		City:    record.City.Names["en"],
	}
}

func (g *GeoIP) Close() error {
	return g.db.Close()
}