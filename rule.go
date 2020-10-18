package main

type Rule struct {
	PosOn                 bool       `json:"pos_on"`
	Pos1                  bool       `json:"pos1"`
	Pos2                  bool       `json:"pos2"`
	Pos3                  bool       `json:"pos3"`
	Search                bool       `json:"search"`
	SearchMode            SearchMode `json:"search_mode"`
	SearchWord            string     `json:"search_word"`
	Stop                  bool       `json:"stop"`
	WorkType              WorkType   `json:"work_type"`
	WorkTextReplaceBefore string     `json:"work_text_replace_before"`
	WorkTextReplaceAfter  string     `json:"work_text_replace_after"`
	WorkTextInsertBegin   string     `json:"work_text_insert_begin"`
	WorkTextInsertEnd     string     `json:"work_text_insert_end"`
	WorkTextSplit         string     `json:"work_text_split"`
	WorkTextSplitAfter    string     `json:"work_text_split_after"`
}

type Config struct {
	Rules []Rule `json:"rules"`
}
