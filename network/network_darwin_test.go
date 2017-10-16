// +build darwin

package network

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetNetwork(t *testing.T) {
	networks, err := Get()
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	for _, network := range networks {
		if network.Name == "en0" && network.RxBytes <= 0 {
			t.Errorf("invalid network value: %+v", network)
		}
	}
	t.Logf("networks value: %+v", networks)
}

func TestCollectNetworkStats(t *testing.T) {
	got, err := collectNetworkStats(strings.NewReader(
		`Name  Mtu   Network       Address            Ipkts Ierrs     Ibytes    Opkts Oerrs     Obytes  Coll
lo0   16384 <Link#1>                       1758927     0  601295906  1758927     0  601295906     0
lo0   16384 127           127.0.0.1        1758927     -  601295906  1758927     -  601295906     -
lo0   16384 ::1/128     ::1                1758927     -  601295906  1758927     -  601295906     -
lo0   16384 fe80::1%lo0 fe80:1::1          1758927     -  601295906  1758927     -  601295906     -
gif0* 1280  <Link#2>                             0     0          0        0     0          0     0
stf0* 1280  <Link#3>                             0     0          0        0     0          0     0
en0   1500  <Link#4>    a4:5e:60:aa:aa:aa 23300388     0 18096041919 17727990     0 8602191509     0
en0   1500  fe80::222:a fe80:4::222:aaaa: 23300388     - 18096041919 17727990     - 8602191509     -
en0   1500  192.168.105   192.168.105.102 23300388     - 18096041919 17727990     - 8602191509     -
en1   1500  <Link#5>    4a:00:00:aa:aa:aa        0     0          0        0     0          0     0
en2   1500  <Link#6>    4a:00:00:aa:aa:aa        0     0          0        0     0          0     0
`))
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	expected := []NetworkStats{
		{"gif0", 0, 0},
		{"stf0", 0, 0},
		{"en0", 18096041919, 8602191509},
		{"en1", 0, 0},
		{"en2", 0, 0},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("invalid network value: %+v (expected: %+v)", got, expected)
	}
}
