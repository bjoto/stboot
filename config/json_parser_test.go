package config

import (
	"bytes"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"testing"

	"github.com/vishvananda/netlink"
)

const (
	goodIPString   = "172.0.0.1"
	goodCIDRString = "127.0.0.1/24"
	goodMACString  = "00:00:5e:00:53:01"
	goodURLString  = "http://server.com"
)

func TestJSONParser(t *testing.T) {
	v := valuesFromGoodStrings(t)

	tests := []struct {
		name string
		json string
		want *HostCfg
	}{
		{
			name: "Version field",
			json: fmt.Sprintf(`{"%s": 1}`, VersionJSONKey),
			want: &HostCfg{Version: 1},
		},
		{
			name: "Network mode field 1",
			json: fmt.Sprintf(`{"%s": "%s"}`, NetworkModeJSONKey, StaticIP.String()),
			want: &HostCfg{NetworkMode: StaticIP},
		},
		{
			name: "Network mode field 2",
			json: fmt.Sprintf(`{"%s": "%s"}`, NetworkModeJSONKey, DynamicIP.String()),
			want: &HostCfg{NetworkMode: DynamicIP},
		},
		{
			name: "Host IP field",
			json: fmt.Sprintf(`{"%s": "%s"}`, HostIPJSONKey, goodCIDRString),
			want: &HostCfg{HostIP: v.cidr},
		},
		{
			name: "Gateway field",
			json: fmt.Sprintf(`{"%s": "%s"}`, DefaultGatewayJSONKey, goodIPString),
			want: &HostCfg{DefaultGateway: v.ip},
		},
		{
			name: "DNS Server field",
			json: fmt.Sprintf(`{"%s": "%s"}`, DNSServerJSONKey, goodIPString),
			want: &HostCfg{DNSServer: v.ip},
		},
		{
			name: "Network interface field",
			json: fmt.Sprintf(`{"%s": "%s"}`, NetworkInterfaceJSONKey, goodMACString),
			want: &HostCfg{NetworkInterface: v.mac},
		},
		{
			name: "Provisioning URLs field 1",
			json: fmt.Sprintf(`{"%s": ["%s"]}`, ProvisioningURLsJSONKey, goodURLString),
			want: &HostCfg{ProvisioningURLs: []*url.URL{v.provURL}},
		},
		{
			name: "Provisioning URLs field 2",
			json: fmt.Sprintf(`{"%s": ["%s", "%s"]}`, ProvisioningURLsJSONKey, goodURLString, goodURLString),
			want: &HostCfg{ProvisioningURLs: []*url.URL{v.provURL, v.provURL}},
		},
		{
			name: "Identity field",
			json: fmt.Sprintf(`{"%s": "some id"}`, IdJSONKey),
			want: &HostCfg{ID: "some id"},
		},
		{
			name: "Authentication field",
			json: fmt.Sprintf(`{"%s": "some auth"}`, AuthJSONKey),
			want: &HostCfg{Auth: "some auth"},
		},
		{
			name: "No fields",
			json: `{}`,
			want: &HostCfg{},
		},
		{
			name: "Empty fields",
			json: fmt.Sprintf(`{"%s": 0, "%s": "", "%s": "", "%s": "", "%s": "", "%s": "", "%s": [], "%s": "", "%s": ""}`, VersionJSONKey, NetworkModeJSONKey, HostIPJSONKey, DefaultGatewayJSONKey, DNSServerJSONKey, NetworkInterfaceJSONKey, ProvisioningURLsJSONKey, IdJSONKey, AuthJSONKey),
			want: &HostCfg{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JSONParser{bytes.NewBufferString(tt.json)}

			got, err := j.Parse()

			assertNoError(t, err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

type values struct {
	ip      *net.IP
	cidr    *netlink.Addr
	mac     *net.HardwareAddr
	provURL *url.URL
}

func valuesFromGoodStrings(t *testing.T) *values {
	t.Helper()

	i := net.ParseIP(goodIPString)
	if i == nil {
		t.Fatal("internal test error: invalid net.IP")
	}

	c, err := netlink.ParseAddr(goodCIDRString)
	if err != nil {
		t.Fatalf("internal test error: %v", err)
	}

	m, err := net.ParseMAC(goodMACString)
	if err != nil {
		t.Fatalf("internal test error: %v", err)
	}

	p, err := url.Parse(goodURLString)
	if err != nil {
		t.Fatalf("internal test error: %v", err)
	}

	v := &values{
		ip:      &i,
		cidr:    c,
		mac:     &m,
		provURL: p,
	}
	return v
}
