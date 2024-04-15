package loggeria

import (
	"errors"
	"fmt"
	"strings"
)

type Level [2]int8

var (
	LevelTrace    Level = [2]int8{-1, 0}
	LevelDebug    Level = [2]int8{0, 0}
	LevelInfo     Level = [2]int8{1, 0}
	LevelSub      Level = [2]int8{1, 1}
	LevelWarn     Level = [2]int8{2, 0}
	LevelError    Level = [2]int8{3, 0}
	LevelPanic    Level = [2]int8{4, 0} // panic is less severe than fatal, because panic can be recovered,
	LevelFatal    Level = [2]int8{5, 0} // and fatal causes explicit OS process exit
	LevelDisabled Level = [2]int8{6, 0}

	LevelToText map[Level]string
	TextToLevel map[string]Level
)

// this architecture allows to add custom user levels
func RegisterNewLevel(lvl Level, text string) {
	text = strings.ToUpper(text)
	LevelToText[lvl] = text
	TextToLevel[text] = lvl
}

func (l Level) String() string {
	if text, ok := LevelToText[l]; ok {
		return text
	} else {
		return fmt.Sprintf("[%v %v]", l[0], l[1])
	}
}

func LevelFromText(text string) (lvl Level, err error) {
	text = strings.ToUpper(text)
	var ok bool
	if lvl, ok = TextToLevel[text]; ok {
		return lvl, nil
	} else {
		return Level{}, errors.New("logging level unknown")
	}
}

func init() {
	LevelToText = make(map[Level]string)
	TextToLevel = make(map[string]Level)

	RegisterNewLevel(LevelTrace, "TRACE")       // TRC
	RegisterNewLevel(LevelDebug, "DEBUG")       // DBG
	RegisterNewLevel(LevelInfo, "INFO")         // INF
	RegisterNewLevel(LevelSub, "SUB")           // SUB
	RegisterNewLevel(LevelWarn, "WARN")         // WRN
	RegisterNewLevel(LevelError, "ERROR")       // ERR
	RegisterNewLevel(LevelPanic, "PANIC")       // PNC
	RegisterNewLevel(LevelFatal, "FATAL")       // FTL
	RegisterNewLevel(LevelDisabled, "DISABLED") // DIS
}
