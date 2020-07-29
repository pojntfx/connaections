package readers

// Get the DB from https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=G4Lm0C60yJsnkdPi&suffix=tar.gz

import (
	"net"

	geodb "github.com/oschwald/geoip2-golang"
)

type Location struct {
	Latitude    float64
	Longitude   float64
	CityName    string
	CountryCode string
}

type LocationReader struct {
	DbPath string

	db *geodb.Reader
}

func (lr *LocationReader) Open() error {
	db, err := geodb.Open(lr.DbPath)
	if err != nil {
		return err
	}

	lr.db = db

	return nil
}

func (lr *LocationReader) GetLocationForIP(ip net.IP) (Location, error) {
	record, err := lr.db.City(ip)
	if err != nil {
		return Location{}, err
	}

	cityNames := record.City.Names["en"]
	cityNameEnd := ""
	if len(cityNameEnd) > 0 {
		cityNameEnd = string(cityNames[0])
	}

	return Location{
		record.Location.Latitude,
		record.Location.Longitude,
		cityNameEnd,
		record.Country.IsoCode,
	}, nil
}
