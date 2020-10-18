package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadAndProcess(t *testing.T) {
	actual := readAndProcess("fixtures/studio_config.json", "fixtures/target_file1.ks")
	assert.Equal(t, `吾輩わがはいは猫である。[r]
名前はまだ無い。[p]

;ここはコメントです。
どこで生れたかとんと見当けんとうがつかぬ。[r]
何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。[r]
吾輩はここで始めて人間というものを見た。[p]

[chara_mod name="あかね" face="笑顔"]
`, actual)
}
