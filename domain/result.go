package domain

type Result struct {
	Document Document
	Rank     float32
}

type Results []Result
