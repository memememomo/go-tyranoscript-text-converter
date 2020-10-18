package main

type MatchType int

const (
	MatchTypePrefix MatchType = iota + 1
	MatchTypeSuffix
	MatchTypeContains
)
