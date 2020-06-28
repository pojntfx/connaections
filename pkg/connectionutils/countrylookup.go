package connectionutils

import (
	"fmt"
	"net"

	"github.com/phuslu/geoip"
)

type Country struct {
	Code string
}

type CountryLookup struct {
}

func (cl *CountryLookup) GetCountryForIP(ip net.IP) string {
	return fmt.Sprintf("%s", geoip.Country(ip))
}
