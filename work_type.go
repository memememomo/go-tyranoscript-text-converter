package main

type WorkType string

const (
	WorkTypeNone    WorkType = "none"
	WorkTypeReplace WorkType = "replace"
	WorkTypeBegin   WorkType = "begin"
	WorkTypeEnd     WorkType = "end"
	WorkTypeSplit   WorkType = "split"
)
