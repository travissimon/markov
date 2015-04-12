package main

import (
	"container/ring"
	"fmt"
)

type Link struct {
	Prefix []string
	Suffixes []Suffix
}

func (l *Link) String() string {
	return fmt.Sprintf("%v -> %v", l.Prefix, l.Suffixes)
}

func NewLink(prefixLen int) *Link {
	l := &Link{}
	l.Prefix = make([]string, prefixLen, prefixLen)
	l.Suffixes = make([]Suffix, 1)

	return l
}

type Suffix struct {
	Val string
	Source string
}

func NewSuffix() *Suffix {
	return &Suffix{}
}

func (s *Suffix) String() string {
	return fmt.Sprintf("%s", s.Val)
}

type Chain struct {
	prefixLen int

	// for building links
	inProgress *ring.Ring

	// map first prefix val to a list of links
	links map[string][]Link
}

func NewChain(prefixLen int) *Chain {
	c := &Chain{}
	c.prefixLen = prefixLen
	c.inProgress = ring.New(prefixLen)
	c.links = make(map[string][]Link)

	return c
}

