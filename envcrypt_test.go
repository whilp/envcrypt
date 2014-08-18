package main

import (
	"testing"
)

func TestParseOK(t *testing.T) {
	in := []byte("FOO=bar\nSPAM=eggs")
	out := []string{"FOO=bar", "SPAM=eggs"}
	if result, _ := parse(in); result[0] != out[0] {
		t.Errorf("")
	}
}
