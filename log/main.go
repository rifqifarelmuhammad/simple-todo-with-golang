package log

import (
	"encoding/json"
	"fmt"
	"os"
)

func setFormat(tag string, metadata interface{}, err error, format string) string {
	data, _ := json.Marshal(metadata)
	return fmt.Sprintln(timeline(), traceline(), tag, string(data), err, format)
}

func setFormatSimplify(tag string, format string) string {
	return fmt.Sprintln(timeline(), traceline(), tag, format)
}

func setFormatFatal(tag string, err error, format string) string {
	return fmt.Sprintln(timeline(), traceline(), tag, err, format)
}

func Print(tag, format string, args ...interface{}) {
	printf(setFormatSimplify(tag, format), args...)
}

func Info(tag string, metadata interface{}, format string, args ...interface{}) {
	infof(setFormat(tag, metadata, nil, format), args...)
}

func Debug(tag string, metadata interface{}, err error, format string, args ...interface{}) {
	debugf(setFormat(tag, metadata, err, format), args...)
}

func Warn(tag string, metadata interface{}, err error, format string, args ...interface{}) {
	warnf(setFormat(tag, metadata, err, format), args...)
}

func Error(tag string, metadata interface{}, err error, format string, args ...interface{}) {
	errorf(setFormat(tag, metadata, err, format), args...)
}

func Fatal(tag string, err error, format string, args ...interface{}) {
	fatalf(setFormatFatal(tag, err, format), args...)
	os.Exit(1)
}
