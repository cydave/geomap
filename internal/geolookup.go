package internal

import (
	"net"

	"github.com/oschwald/geoip2-golang"
)

type LookupService struct {
	db *geoip2.Reader
}

func NewLookupService(dbPath string) (*LookupService, error) {
	db, err := geoip2.Open(dbPath)
	if err != nil {
		return nil, err
	}
	return &LookupService{db}, nil
}

func (ls *LookupService) Close() {
	if ls.db != nil {
		ls.db.Close()
		ls.db = nil
	}
}

type LookupResult struct {
	IP        string  `json:"ip"`
	IsoCode   string  `json:"iso_code"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (ls *LookupService) Lookup(ipAddress net.IP) (*LookupResult, error) {
	record, err := ls.db.City(ipAddress)
	if err != nil {
		return nil, err
	}

	return &LookupResult{
		IP:        ipAddress.String(),
		IsoCode:   record.Country.IsoCode,
		Latitude:  record.Location.Latitude,
		Longitude: record.Location.Longitude,
	}, nil
}
