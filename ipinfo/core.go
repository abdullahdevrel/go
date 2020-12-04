package ipinfo

import (
	"net"
	"net/http"
)

// Core represents data from the Core API.
type Core struct {
	IP       net.IP       `json:"ip"`
	Hostname string       `json:"hostname"`
	City     string       `json:"city"`
	Region   string       `json:"region"`
	Country  string       `json:"country"`
	Location string       `json:"loc"`
	Org      string       `json:"org"`
	Postal   string       `json:"postal"`
	Timezone string       `json:"timezone"`
	ASN      *CoreASN     `json:"asn"`
	Company  *CoreCompany `json:"company"`
	Carrier  *CoreCarrier `json:"carrier"`
	Privacy  *CorePrivacy `json:"privacy"`
	Abuse    *CoreAbuse   `json:"abuse"`
	Domains  *CoreDomains `json:"domains"`
}

// CoreASN represents ASN data for the Core API.
type CoreASN struct {
	ASN    string `json:"asn"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Route  string `json:"route"`
	Type   string `json:"type"`
}

// CoreCompany represents company data for the Core API.
type CoreCompany struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

// CoreCarrier represents carrier data for the Core API.
type CoreCarrier struct {
	Name string `json:"name"`
	Mcc  string `json:"mcc"`
	Mnc  string `json:"mnc"`
}

// CorePrivacy represents privacy data for the Core API.
type CorePrivacy struct {
	Vpn     bool `json:"vpn"`
	Proxy   bool `json:"proxy"`
	Tor     bool `json:"tor"`
	Hosting bool `json:"hosting"`
}

// CoreAbuse represents abuse data for the Core API.
type CoreAbuse struct {
	Address string `json:"address"`
	Country string `json:"country"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Network string `json:"network"`
	Phone   string `json:"phone"`
}

// CoreDomains represents domains data for the Core API.
type CoreDomains struct {
	IP      string   `json:"ip"`
	Total   uint64   `json:"total"`
	Domains []string `json:"domains"`
}

// GetIPInfo returns the details for the specified IP.
func GetIPInfo(ip net.IP) (*Core, error) {
	return DefaultClient.GetIPInfo(ip)
}

// GetIPInfo returns the details for the specified IP.
func (c *Client) GetIPInfo(ip net.IP) (*Core, error) {
	var req *http.Request
	var err error

	if ip == nil {
		req, err = c.NewRequest("json")
	} else {
		req, err = c.NewRequest(ip.String() + "/json")
	}
	if err != nil {
		return nil, err
	}

	v := new(Core)
	_, err = c.Do(req, v)
	return v, err
}
