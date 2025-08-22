package main

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"basic", "hello"}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(); got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
