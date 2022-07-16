package main

import (
	"fmt"
	"testing"
)

func ExampleRun() {
	strIP := "10.0.100.186"
	ip := parseIP(strIP)
	fmt.Println(ip)
	fmt.Println(parseIP("19999.0.100.186"))
	// Output:
	// 10.0.100.186
	// invalid ip address
}

func TestParseIP(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			ip := parseIP(c.input)
			if ip != c.expect {
				t.Errorf("got '%v', expected '%v'", ip, c.expect)
			}
		})
	}
}

var testCases = []struct {
	name   string
	input  string
	expect string
}{
	{"valid", "10.0.100.186", "10.0.100.186"},
	{"invalid", "19999.0.100.186", "invalid ip address"},
	{"empty", "", "invalid ip address"},
}
