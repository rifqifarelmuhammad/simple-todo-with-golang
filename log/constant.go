package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/fatih/color"
)

var (
	timeline = func() string {
		return time.Now().Format("2006-01-02 15:04:05")
	}

	traceline = func() string {
		pc, file, line, ok := runtime.Caller(3)
		if ok {
			file = filepath.Base(file)
			_ = filepath.Base(runtime.FuncForPC(pc).Name())
			return fmt.Sprintf("%s:%d", file, line)
		}
		return "unknown"
	}

	white   = color.New(color.FgWhite)
	green   = color.New(color.FgGreen)
	blue    = color.New(color.FgBlue)
	yellow  = color.New(color.FgYellow)
	red     = color.New(color.FgRed)
	redBold = color.New(color.FgRed, color.Bold)

	printf = white.PrintfFunc()
	infof  = green.PrintfFunc()
	debugf = blue.PrintfFunc()
	warnf  = yellow.PrintfFunc()
	errorf = red.PrintfFunc()
	fatalf = redBold.PrintfFunc()
)
