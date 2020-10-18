package main

type SearchMode string

const (
	SearchModePrefix   SearchMode = "1"
	SearchModeSuffix   SearchMode = "2"
	SearchModeContains SearchMode = "3"
)
