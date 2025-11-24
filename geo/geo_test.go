package geo

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"translate-shell-service/util"
)

type IpInfo struct {
	Status        string  `json:"status"`            // success or fail
	Message       string  `json:"message,omitempty"` // included only when status is fail Can be one of the following: private range, reserved range, invalid query
	Continent     string  `json:"continent"`         // Continent name
	ContinentCode string  `json:"continentCode"`     // Two-letter continent code
	Country       string  `json:"country"`           // Country name
	CountryCode   string  `json:"countryCode"`       // Two-letter country code ISO 3166-1 alpha-2
	Region        string  `json:"region"`            // Region/state short code (FIPS or ISO)
	RegionName    string  `json:"regionName"`        // Region/state
	City          string  `json:"city"`              // City
	District      string  `json:"district"`          // District (subdivision of city)
	Zip           string  `json:"zip"`               // Zip code
	Lat           float64 `json:"lat"`               // Latitude
	Lon           float64 `json:"lon"`               // Longitude
	Timezone      string  `json:"timezone"`          // Timezone (tz)
	Offset        int     `json:"offset"`            // Timezone UTC DST offset in seconds
	Currency      string  `json:"currency"`          // National currency
	Isp           string  `json:"isp"`               // ISP name
	Org           string  `json:"org"`               // Organization name
	As            string  `json:"as"`                // AS number and organization, separated by space (RIR). Empty for IP blocks not being announced in BGP tables.
	Asname        string  `json:"asname"`            // AS name (RIR). Empty for IP blocks not being announced in BGP tables.
	Reverse       string  `json:"reverse"`           // Reverse DNS of the IP (can delay response)
	Mobile        bool    `json:"mobile"`            // Mobile (cellular) connection
	Proxy         bool    `json:"proxy"`             // Proxy, VPN or Tor exit address
	Hosting       bool    `json:"hosting"`           // Hosting, colocated or data center
	Query         string  `json:"query"`             // IP used for the query
}

// 函数用于通过 IP 查询位置
func getIPInfo(ip string) IpInfo {
	host := strings.Join([]string{"http://ip-api.com/json/", ip}, "/")
	headers := map[string]string{"Content-Type": "application/json"}
	params := map[string]string{"fields": "status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,offset,currency,isp,org,as,asname,reverse,mobile,proxy,hosting,query"}
	resp, err := util.HttpGet(headers, params, host)
	if err != nil {
		fmt.Println(err)
	}
	var ipInfo IpInfo
	json.Unmarshal(resp, &ipInfo)
	return ipInfo

}

func TestIP(t *testing.T) {
	ret := getIPInfo("120.244.159.47")
	t.Logf("%+v\n", ret)
}
