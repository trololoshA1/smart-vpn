package geoip

import (
    "log"
    "net"

    "github.com/oschwald/geoip2-golang"
)

type GeoIP struct {
    db *geoip2.Reader
}

// Load GeoIP database (GeoLite2-Country.mmdb)
func NewGeoIP(path string) (*GeoIP, error) {
    db, err := geoip2.Open(path)
    if err != nil {
        return nil, err
    }
    return &GeoIP{db: db}, nil
}

// Get country code by IP (RU, US, EU, etc.)
func (g *GeoIP) Country(ipStr string) string {
    ip := net.ParseIP(ipStr)
    if ip == nil {
        return "UNKNOWN"
    }

    record, err := g.db.Country(ip)
    if err != nil {
        log.Println("GeoIP error:", err)
        return "UNKNOWN"
    }

    return record.Country.IsoCode
}

// Decide routing:
// RU → direct
// others → tunnel
func (g *GeoIP) ShouldUseTunnel(ipStr string) bool {
    country := g.Country(ipStr)
    if country == "RU" {
        return false // direct
    }
    return true // tunnel
}