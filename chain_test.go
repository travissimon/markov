package main

import (
	"log"
	"testing"
)

func Test_ChainHookup(t *testing.T) {
	log.Printf("Hookup succeeded - testing (simple) markov chains")
}

func Test_Creation(t *testing.T) {
	prefixLen := 2
	chain := NewChain(prefixLen)

	if chain == nil {
		t.Errorf("Chain construction failed - got null result")
	}

	link := newLink(prefixLen)

	if link == nil {
		t.Errorf("Link construction failed - got null result")
	}

	if link.Prefix == nil {
		t.Errorf("Link prefix is nil")
	}
	if link.Suffixes == nil {
		t.Errorf("Link suffix is nil")
	}
}
