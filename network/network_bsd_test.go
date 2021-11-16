//go:build darwin || freebsd || netbsd
// +build darwin freebsd netbsd

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
	expected := []Stats{
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

func TestCollectNetworkStats2(t *testing.T) {
	got, err := collectNetworkStats(strings.NewReader(
		`Name    Mtu Network       Address              Ipkts Ierrs Idrop     Ibytes    Opkts Oerrs     Obytes  Coll
vtnet  1500 <Link#1>      08:00:27:aa:aa:aa     2023     0     0     154650     1231     0     134810     0
vtnet     - 10.0.2.0/24   10.0.2.15             2018     -     -     125640     1227     -     117164     -
vtnet  1500 <Link#2>      08:00:27:bb:bb:bb     1022     0     0     239016      991     0     161902     0
vtnet     - 192.168.10.0/ 192.168.10.100        1017     -     -     224437      988     -     147944     -
em0    1500 <Link#3>      08:00:27:cc:cc:cc      114     0     0      17587        1     0         84     0
em0       - 192.168.20.0/ 192.168.20.200           0     -     -          0        0     -          0     -
lo0   16384 <Link#4>      lo0                      0     0     0          0        0     0          0     0
lo0       - ::1/128       ::1                      0     -     -          0        0     -          0     -
lo0       - fe80::%lo0/64 fe80::1%lo0              0     -     -          0        0     -          0     -
lo0       - 127.0.0.0/8   127.0.0.1                0     -     -          0        0     -          0     -
`))
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	expected := []Stats{
		{"vtnet", 154650, 134810},
		{"vtnet", 239016, 161902},
		{"em0", 17587, 84},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("invalid network value: %+v (expected: %+v)", got, expected)
	}
}
