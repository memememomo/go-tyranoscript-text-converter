package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readAndProcess(configFile, targetFile string) string {
	configBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	var config Config
	if err = json.Unmarshal(configBytes, &config); err != nil {
		panic(err)
	}

	target, err := ioutil.ReadFile(targetFile)
	if err != nil {
		panic(err)
	}

	return process(&config, string(target))
}

func process(config *Config, target string) string {
	sentences := parseText(target)
	if os.Getenv("DEBUG") != "" {
		for _, s := range sentences {
			fmt.Printf("%+v\n", s)
		}
	}

	var output []string
	for _, sentence := range sentences {
		for i := range config.Rules {
			rule := config.Rules[i]
			if search(&rule, sentence) {
				sentence.Text = execute(&rule, sentence)
				if rule.Stop {
					break
				}
			}
		}
		output = append(output, sentence.Text)
	}

	return strings.Join(output, "\n")
}

func search(rule *Rule, target *SentenceInfo) bool {
	if rule.Search && matchText(rule.SearchMode, target.Text, rule.SearchWord) {
		return true
	}
	if rule.PosOn && matchPos(rule, target) {
		return true
	}
	return false
}

func execute(rule *Rule, target *SentenceInfo) string {
	switch rule.WorkType {
	case WorkTypeSplit:
		return splitAndReplace(target.Text, rule.WorkTextSplitAfter, rule.WorkTextSplit)
	case WorkTypeBegin:
		return fmt.Sprintf("%s%s", rule.WorkTextInsertBegin, target.Text)
	case WorkTypeEnd:
		return fmt.Sprintf("%s%s", target.Text, rule.WorkTextInsertEnd)
	case WorkTypeReplace:
		return strings.ReplaceAll(target.Text, rule.WorkTextReplaceBefore, rule.WorkTextReplaceAfter)
	default:
		return target.Text
	}
}

func parseText(target string) []*SentenceInfo {
	sentences := strings.Split(target, "\n")
	var infos = make([]*SentenceInfo, len(sentences))

	currentPos := SentencePosBegin
	for i, sentence := range sentences {
		if strings.HasPrefix(sentence, ";") || strings.HasPrefix(sentence, "#") {
			infos[i] = &SentenceInfo{
				Text: sentence,
				Pos:  SentencePosComment,
			}
		} else if sentence == "" {
			if i != 0 && (infos[i-1].Pos == SentencePosIn || infos[i-1].Pos == SentencePosBegin) {
				infos[i-1].Pos = SentencePosEnd
			}
			infos[i] = &SentenceInfo{
				Text: sentence,
				Pos:  SentencePosBlank,
			}
			currentPos = SentencePosBegin
		} else {
			infos[i] = &SentenceInfo{
				Text: sentence,
				Pos:  currentPos,
			}
			currentPos = SentencePosIn
		}
	}

	return infos
}

func matchPos(rule *Rule, target *SentenceInfo) bool {
	if rule.Pos1 && target.Pos == SentencePosBegin {
		return true
	}
	if rule.Pos2 && target.Pos == SentencePosIn {
		return true
	}
	return rule.Pos3 && target.Pos == SentencePosEnd
}

func matchText(searchMode SearchMode, target string, matchText string) bool {
	switch searchMode {
	case SearchModePrefix:
		return strings.HasPrefix(target, matchText)
	case SearchModeSuffix:
		return strings.HasSuffix(target, matchText)
	default:
		return strings.Contains(target, matchText)
	}
}

func splitAndReplace(target, newText, splitter string) string {
	elms := strings.Split(target, splitter)
	return replaceSplitText(newText, elms)
}

func replaceSplitText(target string, elms []string) string {
	_target := target
	for i, t := range elms {
		tag := fmt.Sprintf("${%d}", i+1)
		_target = strings.ReplaceAll(_target, tag, t)
	}
	return _target
}
