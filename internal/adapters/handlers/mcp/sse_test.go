package mcpadapter

import "testing"

func TestTrimHostPort(t *testing.T) {
	for _, tt := range []struct {
		addr, want string
	}{
		{":8081", "localhost:8081"},
		{"127.0.0.1:9090", "127.0.0.1:9090"},
		{"localhost:80", "localhost:80"},
	} {
		if got := trimHostPort(tt.addr); got != tt.want {
			t.Errorf("trimHostPort(%q) = %q, want %q", tt.addr, got, tt.want)
		}
	}
}
