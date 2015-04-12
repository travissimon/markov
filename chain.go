package main

import (
	"container/ring"
	"fmt"
)

type link struct {
	Prefix []string
	Suffixes []suffix
}

func (l *link) String() string {
	return fmt.Sprintf("%v -> %v", l.Prefix, l.Suffixes)
}

func (l *link) AddSuffix(s suffix) {
	l.Suffixes = append(l.Suffixes, s)
}

func newLink(prefixLen int) *link {
	l := &link{}
	l.Prefix = make([]string, prefixLen, prefixLen)
	l.Suffixes = make([]suffix, 1)

	return l
}

type suffix struct {
	Val string
	Source string
}

func newSuffix() *suffix {
	return &suffix{}
}

func (s *suffix) String() string {
	return fmt.Sprintf("%s", s.Val)
}

type Chain struct {
	prefixLen int

	// for building links
	inProgress *ring.Ring

	// map first prefix val to a list of links
	links map[string][]link
}

func NewChain(prefixLen int) *Chain {
	c := &Chain{}
	c.prefixLen = prefixLen
	c.inProgress = ring.New(prefixLen)
	c.links = make(map[string][]link)

	return c
}

