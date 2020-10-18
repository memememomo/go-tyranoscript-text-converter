package main

type SentencePos int

const (
	SentencePosBegin SentencePos = iota + 1
	SentencePosIn
	SentencePosEnd
	SentencePosBlank
	SentencePosComment
)
